## Project Phases

This document describes the planned phases for nextpress-backend: **what is implemented** and **what to do next**.

### Snapshot — where we are

| Phase | Summary |
|-------|---------|
| **1** | Infra, Gin, GORM, config, Zap, migrations (`pkg/migrate`), deploy scripts, Nginx/systemd docs |
| **2** | JWT access/refresh, bcrypt, `user` + `auth` modules, `/v1/auth/*` |
| **3** | RBAC schema, `RequirePermission`, admin RBAC APIs, seed, optional bootstrap |
| **4** | Posts, pages, taxonomy, media, menus; admin + public APIs; rate limits, request ID, OpenAPI, tests |
| **5** *(in progress)* | `plugins` table, admin plugin CRUD, `HookRegistry` chain, `posts/domain.PostSave` wired on create/update |

**Next focus (recommended):** Phase 5 — map enabled plugins to real hook implementations (and/or page/menu hooks), optional dynamic routes; then Phase 6 admin/settings or Phase 7 sample plugin.

---

### Phase 1 – Core Infrastructure

**Goals:**

- Initialize project as a modular monolith in Go.
- Add HTTP server with Gin.
- Configure PostgreSQL connection via GORM with a global DB instance.
- Introduce configuration system and logging.
- Provide a migration system.
- Provide deployment tooling (Makefile, systemd, Nginx, docs, scripts).

**Current status:** Done — see repository `cmd/api`, `internal/server`, `internal/platform`, `migrations/`, `scripts/deploy`, `docs/deployment/`.

---

### Phase 2 – Authentication

**Goals:**

- Implement authentication and basic user management.
- Use JWT-based access and refresh tokens.
- Use bcrypt for password hashing.
- Keep logic in DDD-style `user` and `auth` modules.
- Expose auth endpoints for register, login, and refresh.

**Current status:** Done — `POST /v1/auth/register`, `/login`, `/refresh`; JWT middleware on `/v1/admin/*`.

---

### Phase 3 – RBAC (Roles and Permissions)

**Goals:**

- Implement Role-Based Access Control (RBAC).
- Support users, roles, permissions, and their relations.
- Add middleware to enforce permissions on routes.
- Provide admin APIs to manage roles and permissions.

**Current status:**

- RBAC database schema is present (`roles`, `permissions`, `user_roles`, `role_permissions`).
- Authorization middleware exists (`RequirePermission`) and is wired with a sample protected route:
  - `GET /v1/admin/ping` requires `admin:ping`.
- Admin RBAC APIs exist (guarded by `rbac:manage`):
  - `GET /v1/admin/roles`, `POST /v1/admin/roles`
  - `GET /v1/admin/permissions`, `POST /v1/admin/permissions`
  - `POST /v1/admin/roles/:role_id/permissions` (grant permission to role)
  - `POST /v1/admin/users/:user_id/roles` (assign role to user)
- RBAC defaults are seeded via `make seed` / `go run ./cmd/seed` (admin role + base permissions).
- Optional one-time bootstrap endpoint (guarded by auth + env flag):
  - `POST /v1/admin/bootstrap/claim-admin` (requires `RBAC_BOOTSTRAP_ENABLED=true`)

---

### Phase 4 – CMS Core

**Goals:**

- Implement core CMS entities:
  - posts, pages, media, categories, tags, menus.
- Provide CRUD APIs and relations for content.
- Enable filtering, searching, and listing for content entities.

**Current status:**

- Posts: schema + basic CRUD API exists (RBAC-protected):
  - `GET /v1/admin/posts` (requires `posts:read`; supports `status`, `author_id`, `q`, `limit`, `offset`)
  - `GET /v1/admin/posts/:id` (requires `posts:read`)
  - `POST /v1/admin/posts` (requires `posts:write`)
  - `PUT /v1/admin/posts/:id` (requires `posts:write`)
  - `DELETE /v1/admin/posts/:id` (requires `posts:write`)
  - `PUT /v1/admin/posts/:id/categories` (requires `posts:write`)
  - `PUT /v1/admin/posts/:id/tags` (requires `posts:write`)
- Pages: schema + basic CRUD API exists (RBAC-protected):
  - `GET /v1/admin/pages` (requires `pages:read`; supports `status`, `author_id`, `q`, `limit`, `offset`)
  - `GET /v1/admin/pages/:id` (requires `pages:read`)
  - `POST /v1/admin/pages` (requires `pages:write`)
  - `PUT /v1/admin/pages/:id` (requires `pages:write`)
  - `DELETE /v1/admin/pages/:id` (requires `pages:write`)
- Taxonomy: categories and tags basic CRUD APIs exist (RBAC-protected):
  - `GET /v1/admin/categories` (requires `categories:read`)
  - `POST /v1/admin/categories` (requires `categories:write`)
  - `PUT /v1/admin/categories/:id` (requires `categories:write`)
  - `DELETE /v1/admin/categories/:id` (requires `categories:write`)
  - `GET /v1/admin/tags` (requires `tags:read`)
  - `POST /v1/admin/tags` (requires `tags:write`)
  - `PUT /v1/admin/tags/:id` (requires `tags:write`)
  - `DELETE /v1/admin/tags/:id` (requires `tags:write`)
- Media: upload + list APIs exist (RBAC-protected):
  - `POST /v1/admin/media` (multipart form field `file`, requires `media:write`)
  - `GET /v1/admin/media` (requires `media:read`)
  - `GET /v1/admin/media/:id` (requires `media:read`)
  - uploads are served at `MEDIA_PUBLIC_BASE_URL` (default `/uploads`)
- Menus: schema + basic APIs exist (RBAC-protected):
  - `GET /v1/admin/menus` (requires `menus:read`)
  - `POST /v1/admin/menus` (requires `menus:write`)
  - `GET /v1/admin/menus/:id` (requires `menus:read`)
  - `PUT /v1/admin/menus/:id` (requires `menus:write`)
  - `DELETE /v1/admin/menus/:id` (requires `menus:write`)
  - `GET /v1/admin/menus/:id/items` (requires `menus:read`)
  - `PUT /v1/admin/menus/:id/items` (requires `menus:write`)
- Public APIs (no auth):
  - `GET /v1/posts` (published only; supports `q`, `category_id`, `tag_id`, `limit`, `offset`)
  - `GET /v1/posts/:slug` (published only)
  - `GET /v1/pages/:slug` (published only)
  - `GET /v1/menus/:slug` (items returned as a tree)
- Phase 4 hardening:
  - In-memory rate limiting middleware added for `public`, `auth`, and `admin` API groups.
  - Request correlation via `X-Request-ID` + improved structured request logging (latency/client IP/user id).
  - OpenAPI spec added at `docs/openapi.yaml`.
  - `httptest` coverage for critical middleware/route wiring exists.

---

### Phase 5 – Plugin System

**Goals:**

- Provide a plugin mechanism for extending the CMS without modifying core.
- Start with a database-driven plugin model (no Go plugins initially).
- Allow plugins to register routes, permissions, migrations, and hooks.
- Implement a plugin loader that wires plugins at startup.

**Current status (A0 + A1):**
- Database model: `plugins` table migration added (UUID id, slug/name, enabled, version, config JSONB).
- Plugin registry module added (`internal/modules/plugins`):
  - Admin endpoints: `GET /v1/admin/plugins`, `POST /v1/admin/plugins`, `PUT /v1/admin/plugins/:id` (RBAC: `plugins:manage`).
- Hook infrastructure:
  - `posts/domain.PostSave` is the port; posts `Create`/`Update` call `BeforePostSave` / `AfterPostSave` around persistence.
  - `HookRegistry` chains `PostHooks` implementations; bootstrap registers one noop slot per **enabled** plugin row (ready for real handlers).

---

### Phase 6 – Admin API

**Goals:**

- Provide an admin-facing API for:
  - dashboard/analytics,
  - settings management,
  - plugin management.
- Expose secure endpoints for environment and configuration management.

---

### Phase 7 – Example Ecommerce Plugin

**Goals:**

- Demonstrate plugin capabilities with a non-trivial example:
  - products, orders, cart, payments.
- Implement ecommerce as a plugin, not part of the core.
- Provide example APIs for catalog, cart, and checkout flows.

---

## Continuing development (checklist)

1. **Database:** `make migrate-up` after pulling; `make seed` when RBAC defaults change.
2. **Phase 5:** Replace noop hook slots with real handlers keyed by plugin `slug` / `config` JSON; consider transactions if hooks must roll back with the DB write.
3. **Hardening:** Shared rate-limit store (Redis) for multi-instance; expand `httptest` / integration tests with test DB.
4. **Docs:** Keep `README.md`, `docs/PHASES.md`, and `docs/openapi.yaml` in sync when adding endpoints.
5. **Git:** Follow **`docs/GIT_FLOW.md`** when promoting `dev` → `staging` → `main` or syncing branches after releases.

See also: `docs/README.md` (documentation index).
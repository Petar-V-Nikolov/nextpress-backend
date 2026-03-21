# nextpress-backend

Production-oriented **CMS API** in Go: **modular monolith**, **PostgreSQL + GORM**, **Gin**, **JWT auth**, **RBAC**, **CMS core** (posts, pages, taxonomy, media, menus), **public read APIs**, **rate limiting**, **plugin registry + post-save hooks** (Phase 5 in progress).

---

## Current status (snapshot)

| Area | Status |
|------|--------|
| **Phase 1** – infra, config, migrations, deploy tooling | Done |
| **Phase 2** – register / login / refresh, bcrypt, JWT | Done |
| **Phase 3** – RBAC, seed, admin RBAC APIs | Done |
| **Phase 4** – CMS CRUD + public `/v1/*` reads + hardening (rate limits, `X-Request-ID`, OpenAPI, tests) | Done |
| **Phase 5** – `plugins` table, admin plugin CRUD, `HookRegistry`, posts `PostSave` hooks | **A0–A1 done**; real plugin handlers & loader next |
| **Phase 6–7** | Planned (admin dashboard API, example ecommerce plugin) |

Details and next steps: **`docs/PHASES.md`**.

---

## Stack

- Go 1.26+
- Gin, Zap, GORM, PostgreSQL
- JWT access + refresh (`github.com/golang-jwt/jwt/v5`), bcrypt passwords

---

## Quick start (local)

```bash
cp .env.example .env
# Edit .env — set DB_* for PostgreSQL
go mod download
make migrate-up    # apply SQL migrations
make seed          # RBAC defaults (admin role + permissions)
make run           # or: go run ./cmd/api
```

Health: `GET /health`, `GET /ready` (DB check).

**Tests:** `go test ./...` · **Vet:** `go vet ./...` · **Build:** `make build` → `bin/server`

**API contract:** `docs/openapi.yaml`

---

## Configuration

Primary reference: **`.env.example`** (all variables with short comments).

| Group | Examples |
|-------|-----------|
| App | `APP_NAME`, `APP_ENV`, `APP_PORT` |
| DB | `DB_HOST`, `DB_PORT`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`, `DB_SSLMODE` |
| JWT | `JWT_SECRET`, `JWT_ACCESS_TTL`, `JWT_REFRESH_TTL` |
| RBAC | `RBAC_BOOTSTRAP_ENABLED` (optional first-admin bootstrap) |
| Media | `MEDIA_STORAGE_DIR`, `MEDIA_PUBLIC_BASE_URL`, `MEDIA_MAX_UPLOAD_BYTES` |
| Rate limits | `RATE_LIMIT_ENABLED`, `RATE_LIMIT_*_MAX_PER_MINUTE` |

---

## Repository layout

```text
.
├── cmd/
│   ├── api/           # HTTP API entrypoint
│   ├── migrate/       # SQL migrations runner
│   └── seed/          # Database seeders
├── internal/
│   ├── config/        # Env-based configuration
│   ├── modules/       # Feature slices (auth, user, rbac, posts, pages, …)
│   ├── platform/      # DB, logger, middleware
│   └── server/        # Gin engine, global middleware, health routes
├── migrations/        # Timestamped SQL (pkg/migrate)
├── pkg/
│   ├── migrate/       # Migration runner
│   └── seed/          # Seeder entry + RBAC defaults
├── deploy/            # nginx, systemd templates
├── docs/              # Phases, deployment, OpenAPI, seeding
├── scripts/           # deploy, run_local.sh
└── Makefile
```

---

## API layout (summary)

- **Public (no auth):** `GET /v1/posts`, `GET /v1/posts/:slug`, `GET /v1/pages/:slug`, `GET /v1/menus/:slug`
- **Auth:** `POST /v1/auth/register`, `/login`, `/refresh`
- **Admin (JWT + permission):** `/v1/admin/*` — CMS CRUD, RBAC management, `GET/POST /v1/admin/plugins`, `PUT /v1/admin/plugins/:id` (`plugins:manage`)

Full list: **`docs/openapi.yaml`**.

---

## Git workflow

- **`dev`** — integration / daily work  
- **`staging`** — pre-production  
- **`main`** — production  

Promote: merge `dev` → `staging` → `main` (push each).  
Sync back: merge `main` into `staging` and `dev` when you need them aligned.

Suggested server folders: `/var/www/nextpress-backend-{dev,staging,production}` — see **`docs/DEPLOYMENT.md`**.

---

## Documentation

- **Roadmap & next steps:** `docs/PHASES.md`  
- **Index of all docs:** `docs/README.md`  
- **Deploy:** `docs/DEPLOYMENT.md` + `docs/deployment/*.md`  
- **Seeders:** `docs/SEEDING.md`

---

## Architecture notes

- **Modular monolith** — one process, one DB; boundaries under `internal/modules/*`.
- **Clean layering** — transport → application → domain (+ infrastructure per module).
- **Post-save hooks** — `posts/domain.PostSave` port; `HookRegistry` in plugins module implements it; wired in `cmd/api`.

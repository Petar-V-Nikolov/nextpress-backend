# Database Seeding

Seeders populate reference data and optional development data. They are **idempotent**: running them multiple times will not duplicate rows that use `ON CONFLICT DO NOTHING` (permissions by `code`, roles by `name`, etc.).

## Prerequisites

- Database migrations have been applied (`make migrate-up`).
- `.env` has valid `DB_*` settings so the seed command can connect (same as `cmd/migrate`).

## How to run

```bash
make seed
# or
go run ./cmd/seed
# or build
make seed-build && ./bin/seed
```

## What is seeded

### RBAC defaults (`pkg/seed/rbac_defaults.go`)

**Role**

| Name   | Notes        |
|--------|--------------|
| `admin`| Seeded UUID `00000000-0000-0000-0000-000000000001` |

**Permissions** (all granted to `admin`)

| Code               | Used for |
|--------------------|----------|
| `admin:ping`       | `GET /v1/admin/ping` |
| `rbac:manage`      | Admin RBAC APIs (`/v1/admin/roles`, `/permissions`, assign role/permission) |
| `posts:read` / `posts:write` | Post CRUD + taxonomy assignment on posts |
| `pages:read` / `pages:write` | Page CRUD |
| `categories:read` / `categories:write` | Categories |
| `tags:read` / `tags:write` | Tags |
| `media:read` / `media:write` | Media upload/list |
| `menus:read` / `menus:write` | Menus + items |
| `plugins:manage`   | `GET/POST /v1/admin/plugins`, `PUT /v1/admin/plugins/:id` |

After seeding, assign the `admin` role to a user (via RBAC API or one-time bootstrap — see `docs/PHASES.md` Phase 3).

## Related

- **Migrations:** `make migrate-up` / `cmd/migrate`
- **Deploy:** optional seed on deploy via `RUN_SEED_ON_DEPLOY` (see `scripts/deploy`)

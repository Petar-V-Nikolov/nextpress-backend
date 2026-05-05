# Roadmap

High-level product direction.
Detailed engineering checklist is in [`TODO.md`](TODO.md).

---

## Shipped

- **Platform:** Gin, GORM (AutoMigrate + `cmd/migrate` / `cmd/seed`), config, logging—see [`COMMANDS.md`](COMMANDS.md#database-and-seed-data) and [`DEPLOYMENT.md`](DEPLOYMENT.md).
- **Modular kit:** `kit.Module` registry (`internal/appregistry`), `MODULES` env, shared `cmd/migrate` / `cmd/seed` / API composition (`internal/app`).
- **Auth:** Register/login/refresh, JWT access + refresh, bcrypt.
- **RBAC:** Roles, permissions, middleware, admin APIs, seeded defaults, optional bootstrap.
- **Content & admin APIs:** Posts, pages, taxonomy, media; public + admin HTTP APIs; rate limits, request ID, OpenAPI.
- **Elasticsearch (optional):** posts search + indexing when enabled.

---

## In progress

Current focus:
- better test coverage
- module-combination smoke tests

---

## Later

Main later themes:
- optional object storage backend
- preset product profiles
- plugin model as separate service/kit

---

## Historical note

Numbered **phases** (1-5) were internal planning labels during early development; they are retired in favor of the sections above.

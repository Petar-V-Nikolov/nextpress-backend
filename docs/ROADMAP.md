# Roadmap

**Explanation** — product scope and direction. **Every checkbox** (shipped vs open): [`TODO.md`](TODO.md) — use **`[ ]`** lines as the backlog; this page stays short.

**Related:** [Documentation index](README.md) · [Contributing](../CONTRIBUTING.md) · [Commands](COMMANDS.md) · **REST contract** [`openapi.yaml`](openapi.yaml) · [Module kit](MODULES.md)

Keep this file short: shipped capabilities, what you are actively improving, and rough future themes. Owners and threading: your issue tracker + [`TODO.md`](TODO.md).

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

Typical focus: **test coverage** for transport/application layers, **smoke tests** for minimal `MODULES` sets. Details: [`TODO.md`](TODO.md).

---

## Later

See unchecked items under **Future**, **Testing**, and **Security** in [`TODO.md`](TODO.md). Dynamic **plugins** as a separate kit or service (not in this core repo).

---

## Historical note

Numbered **phases** (1-5) were internal planning labels during early development; they are retired in favor of the sections above.

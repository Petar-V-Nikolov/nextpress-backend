# Kit modules

[← Documentation index](README.md)

NextPressKit is a **modular monolith**: each vertical slice lives under `internal/modules/<name>/` and is registered through a small `module` package plus the shared contract in [`internal/kit`](../internal/kit/module.go).

## Default registry

The compile-time default list (order matters for migrations and HTTP wiring) lives in [`internal/appregistry/registry.go`](../internal/appregistry/registry.go):

| Order | Module id | Responsibility |
|-------|-----------|------------------|
| 1 | `user` | Users table + demo user seed |
| 2 | `rbac` | Roles, permissions, RBAC admin APIs, `/admin/ping`, optional bootstrap |
| 3 | `auth` | Register/login/refresh/logout/me |
| 4 | `taxonomy` | Categories and tags |
| 5 | `media` | Uploads + local static serving when `MEDIA_PUBLIC_BASE_URL` is a path |
| 6 | `posts` | Posts + public routes + Elasticsearch hooks + scheduled publish loop |
| 7 | `pages` | Pages + public routes |

## Choosing modules at runtime

Set **`MODULES`** in `.env` to a comma-separated list of ids (case-insensitive). Empty or unset = full default registry. From an interactive terminal, **`make setup` → Add kit module / Remove kit module** can edit this line for you (then run `migrate-up` / `seed` as needed).

Implicit dependencies are applied so migrations and typical admin stacks stay consistent:

- `auth` → also enables `user`, `rbac`
- Any of `posts`, `pages`, `taxonomy`, `media` → also enables `user`, `rbac`, `auth`
- `posts` → also enables `taxonomy`, `media` (FKs / featured media)

Unknown ids are logged and ignored.

Examples:

```bash
# Minimal API (auth + RBAC only)
MODULES=auth,user,rbac

# Full CMS (same as empty MODULES)
MODULES=user,rbac,auth,taxonomy,media,posts,pages
```

## Choosing modules at compile time

To change the default product permanently, edit [`internal/appregistry/registry.go`](../internal/appregistry/registry.go): add/remove entries from `ModuleRegistry()` or reorder (keep FK-safe migrate order: `user` before `rbac` before content tables, etc.).

## What runs per command

- **`cmd/api`** — resolves modules, runs `Prepare` / `Register*` / `Start` for each.
- **`cmd/migrate`** — runs `AutoMigrate` only for resolved modules (same `MODULES` rules).
- **`cmd/seed`** — runs `SeedRBACDefaults` with merged permission codes from enabled modules, then each module’s `Seed` inside one transaction.

## Adding a new module

1. Create `internal/modules/<name>/` with your usual `domain`, `application`, `persistence`, `transport` layout.
2. Add `internal/modules/<name>/module/module.go` implementing [`kit.Module`](../internal/kit/module.go).
3. Append the module to `ModuleRegistry()` in [`internal/appregistry/registry.go`](../internal/appregistry/registry.go) in the correct position for FKs.
4. Return permission codes from `Permissions()` if the module introduces RBAC codes seeded for the admin role.
5. Update [`docs/openapi.yaml`](openapi.yaml) and Postman templates when you add HTTP surface.

## Plugins

Dynamic WordPress-style plugins are **not** part of this core kit. Post-save extensibility is limited to hooks composed in the posts module (e.g. derived fields + optional Elasticsearch). A separate kit or service can reintroduce plugin discovery later.

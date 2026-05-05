# ADR 0001: Module composition for the kit

## Context

The API entrypoint had grown into a single large `main` that wired every feature, while `cmd/migrate` and `cmd/seed` duplicated which tables and seeds existed. Products built from the kit need to enable or disable slices without hunting through many files.

## Decision

- Introduce a small **`kit.Module`** contract (`Prepare`, `RegisterAuth` / `RegisterPublic` / `RegisterAdmin`, `AutoMigrate`, `Seed`, `Start`, `Permissions`) and shared **`kit.Deps`** for Gin groups, DB, config, JWT, Elasticsearch index handle, and cross-module wiring.
- Keep a single **compile-time registry** in **`internal/appregistry`** (importable from `cmd/*` and `pkg/seed` without import cycles with `internal/app`).
- **`internal/app.Run`** owns process lifetime: infra, module phases, HTTP server, graceful shutdown.
- **`MODULES`** env var filters the registry at runtime; implicit dependencies expand minimal sets.
- **Remove** the in-tree `plugins` module; post-save behavior stays inside posts + platform Elasticsearch.

## Consequences

- One place to add/remove features for a new product (`appregistry` + optional `MODULES`).
- Migrate and seed stay aligned with the API surface.
- Slightly more boilerplate per feature (`module` package).

## Alternatives considered

- **Build tags per product** — stronger compile-time exclusion; can be layered later.
- **Multi-module Go repo** — better for publishing slices independently; higher release overhead for a kit.
- **Runtime-only plugin registry** — flexible but kept complexity and DB-driven hooks out of the core kit.

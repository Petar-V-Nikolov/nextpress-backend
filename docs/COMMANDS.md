# Commands Guide

What to run, in plain language. **Make** has three targets; **nextpresskit** is the full CLI for every workflow.

[← Documentation index](README.md) · [Quick start](../README.md#quick-start) · [Contributing checks](../CONTRIBUTING.md#before-you-open-a-pr)

## Make (three targets)

| Target | Purpose |
|--------|---------|
| `make setup` | **TTY:** text menu — profiles (full / database / build / CI), Custom (numbered CLI steps), **Add/Remove kit module** (`MODULES` in `.env`). **Non-TTY** or **`NP_SETUP_NONINTERACTIVE=1`:** linear `install` → `build-all` → `migrate-up` → `seed` (+ optional `setup-local-https.sh` unless `SKIP_SETUP_LOCAL_HTTPS=1`). |
| `make run` | API in the foreground (same as `./scripts/nextpresskit run`). |
| `make postman-sync` | Sync gitignored `postman/` from templates (same as `./scripts/nextpresskit postman-sync`). |

Running **`make`** with no target prints this summary and points to `./scripts/nextpresskit help`.

## Command families

- `./scripts/nextpresskit ...` (Unix / Git Bash): full command runner.
- `make setup|run|postman-sync`: thin wrappers for day-to-day use.
- `.\scripts\nextpresskit.ps1 ...` (Windows PowerShell): Windows equivalent (same text menu for `setup` when the console is interactive).

## Choose by task

| I want to… | Use this |
|------------|----------|
| Set up a fresh clone | `make setup` or `./scripts/nextpresskit setup` (menu: full = option **1**; automation: `NP_SETUP_NONINTERACTIVE=1`) |
| Run the API now | `make run` or `./scripts/nextpresskit run` |
| Run checks before PR | `./scripts/nextpresskit checks` |
| Create or update database tables | `./scripts/nextpresskit migrate-up` |
| Load demo data (run after migrate) | `./scripts/nextpresskit seed` |
| Start over on a dev database | `./scripts/nextpresskit db-fresh`, then `seed` (see [Database and seed data](#database-and-seed-data)) |
| Generate deploy config | `./scripts/nextpresskit deploy` or `bash scripts/deploy` / `.\scripts\deploy.ps1` |
| Sync Postman env files | `make postman-sync` or `./scripts/nextpresskit postman-sync` |

## Commonly confused commands

| Pair | Difference |
|------|------------|
| `setup` vs `deploy` | `setup` bootstraps local development (deps, env, migrate, seed). `deploy` is deployment/config (nginx/systemd/TLS/release). |
| `setup` vs `install` | `install` is only modules + `.env` if missing. **`setup`** menu can run that plus other steps; profile **1** / non-interactive mode runs install, build-all, migrate-up, and seed. |
| `run` vs `start` | `run` keeps the API in the foreground. `start` runs it in the background (Unix) so your shell stays free. |
| `checks` vs `test` | `checks` is CI-style multi-check flow. `test` runs Go tests only. |

## Most-used commands

| Command | What it does | When to use it |
|---------|---------------|----------------|
| `make setup` / `./scripts/nextpresskit setup` | Text menu, or linear bootstrap when non-interactive / `NP_SETUP_NONINTERACTIVE=1`. | First run on a new clone. |
| `make run` / `./scripts/nextpresskit run` | Runs API in foreground (dev mode). | Normal local development. |
| `./scripts/nextpresskit start` | Starts API in background (Unix). | Keep server running while you use the shell. |
| `./scripts/nextpresskit stop` | Stops background API process (Unix). | Clean shutdown after `start`. |
| `./scripts/nextpresskit checks` | CI-style local checks. | Before pushing or opening a PR. |
| `make postman-sync` / `./scripts/nextpresskit postman-sync` | Creates/updates gitignored `postman/` from templates + env values. | Before importing Postman collections/environments. |
| `./scripts/nextpresskit deploy` | Interactive deployment/config generation wizard. | Generate nginx/systemd snippets or run release steps. |

## Setup, build, run

| Command | Description |
|---------|-------------|
| `./scripts/nextpresskit install` | Downloads modules and initializes local prerequisites (`.env` support included). |
| `./scripts/nextpresskit build` | Builds API binary only. |
| `./scripts/nextpresskit build-all` | Builds API + migrate + seed binaries. |
| `./scripts/nextpresskit run` | Runs API in foreground. |
| `./scripts/nextpresskit start` | Runs API in background (Unix). |
| `./scripts/nextpresskit stop` | Stops background API (Unix). |

<a id="database-and-seed-data"></a>

## Database and seed data

Tables come from the Go models via GORM AutoMigrate. `migrate-up` runs **only the persistence models for modules in the active kit list** (same rules as the API: see [`MODULES`](../.env.example) and [MODULES.md](MODULES.md)). Implementation: [`internal/platform/dbmigrate/migrate.go`](../internal/platform/dbmigrate/migrate.go) + each module’s `AutoMigrate`. There is no separate SQL migration tree: change the models, then run migrate-up.

- migrate-up: sync the database schema with the code.
- seed: add repeatable demo data (roles, sample posts, a superadmin, and more). Run this after migrate-up.
- db-fresh: local development only. Drops every table in the `public` schema (you confirm first), runs migrate-up again, and stops. It does not run seed; run seed yourself if you want the demo dataset back.

PostgreSQL must be running, with `DB_*` set in [`.env`](../.env.example) so migrate and seed can connect.

| Situation | Commands |
|-----------|----------|
| New clone, full setup | `make setup` → **1**, or `NP_SETUP_NONINTERACTIVE=1 make setup`, or `./scripts/nextpresskit migrate-up` then `seed` |
| Clean slate on your machine | `./scripts/nextpresskit db-fresh`, then `./scripts/nextpresskit seed` |
| Pulled code with model changes | `./scripts/nextpresskit migrate-up` |
| Refresh demo data only | `./scripts/nextpresskit seed` (safe to repeat) |

On Windows use `.\scripts\nextpresskit.ps1` with the same subcommand names. After `./scripts/nextpresskit build-all`, `./bin/migrate -command=up` and `./bin/seed` match migrate-up and seed.

| Subcommand | Meaning |
|------------|---------|
| migrate-up | AutoMigrate from module [`persistence`](../internal/modules) models (FK-safe order). |
| migrate-down | Removed (no versioned SQL downs). Locally use db-fresh or migrate-drop plus migrate-up. |
| migrate-version | Prints a note: there is no `schema_migrations` table with AutoMigrate. |
| migrate-drop | Drops all `public` tables after confirmation (`ALLOW_SCHEMA_DROP` is set for you). |
| db-fresh | migrate-drop plus migrate-up only; no seed. |
| seed | Upserts RBAC defaults (permissions from enabled modules) and demo data per module ([SEEDING.md](SEEDING.md)). Respects `MODULES` like migrate-up. |

On servers, releases usually run `bin/migrate -command=up` and sometimes `bin/seed` ([DEPLOYMENT.md](DEPLOYMENT.md)). Do not use db-fresh in production.

## API contract and quality checks

| Command | Description |
|---------|-------------|
| `./scripts/nextpresskit checks` | Runs project checks used in CI-style validation. |
| `./scripts/nextpresskit test` | Runs Go test suites (`go test -v ./...`). |
| `go vet ./...` | Static analysis for suspicious constructs. |
| `./scripts/nextpresskit security-check` | Runs vulnerability checks (`govulncheck`). |

## Deployment workflow commands

| Command | Description |
|---------|-------------|
| `./scripts/deploy` | Interactive deployment wizard (Linux/macOS/Git Bash). |
| `.\scripts\deploy.ps1` | Interactive deployment wizard (PowerShell). |
| `./scripts/nextpresskit deploy` | Same as `bash scripts/deploy` on Unix. |

For full production guidance, branch promotion model, and TLS options, see `DEPLOYMENT.md`.

## Postman and environments

| Command | Description |
|---------|-------------|
| `./scripts/nextpresskit postman-sync` | Syncs environment JSON into gitignored `postman/`. |
| `./scripts/nextpresskit postman-sync --dry-run` | Shows what would change, without writing files. |
| `POSTMAN_CLEAR_TOKENS=1 ./scripts/nextpresskit postman-sync` | Clears token placeholders in generated Postman env files. |

## Help and discovery

| Command | Description |
|---------|-------------|
| `./scripts/nextpresskit help` | Full command list and usage. |
| `make` | Lists the three Make targets and points to `nextpresskit help`. |

If a command is unfamiliar, run `./scripts/nextpresskit help` and pick the command family for your OS.

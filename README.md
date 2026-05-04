# NextPressKit

NextPressKit is a starter kit for building modern backend APIs using Go, Gin, and PostgreSQL.

The goal of this project is to give developers a strong starting point they can clone and build on, with common product needs already in place: authentication handling, content creation flows, and an administration area.

* Website: [nextpresskit.com](https://nextpresskit.com)
* Frontend repository: [nextpresskit/web](https://github.com/nextpresskit/web)

## Where to go next

Use this table if you're not sure which doc to open first.

| Goal | Start here |
|------|------------|
| Understand what each command does | [docs/COMMANDS.md](./docs/COMMANDS.md) |
| See every doc and how it fits together | [docs/README.md](./docs/README.md) |
| Deploy to Ubuntu (nginx, systemd, HTTPS) | [docs/DEPLOYMENT.md](./docs/DEPLOYMENT.md) |
| HTTPS + nginx on your laptop (mkcert) | [docs/deployment/local.md](./docs/deployment/local.md) · [macOS](./docs/deployment/macos.md) |
| REST request/response shapes | [docs/openapi.yaml](./docs/openapi.yaml) |
| Try endpoints in Postman | [postman-templates/README.md](./postman-templates/README.md) (`postman-sync`) |
| JWT cookies, CORS, hardening | [docs/SECURITY.md](./docs/SECURITY.md) |
| Database (migrate, seed, fresh start) | [docs/COMMANDS.md](./docs/COMMANDS.md#database-and-seed-data) · [docs/SEEDING.md](./docs/SEEDING.md) |
| Contribute or run checks before a PR | [CONTRIBUTING.md](./CONTRIBUTING.md) |

## Project Concepts

* Modular kit: enable slices via `MODULES` or edit [`internal/appregistry`](./internal/appregistry/registry.go); see [docs/MODULES.md](./docs/MODULES.md).
* Starter-first architecture for bootstrapping and customization.
* Auth-ready foundations (patterns and services for sign-in and tokens).
* Content-oriented workflows (creation and publishing APIs).
* Admin routes and structures for managing app data.
* Modern API stack: Go, Gin, GORM.

## Tech Stack

* Go
* Gin
* PostgreSQL
* GORM
* JWT
* Prometheus

## Quick start

You need Go (see `go.mod`) and PostgreSQL. Start the server, then create an empty database (and user if needed) that match `DB_*` in [.env.example](./.env.example), for example database `nextpresskit` and user `postgres` on `localhost:5432`.

On an **interactive terminal**, `./scripts/nextpresskit setup` (or `make setup`) prints a **text menu**: pick a profile (full first-time, database-only, build, CI checks), **Custom** to run other `nextpresskit` steps by number, **Add kit module** / **Remove kit module** to edit `MODULES` in `.env` (see [docs/MODULES.md](./docs/MODULES.md)). Destructive steps ask `y/N` first. For **CI or scripts** with no prompts, use `NP_SETUP_NONINTERACTIVE=1 make setup` — that runs only the linear path: `install` → `build-all` → `migrate-up` → `seed` (and optional `scripts/setup-local-https.sh` on a TTY unless `SKIP_SETUP_LOCAL_HTTPS=1`). Edit at least `JWT_SECRET` and double-check `DB_*` before migrate runs.

Quick path:

1. `make setup` or `./scripts/nextpresskit setup` (choose **1** for full first-time, or set `NP_SETUP_NONINTERACTIVE=1` to skip the menu)
2. `make run` or `./scripts/nextpresskit run`
3. Open `http://localhost:9090/health` (replace `9090` with `APP_PORT` from `.env` if you changed it). Use `/ready` if you want to confirm PostgreSQL is wired up.

More database commands (`migrate-up`, `seed`, `db-fresh`): [docs/COMMANDS.md](./docs/COMMANDS.md#database-and-seed-data) and [docs/SEEDING.md](./docs/SEEDING.md).

### Linux / macOS / Git Bash (same commands)

Copy-paste:

```bash
./scripts/nextpresskit setup
./scripts/nextpresskit run
```

Make exposes three targets; they call the same scripts as `nextpresskit`:

```bash
make setup
make run
make postman-sync
```

Full setup (menu **1** or `NP_SETUP_NONINTERACTIVE=1`) may run `scripts/setup-local-https.sh` on a TTY (unless `SKIP_SETUP_LOCAL_HTTPS=1`): mkcert, certs under `~/.local/share/nextpresskit-ssl/`, optional `/etc/hosts` hint, and on Linux with nginx may run `deploy apply-nginx --no-tls-menu`.

The API listens on `APP_PORT` (default 9090). Foreground `run` frees the port if another same-repo `bin/server` or `go run ./cmd/api` is still listening; systemd units named `nextpresskit-backend@*` are left alone (stop with `systemctl`).

### Windows (PowerShell)

From the repo root:

```powershell
.\scripts\nextpresskit.ps1 setup
.\scripts\nextpresskit.ps1 run
```

Deploy wizards: `.\scripts\nextpresskit.ps1 deploy` (PowerShell) or `bash scripts/deploy` / `./scripts/nextpresskit deploy` (Unix).

### HTTPS / Nginx locally

For HTTPS (browser cookie auth) and reverse-proxy setup, see [docs/deployment/local.md](./docs/deployment/local.md) and [docs/deployment/macos.md](./docs/deployment/macos.md). Run the deploy wizard to write snippets under `deploy/generated/`: on Linux/macOS/Git Bash use `bash scripts/deploy` (or `./scripts/nextpresskit deploy`); on Windows PowerShell use `.\scripts\deploy.ps1`.

Background mode (Unix): `./scripts/nextpresskit start` / `stop`.

## Commands (summary)

Need command-by-command explanations? Open [docs/COMMANDS.md](./docs/COMMANDS.md).

Most common confusion: `setup` is for local bootstrap; `deploy` is for deployment/config and release flows.

| Make (3 targets) | What it does |
|------------------|--------------|
| `make setup` | Text menu (profiles + custom steps) on a TTY; `NP_SETUP_NONINTERACTIVE=1` = linear install → build-all → migrate-up → seed only. |
| `make run` | API in the foreground. |
| `make postman-sync` | Refresh gitignored `postman/` from templates. |

Everything else — `install`, `build`, `migrate-up`, `seed`, `db-fresh`, `checks`, `deploy`, `start`, `stop`, and more — is available from **setup → Custom** (by step number) or directly: `./scripts/nextpresskit help`.

Running bare **`make`** prints the three targets and points to `nextpresskit help`.

Postman templates live under [postman-templates/](./postman-templates/). Run `./scripts/nextpresskit postman-sync` to create a gitignored [postman/](./postman/) workspace. Options: `--dry-run`, tier URLs such as `POSTMAN_DEV_BASE_URL`. Details: [postman-templates/README.md](./postman-templates/README.md).

## Frontend Integration

The NextPressKit backend is designed to work with the frontend web project:

* Frontend repo: [nextpresskit/web](https://github.com/nextpresskit/web)
* Backend repo: [nextpresskit/backend](https://github.com/nextpresskit/backend)
* API responsibilities include authentication, content APIs, and admin-related backend operations.
* This project can also be used separately with a different frontend or mock/local API consumers.

## API contract

This project includes an OpenAPI-first REST API contract.

* OpenAPI spec: [docs/openapi.yaml](./docs/openapi.yaml).
* REST endpoints cover auth, public content, and admin operations.
* API base path is configurable with `API_BASE_PATH` (see [.env.example](./.env.example)).

### Authentication (JWT)

* **`JWT_AUTH_SOURCE=cookie` (default):** access and refresh tokens are issued as **HttpOnly** cookies (`JWT_ACCESS_COOKIE_NAME`, `JWT_REFRESH_COOKIE_NAME`). Login and refresh return **`user`** only in JSON. Protected routes accept the access cookie or, if you switch the server to header mode, a Bearer token.
* **`JWT_AUTH_SOURCE=header`:** tokens are returned in JSON (`tokens` + `user` on login/refresh); send **`Authorization: Bearer <access_jwt>`** to protected routes.

Cross-site browser apps must set **`CORS_ORIGINS`** to the real frontend origin and use **`credentials: 'include'`**. See [docs/SECURITY.md](./docs/SECURITY.md) and [.env.example](./.env.example).

## Documentation

The single map of all docs is [docs/README.md](./docs/README.md). Skim that page whenever you feel lost.

Common links: [Module kit](./docs/MODULES.md) · [API versioning](./docs/API_VERSIONING.md) · [Seeding](./docs/SEEDING.md) · [Elasticsearch runbook](./docs/ELASTICSEARCH_OPERATIONS.md) · [Roadmap](./docs/ROADMAP.md) · [Task checklist](./docs/TODO.md) · [Changelog](./CHANGELOG.md)

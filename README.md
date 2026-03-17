## nextpress-backend

nextpress-backend is a **production-ready CMS backend** written in Go, designed as a **modular monolith** with a **global PostgreSQL database instance**, a **clean service layer**, and a future-proof **plugin system** (planned for later phases).

This repository currently contains **Phase 1 – Project Foundation**:

- Strict, opinionated folder structure
- Application bootstrap with Gin
- Environment and configuration loading
- Deployment-oriented scripts and configuration skeletons

### High-level architecture

- **Modular Monolith**: Features are grouped by domain under `internal/modules`, sharing a single process and database, while keeping clear boundaries for maintainability.
- **Global DB instance**: A single Postgres connection (via GORM) is initialized once and injected where needed to avoid duplicated connection logic.
- **Service layer**: Thin HTTP handlers, reusable application services, and clear separation of concerns without over-engineered DDD layers.

### Folder structure (Phase 1)

```text
.
├── cmd/
│   └── nextpress/
│       └── main.go          # Application entrypoint
├── internal/
│   ├── config/              # Configuration loading and environment helpers
│   ├── http/                # HTTP server setup (Gin engine, middlewares, routing)
│   ├── db/                  # Global DB initialization and lifecycle management
│   ├── modules/             # Feature modules (content, auth, media, etc.)
│   ├── logging/             # Centralized logger setup
│   └── shared/              # Shared helpers (errors, responses, etc.)
├── deployments/
│   ├── docker/
│   │   ├── Dockerfile       # Production Docker image
│   │   └── docker-compose.local.yml  # Local dev + Postgres
│   └── configs/
│       ├── app.local.yaml   # Local development defaults
│       ├── app.dev.yaml     # Remote dev/staging defaults
│       └── app.prod.yaml    # Production defaults
├── scripts/
│   ├── run_local.sh         # Local run helper
│   ├── run_dev.sh           # Dev/staging run helper
│   └── run_prod.sh          # Production entrypoint (for Docker / systemd)
└── Makefile                 # Common developer tasks (build, run, test)
```

> Later phases will populate `internal/modules` with concrete CMS features (auth, content, media, plugins, etc.).

### Getting started (Phase 1)

#### Prerequisites

- Go 1.26+
- Docker & Docker Compose (for local Postgres)

#### Local development (with Docker)

```bash
make up-local       # Start Postgres via docker-compose
make run-local      # Run API locally with hot reload (once added) or plain go run
```

#### Configuration

- `.env` (optional) – for local overrides, loaded via `godotenv`.
- `deployments/configs/app.<env>.yaml` – environment-specific defaults.

Key environment variables (Phase 1):

- `APP_ENV` – one of `local`, `dev`, `prod` (default: `local`)
- `APP_HTTP_PORT` – port for the HTTP server (default: `8080`)

Database, JWT, and module-level configuration will be introduced in later phases.

# NextPress Backend

NextPress is a modular CMS backend written in Go.

## Stack

- Go 1.26
- Gin HTTP Framework
- PostgreSQL
- GORM
- Zap Logger
- JWT Authentication

## Architecture

The project follows:

- Clean Architecture
- Modular Monolith
- Domain Driven Design

## Project Structure

cmd/api → application entry point

internal/config → configuration system  
internal/platform → shared infrastructure  
internal/modules → domain modules

pkg → reusable utilities

## Development

Run server:

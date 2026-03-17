# Local

Run nextpress-backend on your machine for development or testing. No systemd or Nginx; you run the binary from the project folder.

| Item   | Value                    |
|--------|--------------------------|
| Folder | project root (any path)  |
| Branch | your branch              |

---

## Prerequisites

| Requirement   | Details |
|---------------|---------|
| **Go**        | Version in `go.mod`. The project will not build with an older or incompatible version. |
| **PostgreSQL**| Installed and running (for future phases). Configure `DB_*` variables in `.env` when DB is used. |
| **Git**       | To clone the repository. |

---

## 1. Clone and setup

Clone the repo, download dependencies, and create the environment file:

```bash
git clone <repo-url> nextpress-backend
cd nextpress-backend
go mod download
cp .env.example .env
```

Edit `.env` as needed; at minimum you can leave defaults for local runs:

- `APP_NAME` – name for logs and metrics.
- `APP_ENV` – usually `development` locally.
- `APP_PORT` – HTTP port (default `9090`).

---

## 2. Run the API

Using the helper script:

```bash
./scripts/run_local.sh
```

Or directly with Go:

```bash
APP_ENV=development APP_PORT=9090 go run ./cmd/api
```

The server listens on `APP_PORT` (e.g. `http://localhost:9090`).

---

[← Menu](../DEPLOYMENT.md)


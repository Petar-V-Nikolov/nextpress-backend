# nextpress-backend – Deployment

Deploy the CMS backend to Ubuntu using **`scripts/deploy`**, Nginx, and systemd. Each environment has its own guide under `docs/deployment/`.

Each environment (production, staging, dev) can run on a **separate server** or share a server; you configure folder, `APP_PORT`, and database per server via `.env`.

---

## Menu

| Environment | Guide | Branch | Folder |
|------------|-------|--------|--------|
| Production | [deployment/production.md](deployment/production.md) | `main`    | `/var/www/nextpress-backend-production` |
| Staging    | [deployment/staging.md](deployment/staging.md)       | `staging` | `/var/www/nextpress-backend-staging`    |
| Dev        | [deployment/dev.md](deployment/dev.md)               | `dev`     | `/var/www/nextpress-backend-dev`        |
| Local      | [deployment/local.md](deployment/local.md)           | —         | project root                            |

Port is set in `.env` on each server (e.g. 9090, 9091, 9092; if multiple envs share one server, use different ports).

Branch flow (recommended): work flows **dev → staging → main**. Promote by merging and pushing between branches as needed. Step-by-step commands and optional feature/hotfix flows: **[GIT_FLOW.md](GIT_FLOW.md)**.

---

## Overview

| Item            | Description |
|-----------------|-------------|
| **Deploy script** | `./scripts/deploy [production\|staging\|dev]` — run from the **root of that environment's folder** |
| **Binary**        | Builds **`bin/server`** from `cmd/api`; script runs **`migrate -command=up`**; optional **`bin/seed`** if `RUN_SEED_ON_DEPLOY=true` |
| **Systemd**       | Template unit `nextpress-backend@.service`; instances: `nextpress-backend@production`, `nextpress-backend@staging`, `nextpress-backend@dev` |

---

## Prerequisites (server environments)

- Ubuntu 22.04 LTS (or similar)
- Go (same major version as in `go.mod`)
- PostgreSQL (accessible from the server)
- Nginx, Git, Make, systemd
- SSH key and GitHub access for the deploy user

---

## One-time: server layout

On **each server** where you deploy, create the folder and clone once (one environment per server). If you run several environments on the **same** server, create one folder per environment and set a different `APP_PORT` in each `.env`.

Example (one env per server):

```bash
sudo mkdir -p /var/www
sudo chown "$USER" /var/www

git clone <repo-url> /var/www/nextpress-backend-production
```

Example (all three envs on one server):

```bash
sudo mkdir -p /var/www
sudo chown "$USER" /var/www

git clone <repo-url> /var/www/nextpress-backend-production
git clone <repo-url> /var/www/nextpress-backend-staging
git clone <repo-url> /var/www/nextpress-backend-dev
```

Then follow the guide for that environment: [production](deployment/production.md), [staging](deployment/staging.md), [dev](deployment/dev.md). For local development see [local](deployment/local.md).


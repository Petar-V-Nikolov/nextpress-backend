# Commands Guide

Simple command reference for day-to-day work.

[← Docs index](README.md)

## Make commands

- `make setup`: first-time setup
- `make run`: run API
- `make postman-sync`: generate/update local Postman files

Direct Make aliases are also available for common script commands:

- `make migrate-up`
- `make seed`
- `make db-fresh`
- `make build`, `make build-all`
- `make test`, `make checks`

Under the hood, these map to `./scripts/nextpresskit <command>`.

<a id="database-and-seed-data"></a>

## Database and seed data

- `migrate-up`: create/update tables from models
- `seed`: load demo data (safe to rerun)
- `db-fresh`: drop all public tables, then run `migrate-up` (local only)

Typical flows:

```bash
make migrate-up
make seed
```

```bash
make db-fresh
make seed
```

Never use `db-fresh` in production.

## Quality checks

- `./scripts/nextpresskit checks`: CI-style checks
- `./scripts/nextpresskit test`: Go tests
- `go vet ./...`: static checks
- `./scripts/nextpresskit security-check`: vulnerability scan

## Deployment commands

- `./scripts/nextpresskit deploy`
- `./scripts/deploy` (Unix)
- `.\scripts\deploy.ps1` (PowerShell)

Full server instructions: [`DEPLOYMENT.md`](DEPLOYMENT.md).

## Postman commands

- `./scripts/nextpresskit postman-sync`
- `./scripts/nextpresskit postman-sync --dry-run`
- `POSTMAN_CLEAR_TOKENS=1 ./scripts/nextpresskit postman-sync`

## Help

- `./scripts/nextpresskit help`
- `make` (prints primary Make targets)

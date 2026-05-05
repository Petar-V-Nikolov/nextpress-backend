# Postman - NextPressKit API

[Docs index](../docs/README.md) · [Commands](../docs/COMMANDS.md)

Templates are versioned in `postman-templates/`.
Generated local files go to gitignored `postman/`.

## Quick Start

```bash
./scripts/nextpresskit postman-sync
```

Then import JSON files from `postman/` into Postman.

## Collections

- Public collection: health/auth/public content routes
- Admin collection: `/admin/*` routes

### `jwt_auth_source`

Matches server behavior controlled by `JWT_AUTH_SOURCE` in `.env`:

| Value | Meaning |
|-------|---------|
| `cookie` (default) | After `POST /auth/login`, Postman stores HttpOnly cookies for `{{base_url}}`. Protected requests **do not** send `Authorization`; the collection pre-request script removes that header so the cookie jar is used. |
| `header` | Login/refresh responses include `tokens` in JSON. The collection scripts set `Authorization: Bearer …` from `access_token` (Public) or `admin_access_token` (Admin). |

Set this variable in each imported environment.

## Collections

| Collection | File | Contents |
|------------|------|----------|
| **NextPressKit Public API** | `NextPressKit-Public-API.postman_collection.json` | Root/health/ready endpoints plus `/auth/*` and public content APIs. |
| **NextPressKit Admin API** | `NextPressKit-Admin-API.postman_collection.json` | All `/admin/*` endpoints requiring admin token and permissions. |

## Environments

Use one environment per target. Both collections rely on `{{base_url}}`. **`POST /auth/login`** in the Public collection uses `{{superadmin_email}}` and `{{superadmin_password}}`, which default to the seed superadmin (`SEED_SUPERADMIN_EMAIL` / `SEED_SUPERADMIN_PASSWORD` in `.env.example`). Override per environment (required for real staging/production accounts).

| Environment | File | Use case | `base_url` |
|-------------|------|----------|------------|
| **NextPressKit - Local** | `NextPressKit-Local.postman_environment.json` | Local Nginx + TLS (`bash scripts/deploy` or `./scripts/nextpresskit deploy`, `nextpresskit.local` in `/etc/hosts`) | `https://nextpresskit.local` |
| **NextPressKit - Dev** | `NextPressKit-Dev.postman_environment.json` | Dev deployment | `https://api-dev.example.com` |
| **NextPressKit - Staging** | `NextPressKit-Staging.postman_environment.json` | Staging deployment | `https://api-staging.example.com` |
| **NextPressKit - Production** | `NextPressKit-Production.postman_environment.json` | Production deployment | `https://api.example.com` |

> Replace the dev/staging/production `base_url` values with your actual domains. For **direct** `go run` / `make run` without Nginx, set local `base_url` to `http://127.0.0.1:9090` (or your `APP_PORT`).

### Sync from repo env files

Refresh **`postman/*.postman_environment.json`** from `.env.example` and `.env` (and optional shell overrides). On a fresh clone, this also creates **`postman/`** from these templates when needed:

```bash
./scripts/nextpresskit postman-sync
# or
make postman-sync
```

- Preview: `./scripts/nextpresskit postman-sync --dry-run`
- Windows: `.\scripts\nextpresskit.ps1 postman-sync`
- Tier URLs: `POSTMAN_LOCAL_BASE_URL`, `POSTMAN_DEV_BASE_URL`, `POSTMAN_STAGING_BASE_URL`, `POSTMAN_PRODUCTION_BASE_URL`, or set `NEXTPRESS_PUBLIC_HOST` for local `https://<host>`
- Clear token placeholders in the JSON: `POSTMAN_CLEAR_TOKENS=1 ./scripts/nextpresskit postman-sync`

Collections are not rewritten (requests use `{{base_url}}` only).

### Usage

1. Run **`postman-sync`** once so **`postman/`** contains the JSON (then import from that folder).
2. Import the two collections and the four environment files into Postman.
3. Select one environment.
4. Run **`POST /auth/login`** from the Public collection.
   - **`jwt_auth_source=cookie`:** cookies are stored automatically; response body is `{ "user": … }` only. Then run Admin requests against the same `base_url`.
   - **`jwt_auth_source=header`:** the login tests store `access_token`, `refresh_token`, and **`admin_access_token`** (copy of access) for the Admin collection script.
5. Cookie mode uses cookie jar automatically. Header mode uses bearer tokens.

### Notes

- `POST /admin/bootstrap/claim-admin` is only available when `RBAC_BOOTSTRAP_ENABLED=true`.
- `GET /posts/search` and `POST /admin/posts/search/reindex` require Elasticsearch to be enabled.

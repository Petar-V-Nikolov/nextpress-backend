# Postman - NextPress Backend API

Postman collections and environments for the NextPress backend.

## Route groups (Public/Auth and Admin)

The API is split into two major groups:

| Type | Base URL | Auth | Use case |
|------|----------|------|----------|
| **Public/Auth** | `{{base_url}}` | Public routes: none. Auth routes: none. | Health/readiness checks, auth (`/auth/*`), and public content routes (`/posts/*`, `/pages/*`, `/menus/*`). |
| **Admin** | `{{base_url}}` | JWT (`Authorization: Bearer {{admin_access_token}}`) | CMS and management routes under `/admin/*`: posts, pages, taxonomy, media, menus, RBAC, plugin management, and bootstrap/admin checks. |

## Collections

| Collection | File | Contents |
|------------|------|----------|
| **NextPress Public API** | `NextPress-Public-API.postman_collection.json` | Root/health/ready endpoints plus `/auth/*` and public content APIs. |
| **NextPress Admin API** | `NextPress-Admin-API.postman_collection.json` | All `/admin/*` endpoints requiring admin token and permissions. |

## Environments

Use one environment per target. Both collections rely on `{{base_url}}`.

| Environment | File | Use case | `base_url` |
|-------------|------|----------|------------|
| **NextPress - Local** | `NextPress-Local.postman_environment.json` | Local development (`APP_PORT=9090` by default) | `http://localhost:9090` |
| **NextPress - Dev** | `NextPress-Dev.postman_environment.json` | Dev deployment | `https://api-dev.example.com` |
| **NextPress - Staging** | `NextPress-Staging.postman_environment.json` | Staging deployment | `https://api-staging.example.com` |
| **NextPress - Production** | `NextPress-Production.postman_environment.json` | Production deployment | `https://api.example.com` |

> Replace the dev/staging/production `base_url` values with your actual domains.

### Setup

1. Import the two collections and the four environment files into Postman.
2. Select one environment.
3. Run `POST /auth/login` from Public collection to auto-store `access_token` and `refresh_token`.
4. Copy `access_token` to `admin_access_token` for admin requests.

### Notes

- `POST /admin/bootstrap/claim-admin` is only available when `RBAC_BOOTSTRAP_ENABLED=true`.
- `GET /posts/search` and `POST /admin/posts/search/reindex` require Elasticsearch to be enabled.
- GraphQL (`/graphql`) is optional and controlled by `GRAPHQL_ENABLED`; it is not part of these REST collections.

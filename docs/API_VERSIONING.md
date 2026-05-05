# API Versioning Strategy

[← Docs index](README.md) · [OpenAPI](openapi.yaml)

This document records the current API versioning decision for NextPressKit and how to evolve it safely.

## Summary

API routes are unversioned by default.
When needed, set `API_BASE_PATH=/v1`.

## Current Decision

- Default strategy: **unversioned paths**.
- Optional path prefix: **`API_BASE_PATH`**.
- Default value: empty (`""`) so routes are unversioned.

Examples:

- `API_BASE_PATH=""` -> `POST /auth/login`, `GET /posts`, `GET /admin/ping`
- `API_BASE_PATH="/v1"` -> `POST /v1/auth/login`, `GET /v1/posts`, `GET /v1/admin/ping`

This keeps the API simple now while making URL-path versioning a no-refactor config change later.

## Why

- Simple now
- Easy to move to `/v1` later
- No handler refactor required

## Runtime Behavior

- All REST route groups are mounted under `API_BASE_PATH`.

## Configuration

In `.env`:

```env
# Leave empty for unversioned endpoints.
API_BASE_PATH=
```

Rules:

- Empty or `/` means no prefix.
- Missing leading slash is normalized (`v1` -> `/v1`).
- Trailing slash is trimmed (`/v1/` -> `/v1`).

## Move To URL Versioning Later

1. Set `API_BASE_PATH=/v1`.
2. Update external clients to call `/v1/*`.
3. Keep unversioned compatibility only if needed (temporary rewrite/proxy rule).
4. Announce migration window and cutoff date.

No handler-level route refactor is required.

## Header Versioning Option

Possible, but path versioning is usually easier to debug and cache.

## Compatibility Policy (Baseline)

Until formal versioning is introduced:

- Additive changes only by default (new optional fields/endpoints).
- Avoid breaking field renames/removals without a migration plan.
- Document contract-affecting changes in `CHANGELOG.md`.

## Deprecation Policy Template

When versioning starts, apply this baseline:

- Announce deprecation with timeline.
- Provide migration guide and examples.
- Track old/new version usage.
- Remove old behavior only after the published sunset date.

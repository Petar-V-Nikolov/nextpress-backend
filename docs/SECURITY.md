# Security and hardening

[← Docs index](README.md) · [`.env.example`](../.env.example)

Baseline security checklist for NextPressKit.

## Dependency and CVE review

Run regularly (weekly and before releases):

```bash
./scripts/nextpresskit security-check
go list -m -u all
```

- `security-check` runs `govulncheck`.
- Upgrade vulnerable dependencies quickly.

## CORS policy

`CORS_ORIGINS` controls allowed origins. Keep it explicit in production:

- **Production:** set exact origins (comma-separated), e.g. `https://app.example.com,https://admin.example.com`.
- **Local/dev:** empty `CORS_ORIGINS` allows all origins for faster iteration.

If `CORS_ORIGINS` is empty, browser credentialed cross-origin requests will not work.

## JWT mode

Set `JWT_AUTH_SOURCE`:

| Mode | Access token | Refresh token | Login/refresh JSON |
|------|----------------|---------------|---------------------|
| `cookie` (default) | `JWT_ACCESS_COOKIE_NAME` (default `access_token`) | `JWT_REFRESH_COOKIE_NAME` (default `refresh_token`) | `user` only; tokens are not returned in the body |
| `header` | Client sends `Authorization: Bearer <jwt>` | Client stores refresh from JSON and sends it on `POST /auth/refresh` | `tokens` + `user` |

- `cookie` for browser sessions
- `header` for Bearer token clients

Cross-site browser cookies require:
- HTTPS
- Exact `CORS_ORIGINS`
- `credentials: include` from frontend

## Rate limits

- Start with defaults from `.env.example`.
- Keep different limits for public/auth/admin paths.
- Tune using real traffic and logs.

## JWT secret rotation

Current implementation uses one signing secret (`JWT_SECRET`) for access and refresh tokens.

Simple rotation flow:
1. Announce maintenance window for token invalidation.
2. Set new `JWT_SECRET` and redeploy API instances.
3. Invalidate old refresh tokens (users re-authenticate).
4. Confirm login + refresh works.


# Troubleshooting

[← Docs index](README.md) · [Commands](COMMANDS.md)

Common local issues and quick fixes.

## API Does Not Start

- Check logs in your terminal.
- Verify `.env` exists and has valid values.
- Confirm PostgreSQL is running.
- Run:

```bash
./scripts/nextpresskit migrate-up
./scripts/nextpresskit run
```

## Database Connection Errors

- Recheck `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASSWORD`, `DB_NAME`.
- Make sure the database exists.
- Test with your SQL client using same credentials.

## Port Already In Use

- Change `APP_PORT` in `.env`, or stop the other process.
- Try `make run` again.

## Login Works But Browser Is Not Authenticated

- If using cookie mode, use HTTPS.
- Set `CORS_ORIGINS` to your frontend origin exactly.
- Ensure frontend sends requests with credentials enabled.

## Admin Endpoints Return Forbidden

- Login with an admin/superadmin account.
- Run seed again to restore default roles and permissions:

```bash
./scripts/nextpresskit seed
```

## Elasticsearch Endpoints Return 501

- Set `ELASTICSEARCH_ENABLED=true`.
- Set valid `ELASTICSEARCH_URLS`.
- Restart API.

## Need A Clean Local Reset

```bash
./scripts/nextpresskit db-fresh
./scripts/nextpresskit seed
```

Use only on local/dev databases.

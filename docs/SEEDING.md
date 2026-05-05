# Database seeding

[← Docs index](README.md) · [Commands](COMMANDS.md#database-and-seed-data)

Seeding loads demo data so you can use the API right away.
Run migrations first.

```bash
./scripts/nextpresskit migrate-up && ./scripts/nextpresskit seed   # usual path
./scripts/nextpresskit db-fresh && ./scripts/nextpresskit seed      # wipe public schema, recreate tables, then seed (dev only)
./scripts/nextpresskit seed                                         # run again anytime; upserts, not duplicates
```

## What gets seeded

- RBAC defaults (roles/permissions)
- Superadmin user
- Demo rows for enabled modules

Reruns are safe: seeders use upserts so you do not pile up duplicate keys.

## Superadmin credentials

Configure in `.env` (defaults shown in `.env.example`):

```bash
SEED_SUPERADMIN_EMAIL=superadmin@nextpresskit.local
SEED_SUPERADMIN_PASSWORD=SuperAdmin123!
```

The seeded superadmin user is deterministic and updated on reruns (same identity, latest configured credentials).

## Notes

- Rerunning `seed` is safe (upsert behavior)
- `MODULES` in `.env` decides which module seeds run
- Use `db-fresh` only on local/dev databases

## Deploy

If you want the deploy wizard release step to run `./bin/seed` automatically, set `RUN_SEED_ON_DEPLOY=true` in `.env` ([DEPLOYMENT.md](DEPLOYMENT.md)).

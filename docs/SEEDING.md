# Database seeding

[← Documentation index](README.md) · [Commands: migrate and seed](COMMANDS.md#database-and-seed-data) · [Module kit](MODULES.md)

Seeding adds predictable demo content so you can try the API without entering everything by hand. Tables must already exist: run migrate-up first, or use setup, which does migrate and seed. All command options (including db-fresh) are in [COMMANDS.md](COMMANDS.md#database-and-seed-data).

**`MODULES`** in `.env` filters which modules run their `Seed` step (same resolution as `cmd/api` and `cmd/migrate`). Empty = full default registry in [`internal/appregistry`](../internal/appregistry/registry.go).

```bash
./scripts/nextpresskit migrate-up && ./scripts/nextpresskit seed   # usual path
./scripts/nextpresskit db-fresh && ./scripts/nextpresskit seed      # wipe public schema, recreate tables, then seed (dev only)
./scripts/nextpresskit seed                                         # run again anytime; upserts, not duplicates
```

You can also run `go run ./cmd/seed`, `./bin/seed` after `./scripts/nextpresskit build-all`, or `./scripts/nextpresskit seed`; behavior is the same.

## What gets seeded

1. **RBAC defaults** (`pkg/seed/rbac_defaults.go`): admin role, permission rows for codes returned by each enabled module’s `Permissions()`, and admin↔permission links for those codes only.
2. **Per-module demo data** (`internal/modules/*/persistence/seed_demo.go`): users, RBAC demo roles/extra permissions, taxonomy, media, posts (+ relations), pages — each module seeds only when enabled.
3. **Superadmin:** one privileged user tied to both superadmin and admin roles (from user + RBAC seed steps).

Reruns are safe: seeders use upserts so you do not pile up duplicate keys.

## Superadmin credentials

Configure in `.env` (defaults shown in `.env.example`):

```bash
SEED_SUPERADMIN_EMAIL=superadmin@nextpresskit.local
SEED_SUPERADMIN_PASSWORD=SuperAdmin123!
```

The seeded superadmin user is deterministic and updated on reruns (same identity, latest configured credentials).

## Tables seeded with ~100 rows (full default modules)

- `users` (includes `superadmin` as one of the 100)
- `roles`
- `permissions` (RBAC defaults for enabled modules + generated seed permissions)
- `role_permissions`
- `user_roles`
- `posts`, `pages`
- `categories`, `tags`
- `media`
- `post_categories`, `post_tags`
- `post_seo`, `post_metrics`
- `series`, `post_series`
- `post_coauthors`, `post_gallery_items`, `post_changelog`, `post_syndication`
- `translation_groups`, `post_translations`

## Deploy

If you want the deploy wizard release step to run `./bin/seed` automatically, set `RUN_SEED_ON_DEPLOY=true` in `.env` ([DEPLOYMENT.md](DEPLOYMENT.md)).

---

See also: [Documentation index](README.md) · [Deployment](DEPLOYMENT.md) · [TODO](TODO.md)

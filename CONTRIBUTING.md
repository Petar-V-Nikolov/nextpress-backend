# Contributing

[Docs index](docs/README.md) · [Commands](docs/COMMANDS.md)

Keep changes small, tested, and documented.

## Before Opening A PR

Run:

```bash
./scripts/nextpresskit checks
```

If needed, run only the basics:

```bash
./scripts/nextpresskit test
go vet ./...
```

Fix failures or explain them in the PR description.

## Keep Docs Updated

Update docs in the same PR as code changes.

- API route or payload changed: update `docs/openapi.yaml`
- Command behavior changed: update `docs/COMMANDS.md`
- Seed/demo data changed: update `docs/SEEDING.md`
- Deploy flow changed: update `docs/DEPLOYMENT.md`
- Scope or priorities changed: update `docs/TODO.md` and `docs/ROADMAP.md`

## Branch And Environment Model

See `docs/DEPLOYMENT.md` for `dev -> staging -> main` flow and server layout.

## Need Help?

Open an issue or discussion in the repository.

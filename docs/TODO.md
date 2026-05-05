# TODO (maintenance checklist)

`[x]` done · `[ ]` open.

---

## Modular composition

- [x] `kit.Module` + `kit.Deps` contract ([`internal/kit`](../internal/kit/module.go))
- [x] `internal/app.Run` composition root + phased HTTP registration
- [x] `internal/appregistry` compile-time default module list
- [x] `MODULES` env filtering + implicit dependencies ([`internal/kit/manifest.go`](../internal/kit/manifest.go))
- [x] `cmd/migrate` uses same resolved module list as API
- [x] `cmd/seed` uses merged `Permissions()` + per-module `Seed`

---

## Plugins removal

- [x] Remove `internal/modules/plugins` and admin plugin routes
- [x] Post save chain: derived fields + optional ES only (no DB-driven plugin hooks)
- [x] OpenAPI + Postman: drop `/admin/plugins`
- [x] RBAC: drop `plugins:manage` from default seed / `SeedRBACDefaults` known codes

---

## Documentation

- [x] [`MODULES.md`](MODULES.md) — how to enable/disable and add modules
- [x] [ADR 0001](adr/0001-module-composition.md)
- [ ] Keep [openapi.yaml](openapi.yaml) in sync whenever REST changes
- [x] Add troubleshooting guide for common setup/run issues

---

## Testing & quality

- [ ] Unit tests for `internal/app` wiring (smoke: minimal `MODULES` boots)
- [ ] **Unit tests** for transport layers across modules (ongoing)
- [x] Integration tests with Postgres (CI)
- [x] CI workflow (`go test`, `go vet`, lint)

---

## Future (not in core kit)

- [ ] Optional **object storage** backend for `Storage` (S3-compatible)
- [ ] **Separate kit or service** for dynamic plugins / marketplace-style extensions
- [ ] Example “product presets” (blog vs admin-only API) as docs or tiny `cmd/*` samples

---

## Security & ops

- [x] Periodic dependency / CVE review
- [x] Documented CORS, rate limits, JWT modes
- [ ] Revisit JWT rotation / multi-key when product needs it

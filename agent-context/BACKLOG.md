# BACKLOG — provider-gridscale (private lane detail)

Public build order: [docs/ROADMAP.md](../docs/ROADMAP.md). OpenSpec changes: `openspec/changes/`.

## Status

| Phase | State |
| --- | --- |
| Codegen (E6g-S01 in kaddy's terms) | **complete** — 32 resources generated & building |
| Design (this harness: openspec, ADRs, ROADMAP, backlog) | **in progress** — E1 lane |
| Implementation E1–E6 | **pending** — `/agent-loop` |
| E4-S04 (docs site), E5-S06 (signing), E6-S05 (assurance) | **decision-gated / stretch** |

## Lane dependency graph

```
E1 (foundation)
 ├─ E2 tests        (parallel after E1)
 ├─ E3 branding     (parallel after E1)
 ├─ E4 docs         (parallel after E1; E4-S01/02 need no assets; E4-S04 gated on D-002)
 ├─ E6 governance   (parallel after E1)
 └─ E5 CI/supply    (after E2 coverage lands + E3 metadata/badges; E5-S04 anytime)
```

Disjoint path sets (safe to run as parallel worktrees):
- **E2** → `config/*_test.go`, `examples/**`, `test/**`, Makefile test targets
- **E3** → `extensions/**`, `docs/assets/**`, `README.md`, `package/crossplane.yaml`
- **E4** → `docs/api/**`, `docs/adr/**`, `.crd-ref-docs.yaml`, `mkdocs.yml`
- **E6** → `CONTRIBUTING.md`, `SECURITY.md`, `GOVERNANCE.md`, `CODEOWNERS`, `OWNERS.md`, `.github/*_TEMPLATE*`
- **E5** → `.github/workflows/**`, `.pre-commit-config.yaml`, `codecov.yml`, `cliff.toml`, `.github/gitleaks.toml`

Collision watch: **E3-S02 (README)** and **E3-S04 (badges)** both touch `README.md` — sequence them.
**E5** edits `.github/workflows/` which the stock upjet template owns; prefer *adding* workflow files
over editing `ci.yml` where possible.

## Story ID scheme

`E<n>-S<nn>` — e.g. `E2-S01`, `E5-S06`. REQ IDs: `REQ-E<n>-S<nn>-<mm>` (see `openspec/config.yaml`).
OpenSpec change slug: `e2-test-foundation`, `e5-ci-supplychain`, … (lowercase, hyphenated).

> **Numbering note (D-001):** this repo uses a *fresh, repo-local* E1–E6 scheme. Kaddy tracks this
> whole provider as its epic **E6g** — that is kaddy's external handle for us and does not appear in
> this repo's IDs (workspace rule: repos never cross-reference). The one legacy touchpoint is the
> existing commit `…(E6g-S01)` for the initial generation; we do not renumber history.

## Provider-specific hazards (carry into every lane)

- **Never hand-edit generated code**: `apis/**/zz_*.go`, `internal/controller/**`, `package/crds/**`.
  Route all changes through `config/` or codegen inputs. `make reviewable` / `make check-diff` guard this.
- **BSL terraform pin**: `TERRAFORM_VERSION=1.5.7` is deliberate (BSL after that). Don't bump it.
- **Kind overrides exist**: `config/gridscale.go` overrides mysql8 (avoids Go kind "0") and
  storageimport (avoids Go keyword `Import`) — E2-S02 tests must assert these specifically.
- **golangci local-prefix is stale**: points at `upjet-provider-template`, not
  `PlatformRelay/provider-gridscale` — worth a fix (belongs in E5-S03 or a tidy commit).
- **uptest needs real creds**: E-level tests hit the live gridscale API; keep them behind
  `/test-examples` + nightly, never on every push. Document the creds contract in E2-S05.

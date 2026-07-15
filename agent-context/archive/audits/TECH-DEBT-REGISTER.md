# TECH-DEBT-REGISTER — provider-gridscale

First established: 2026-07-15 (baseline audit; no prior register). Audited against `origin/main` @ `00e1c6c`.
Severity: P0 (blocker) · P1 (high) · P2 (medium) · P3 (low). Status: NEW / open / fixed / regressed.

| ID | Dimension | Sev | Description | Evidence (path:line) | Est. effort | First-seen | Status |
|----|-----------|-----|-------------|----------------------|-------------|-----------|--------|
| CRED-1 | Security | P0 | Credentials extracted then discarded — `ps.Configuration` assignment commented out; provider never authenticates | `internal/clients/gridscale.go:53-57` | S | 2026-07-15 | NEW |
| CRED-2 | Security | P0 | Wrong credential keys `username`/`password`; gridscale expects `uuid`/`token` (schema.json provider block) | `internal/clients/gridscale.go:54-57`; `examples/cluster/providerconfig/secret.yaml.tmpl`; `examples/namespaced/providerconfig/secret.yaml.tmpl` | S | 2026-07-15 | NEW |
| TEST-1 | Tests | P1 | Zero tests in repo; `make test` gate passes vacuously; no coverage floor | `.github/workflows/ci.yml:186-193` | M | 2026-07-15 | NEW |
| ARCH-1 | Architecture | P1 | No cross-resource references; generated examples carry unresolved `${...}` HCL interpolations | `config/` (none); `examples-generated/cluster/ssl/v1alpha1/certificate.yaml` | M | 2026-07-15 | NEW |
| DOC-1 | Docs | P1 | ProviderConfig example uses wrong group `gridscale.crossplane.io` vs actual `gridscale.platformrelay.io` | `examples/cluster/providerconfig/providerconfig.yaml:1` vs `apis/cluster/v1beta1/register.go:12` | S | 2026-07-15 | NEW |
| ARCH-2 | Architecture | P2 | provider-metadata incomplete/malformed: 31 keyed vs 32 resources; ISO-image keyed by title `ISO Image:` with corrupted fields | `config/provider-metadata.yaml:3,10` | S | 2026-07-15 | NEW |
| TEST-2 | Tests | P2 | e2e wired but non-blocking (comment-gated) and unrunnable today due to CRED-1/ARCH-1 | `.github/workflows/e2e.yaml:96-149` | M | 2026-07-15 | NEW |
| SEC-1 | Security | P2 | golangci suppresses gosec G101 & G104 globally (not just tests) | `.golangci.yml:102-107` | S | 2026-07-15 | NEW |
| DOC-2 | Docs | P2 | Generated examples cover only 6 of ~20 resource groups | `examples-generated/` | M | 2026-07-15 | NEW |
| DOC-3 | Docs | P2 | README is unedited upjet template; broken placeholder text, no gridscale setup | `README.md:10` | S | 2026-07-15 | NEW |
| DIR-2 | Direction | P2 | Scaffold leftovers: `provider-upjet-aws` changelog label, `upjet-provider-template` goimports prefix, 3-way Go version drift | `cmd/provider/main.go:211`; `.golangci.yml:126`; `Makefile:47`/`ci.yml:13`/`e2e.yaml:12` | S | 2026-07-15 | NEW |
| ARCH-3 | Architecture | P3 | Verify UUID external-name is right for name-identified resources (bucket, storage_import, clone) | `config/external_name.go:19-31` | S | 2026-07-15 | NEW |
| TEST-3 | Tests | P3 | `schema-version-diff` gate dormant (`if: inputs.upjet-based-provider` never set) | `.github/workflows/ci.yml:53-56` | S | 2026-07-15 | NEW |
| DOC-4 | Docs | P3 | CODEOWNERS/OWNERS.md are placeholders; no real maintainer | `CODEOWNERS:16`; `OWNERS.md:9` | S | 2026-07-15 | NEW |
| DIR-3 | Direction | P3 | local-deploy download URL wrong (releases.hashicorp.com) — prior 404 | `Makefile:17` | S | 2026-07-15 | **FIXED** (commit `797adad`, GitHub releases) |
| ~~DIR-1~~ | Direction | — | WITHDRAWN — GUIDELINES.md/BACKLOG.md/decisions.md/ROADMAP exist on main; based on stale local checkout | `agent-context/GUIDELINES.md:1` | — | 2026-07-15 | WITHDRAWN |

## Summary
- **Open:** 14 · **Fixed:** 1 (DIR-3) · **Withdrawn:** 1 (DIR-1) · **Regressed:** 0 (baseline — nothing to compare against)
- **Open by severity:** P0 × 2 · P1 × 3 · P2 × 6 · P3 × 3

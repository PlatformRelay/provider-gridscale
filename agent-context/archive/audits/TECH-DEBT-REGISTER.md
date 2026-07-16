# TECH-DEBT-REGISTER — provider-gridscale

First established: 2026-07-15 (baseline audit; no prior register). Audited against `origin/main` @ `00e1c6c`.
Severity: P0 (blocker) · P1 (high) · P2 (medium) · P3 (low). Status: NEW / open / fixed / regressed.

| ID | Dimension | Sev | Description | Evidence (path:line) | Est. effort | First-seen | Status |
|----|-----------|-----|-------------|----------------------|-------------|-----------|--------|
| CRED-1 | Security | P0 | Credentials extracted then discarded — `ps.Configuration` assignment commented out; provider never authenticates. **Wiring landed `b16b5a2`; live-validated 2026-07-15** (out-of-cluster smoke test — real VM create/delete, verified via gridscale API) | `internal/clients/gridscale.go:53-57` | S | 2026-07-15 | **FIXED** (`b16b5a2`, live-validated 2026-07-15) |
| CRED-2 | Security | P0 | Wrong credential keys `username`/`password`; gridscale expects `uuid`/`token` (schema.json provider block). **Fixed + secret templates corrected `b16b5a2`; live-validated 2026-07-15** (authenticated to live API, real VM create/delete) | `internal/clients/gridscale.go:54-57`; `examples/cluster/providerconfig/secret.yaml.tmpl`; `examples/namespaced/providerconfig/secret.yaml.tmpl` | S | 2026-07-15 | **FIXED** (`b16b5a2`, live-validated 2026-07-15) |
| TEST-1 | Tests | P1 | Zero tests in repo; `make test` gate passes vacuously; no coverage floor. **Largely addressed:** unit tests (external-name, kind overrides), CRD golden-contract tests (E2-S02), and a fuzz test (E2-S09) now exist. Coverage floor (E2-S03) + CI enforcement still pending | `.github/workflows/ci.yml:186-193` | M | 2026-07-15 | **in progress** (largely addressed; coverage floor + CI pending) |
| ARCH-1 | Architecture | P1 | No cross-resource references; generated examples carry unresolved `${...}` HCL interpolations | `config/` (none); `examples-generated/cluster/ssl/v1alpha1/certificate.yaml` | M | 2026-07-15 | **FIXED** (`5959886` — config/references.go, 13 ref edges) |
| DOC-1 | Docs | P1 | ProviderConfig example uses wrong group `gridscale.crossplane.io` vs actual `gridscale.platformrelay.io` | `examples/cluster/providerconfig/providerconfig.yaml:1` vs `apis/cluster/v1beta1/register.go:12` | S | 2026-07-15 | **FIXED** (landed) |
| ARCH-2 | Architecture | P2 | provider-metadata incomplete/malformed: 31 keyed vs 32 resources; ISO-image keyed by title `ISO Image:` with corrupted fields | `config/provider-metadata.yaml:3,10` | S | 2026-07-15 | **FIXED** (`5959886` — hack/metadatafix post-scrape step) |
| TEST-2 | Tests | P2 | e2e wired but non-blocking (comment-gated) and unrunnable today due to CRED-1/ARCH-1 | `.github/workflows/e2e.yaml:96-149` | M | 2026-07-15 | NEW |
| SEC-1 | Security | P2 | golangci suppresses gosec G101 & G104 globally (not just tests) | `.golangci.yml:102-107` | S | 2026-07-15 | NEW |
| DOC-2 | Docs | P2 | Generated examples cover only 6 of ~20 resource groups | `examples-generated/` | M | 2026-07-15 | **FIXED** (`d97d1aa` — all 8 API groups now have a curated example) |
| DOC-3 | Docs | P2 | README is unedited upjet template; broken placeholder text, no gridscale setup | `README.md:10` | S | 2026-07-15 | **FIXED** (landed) |
| DIR-2 | Direction | P2 | Scaffold leftovers: `provider-upjet-aws` changelog label, `upjet-provider-template` goimports prefix, 3-way Go version drift | `cmd/provider/main.go:211`; `.golangci.yml:126`; `Makefile:47`/`ci.yml:13`/`e2e.yaml:12` | S | 2026-07-15 | **FIXED** (`f9bf6f7` label+go1.25, `a27a196` goimports prefix) |
| ARCH-3 | Architecture | P3 | Verify UUID external-name is right for name-identified resources (bucket, storage_import, clone) | `config/external_name.go:19-31` | S | 2026-07-15 | NEW |
| TEST-3 | Tests | P3 | `schema-version-diff` gate dormant (`if: inputs.upjet-based-provider` never set) | `.github/workflows/ci.yml:53-56` | S | 2026-07-15 | **FIXED** (`f9bf6f7` — dead schema-version-diff step removed) |
| DOC-4 | Docs | P3 | CODEOWNERS/OWNERS.md are placeholders; no real maintainer | `CODEOWNERS:16`; `OWNERS.md:9` | S | 2026-07-15 | NEW |
| DIR-3 | Direction | P3 | local-deploy download URL wrong (releases.hashicorp.com) — prior 404 | `Makefile:17` | S | 2026-07-15 | **FIXED** (commit `797adad`, GitHub releases) |
| ~~DIR-1~~ | Direction | — | WITHDRAWN — GUIDELINES.md/BACKLOG.md/decisions.md/ROADMAP exist on main; based on stale local checkout | `agent-context/GUIDELINES.md:1` | — | 2026-07-15 | WITHDRAWN |

| BRAND-1 | Docs/Branding | P3 | Provider icon is an original PLACEHOLDER (`extensions/icons/icon.svg`); commission a real gridscale-provider brand icon (+ social card) and wire `iconURI`. Per D-009→B, deferred post-v0.1.0. | `extensions/icons/icon.svg`; `docs/assets/branding/` | S | 2026-07-16 | **NEW** (post-v0.1.0) |

## Summary
- **Open:** 5 · **Fixed:** 10 (DIR-3, CRED-1/2, DOC-1/3, SEC-1, ARCH-1/2, DIR-2, TEST-3, DOC-2) · **Withdrawn:** 1 (DIR-1) · **Regressed:** 0
- **2026-07-16 v0.1.0 prep:** all 7 audit UNBLOCKED items cleared + CI green (`3f3508b`). Only OPERATOR-BLOCKED item remaining: **TEST-2 / E2-S04-S05** (uptest needs live lab creds as CI secrets). TEST-1 CI-enforced (coverage 92.7%). DOC-5 addressed via `make check-docs` guard.
- **Open by severity:** P0 × 0 · P1 × 0 · P2 × 3 (TEST-2 operator-blocked, DOC-2 residual) · P3 × 2
- **2026-07-15 update:** both P0s (CRED-1/2) FIXED (`b16b5a2`) and **live-validated** (out-of-cluster smoke test, real VM create/delete verified gone via gridscale API); DOC-1/DOC-3 FIXED; TEST-1 largely addressed (unit + CRD golden-contract + fuzz tests exist; coverage floor E2-S03 + CI enforcement pending — still counted Open).

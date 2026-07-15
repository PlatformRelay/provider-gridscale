# PlatformRelay Health & Direction Audit — provider-gridscale

- **Date:** 2026-07-15
- **Auditor:** read-only replayable audit (`/replayable-audit`)
- **Repo:** /Users/A242168/Projects/PlatformRelay/provider-gridscale
- **Commit audited:** `origin/main` @ `00e1c6c` (findings reconciled against real main, not the stale local working checkout `b95c469`)
- **Scope:** Upjet v2 Crossplane provider wrapping `gridscale/gridscale` Terraform provider v2.3.0. Generated `apis/` boilerplate was NOT audited line-by-line; findings target the hand-authored surface.

## Verdict: NEEDS-WORK (not shippable as v0.1.0)

One hard functional blocker (credentials are never passed to Terraform) makes any real resource reconcile impossible. The rest are quality/hygiene gaps consistent with a 2-commit scaffold. The trajectory is coherent (dual cluster/namespaced scope is correctly wired, CI is well-built and SHA-pinned), so this is NEEDS-WORK, not AT-RISK.

## Delta vs prior audit

**This is the BASELINE audit.** The directory `agent-context/archive/audits/` was empty — no prior audit exists, so there is no delta to compute. All findings below are established as first-seen on 2026-07-15. Future runs should diff against this report.

**Reconciliation note (important):** The `design-auditor` first ran against the local working checkout (`b95c469`, 2 commits), which was 4 commits *behind* `origin/main` (`00e1c6c`). All findings were then re-verified against real `main`. Two were affected and corrected below: **DIR-3 was already fixed** on main (commit `797adad`) and **DIR-1 is withdrawn** (roadmap/guidelines docs exist on main). Every other finding was confirmed present on `origin/main` at the cited `path:line`.

---

## Dimension 1 — Architecture

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| ARCH-1 | High | No cross-resource references configured anywhere. All 32 resources rely on raw UUID inputs; relationships (server→ipv4, loadbalancer→ssl cert, etc.) are not wired. Concrete symptom: generated loadbalancer example contains unresolved HCL interpolations `${gridscale_ipv4.server.ip}` / `${gridscale_ssl_certificate...id}` that Crossplane will treat as literal strings. | `config/` (no `Reference{}` present, grep-confirmed); `examples-generated/cluster/ssl/v1alpha1/certificate.yaml` (backendServer host / forwardingRule certificateUuid) |
| ARCH-2 | Medium | `provider-metadata.yaml` is incomplete and malformed: only 31 keyed resource entries vs 32 generated resources, and the ISO-image entry is keyed by a human-readable title `ISO Image:` instead of `gridscale_isoimage`, with corrupted field text (e.g. "storage/ISO Image/ISO Image/snapshot"). | `config/provider-metadata.yaml:3`, `:10` |
| ARCH-3 | Low | External-name strategy `IdentifierFromProvider` is applied to every resource via `.+`. This is a *documented, defensible* choice (gridscale objects are UUID-identified), NOT a placeholder/TODO. Verify a handful that might be name-identified: `object_storage_bucket`, `storage_import`, `storage_clone`. | `config/external_name.go:7-9,19-26,31` |
| ARCH-4 | Info (PASS) | Dual cluster + namespaced scope (Upjet v2 M2 pattern) is correctly wired: `GetProvider`/`GetProviderNamespaced` with distinct root groups, both scopes registered in `main.go`, 32 controllers generated per scope. | `config/provider.go:23-65`; `cmd/provider/main.go:141-142,228-233` |

## Dimension 2 — Tests

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| TEST-1 | High | Zero test files in the entire repo (no `*_test.go`, no `TestExternalName`, no config unit tests, no example-validation test). The CI `unit-tests` job runs `make test` and passes vacuously with no coverage floor. For a generated provider this is the primary quality risk. | `.github/workflows/ci.yml:186-193` (`make test` with no tests present) |
| TEST-2 | Medium | e2e (`uptest`) is wired but gated behind an `/test-examples` maintainer comment and depends on `UPTEST_CLOUD_CREDENTIALS`; it is not part of the PR-blocking gate. Combined with CRED-1 and ARCH-1, e2e cannot pass today even if triggered. | `.github/workflows/e2e.yaml:96-149` |
| TEST-3 | Low | `report-breaking-changes` job's `schema-version-diff` step is guarded by `if: ${{ inputs.upjet-based-provider }}`, an input never provided by push/pull_request triggers, so the native-schema-drift check never runs (dormant gate). | `.github/workflows/ci.yml:53-56` |

## Dimension 3 — Security

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| CRED-1 | **Critical** | **Credentials are extracted then silently discarded.** `TerraformSetupBuilder` calls `CommonCredentialExtractor` and unmarshals the secret, but the block that assigns `ps.Configuration` (username/password) is commented out and never re-added. The gridscale TF provider is therefore never given `uuid`/`token`, so no resource can authenticate. `go vet` on this package is clean — it compiles and fails only at runtime. | `internal/clients/gridscale.go:53-57` |
| CRED-2 | **Critical** | Wrong credential schema throughout. gridscale's provider config expects keys `uuid` and `token` (schema.json provider block: `api_url, http_headers, max_n_retries, request_delay_interval, token, uuid`), but the commented code and both secret templates use `username`/`password`. Even after un-commenting CRED-1, auth would fail. | `internal/clients/gridscale.go:54-57`; `examples/cluster/providerconfig/secret.yaml.tmpl`; `examples/namespaced/providerconfig/secret.yaml.tmpl` |
| SEC-1 | Medium | `.golangci.yml` globally suppresses gosec `G101` (hardcoded credentials) and `G104` (unhandled errors) for all files, not just tests — weakening two security-relevant linters across hand-authored code. | `.golangci.yml:102-107` |
| SEC-2 | Info (PASS) | Supply-chain hygiene is good: all GitHub Actions are SHA-pinned; e2e uses the safer `issue_comment` + repo-permission gate (admin/write) rather than `pull_request_target`; secrets are referenced, not echoed. | `.github/workflows/ci.yml:37,71,99`; `.github/workflows/e2e.yaml:15-36,100` |

## Dimension 4 — Docs

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| DOC-1 | High | ProviderConfig example is unusable: declares `apiVersion: gridscale.crossplane.io/v1beta1`, but the actual generated group is `gridscale.platformrelay.io` (register.go). Applying the example fails "no matches for kind". | `examples/cluster/providerconfig/providerconfig.yaml:1` vs `apis/cluster/v1beta1/register.go:12` |
| DOC-2 | Medium | Generated examples do not cover all 32 resources — only 6 distinct resource groups have examples (`gridscale, mysql8, paas, redis, ssl, storage`), 46 yaml files total. Most resources ship with no example. | `examples-generated/` (6 of ~20 groups populated) |
| DOC-3 | Medium | README is the unedited upjet template: broken placeholder sentence "This gridscale serves as a starting point…", no ProviderConfig/credential setup, no install instructions specific to gridscale. | `README.md:10` |
| DOC-4 | Low | Ownership files are placeholders: CODEOWNERS has only a commented `#* @username` fallback; OWNERS.md lists `Full Name <email@example.com> (githubusername)`. No real maintainer is assigned. | `CODEOWNERS:16`; `OWNERS.md:9` |

## Dimension 5 — Direction

| ID | Sev | Finding | Location |
|----|-----|---------|----------|
| ~~DIR-1~~ | ~~Medium~~ WITHDRAWN | Original finding ("no AGENTS.md/GUIDELINES.md/roadmap in-repo") was based on the stale local checkout. On real `main` these exist: `agent-context/GUIDELINES.md`, `agent-context/BACKLOG.md`, `agent-context/decisions.md`, plus pointers to `docs/ROADMAP.md`, `openspec/changes/`, and a workspace `AGENTS.md`. **Not a finding.** | `agent-context/GUIDELINES.md:1`; `agent-context/BACKLOG.md:1-3` |
| DIR-2 | Medium | Multiple copy-paste scaffold artifacts signal "generate-all-32-at-once" outran hand-configuration: (a) changelog provider label hardcoded `provider-upjet-aws:%s`; (b) goimports local-prefix still `upjet-provider-template`; (c) three-way Go version drift (Makefile `1.24`, ci.yml `1.25`, e2e.yaml `1.24`). **All three confirmed on `origin/main`.** | `cmd/provider/main.go:211`; `.golangci.yml:126`; `Makefile:47` / `ci.yml:13` / `e2e.yaml:12` |
| ~~DIR-3~~ | ~~Low~~ **ALREADY FIXED on main** | `local-deploy` download URL previously defaulted to `releases.hashicorp.com` (the prior 404). Fixed on `origin/main` by commit `797adad` — `Makefile:17` now derives the URL from `$(TERRAFORM_PROVIDER_REPO)/releases/download/...` (GitHub releases). Recorded here for the baseline trail; **no action needed.** | `Makefile:17` (fixed) |
| DIR-4 | Info | **Upstream-bug note (operator memory):** No evidence of upstream `terraform-provider-gridscale` bugs or workarounds was found in this repo's hand-authored config this run (no patches/overrides dir, no BUG/HACK/workaround markers in `config/`, `internal/clients/`, `cmd/`). Nothing to PR upstream surfaced from this baseline. | (searched `config/*.go`, `internal/clients/*.go`, `cmd/`, `extensions/`) |

---

## Top 3 to fix next

1. **CRED-1 + CRED-2 (Critical):** Re-enable `ps.Configuration` in `internal/clients/gridscale.go` and map the secret to gridscale's real keys `uuid`/`token` (and optional `api_url`). Fix both secret templates to match. Without this, nothing reconciles. → `/write-story`
2. **TEST-1 (High):** Add a `TestExternalName` / config unit test and an example-validation test so the vacuous `unit-tests` gate has teeth; add a coverage floor. → `/write-story`
3. **DOC-1 + ARCH-1 (High):** Fix the ProviderConfig example group (`gridscale.platformrelay.io`) and add cross-resource references for the obvious edges (server↔ipv4, loadbalancer↔ssl cert/ipv4/ipv6) so generated examples resolve and can pass e2e. → `/write-story`

## Gate re-run note (this session)
- `go vet ./config/... ./internal/clients/... ./cmd/...` → **exit 0 (clean)**. Confirms CRED-1 is a silent runtime failure, not a compile error. No full generated-`apis/` build was run (out of scope; read-only).
- CRED-1 independently re-verified by coordinator against `internal/clients/gridscale.go:44-58` before filing.

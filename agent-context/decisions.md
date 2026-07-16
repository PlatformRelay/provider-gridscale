# decisions.md — provider-gridscale (append-only)

Format per entry: **Decision** · date · context · choice · agent counterpoints (kept even when
overruled). Open decisions awaiting the operator live in [`INBOX.md`](INBOX.md); when answered they
move here.

---

## D-000 — Adopt kaddy's spec-driven (OpenSpec) workflow, committed

**Date:** 2026-07-15
**Context:** provider-gridscale was scaffolded with no agent-context, specs, or backlog. The
operator asked to port kaddy's spec-driven workflow and raise the repo to kollect/mkurator OSS
maturity, and said agent files may be *kept* (committed) here.
**Decision:** Port OpenSpec (`openspec/`), `agent-context/`, root `AGENTS.md`, and `docs/adr/`.
**Commit** them (unlike kaddy, where they're gitignored) — for a public provider a visible
spec-driven harness is a maturity signal. Only `.claude/` stays local/gitignored.
**Counterpoints:** Committing process docs adds surface a casual reader may skip; mitigated by
keeping `AGENTS.md` a thin doc-map and putting detail behind links.

---

## D-005 — Provider-shaped test taxonomy, not kaddy's L0–L4 (signed off)

**Date:** 2026-07-15 · **Status:** Accepted — operator signed off 2026-07-15 (approved as-is).
**Context:** kaddy's pyramid is L0 tofu / L1 conftest / L2 chainsaw / L3 k6 / L4 scorecard — a
*platform* shape. kollect/mkurator hand-write reconcilers so lean on Ginkgo + envtest + mockery. A
*generated* provider has neither shape: its testable surface is a small hand-authored `config/`
layer, the generated CRD schemas, and example-as-test e2e.
**Decision:** Adopt the four-tier taxonomy in **ADR-0002** — **U** (Go unit on custom
external-name logic) · **M** (CRD golden-contract, flagship + structural invariants) · **E** (uptest,
which already renders Chainsaw) · **D** (doc-sync). Explicitly skip envtest/Ginkgo reconcile tests,
mockery, standalone Chainsaw, and repo-wide coverage floors (coverage scoped to `./config/...`).
Keep kaddy's discipline (every REQ has one **Test:** + one **Verify:**).
**Counterpoints:** Divergent taxonomy across sibling repos costs some cross-repo familiarity; worth
it because the siblings' style is a consequence of hand-written controllers we don't have, and
forcing chainsaw/tofu/envtest onto a generated provider would produce hollow, high-churn tests.

---

## D-001…D-004 — Epic-shape approval (all answered **A**)

**Date:** 2026-07-15 · **Status:** Accepted — operator approved the E1–E6 plan and each decision's
Recommended option (**A**) 2026-07-15.

- **D-001 — Epic numbering → A (fresh repo-local E1–E6).** kaddy's "E6g" stays kaddy's external
  handle; this repo's IDs are `E1-S01…E6-Snn`. History not renumbered. Counterpoint: loses one
  visible lineage to kaddy — accepted, per the workspace repo-independence rule.
- **D-002 — Documentation depth → A (provider-native).** `crd-ref-docs` API reference + curated
  `examples/` + rich README + Marketplace README; no MkDocs site. **E4-S04 dropped.** Counterpoint:
  diverges from kollect/mkurator's site convention — accepted; a generated provider's docs are its
  CRD reference, and a site is re-addable later (was option C).
- **D-003 — Build tooling → A (keep Makefile + `build/` submodule).** Add thin `hack/` scripts and
  new make targets (`test.spec`, `docs`) around upstream; no Taskfile. Counterpoint: runner
  differs from sibling app repos — accepted; matches upjet tooling and avoids fighting upstream.
- **D-004 — Supply-chain scope → A (pragmatic baseline).** codecov + Scorecard + CodeQL +
  govulncheck + gitleaks + pre-commit + dependabot + cliff + cosign/SBOM. **SonarCloud, OpenVEX,
  and the assurance case (E6-S05) stay out/stretch.** Counterpoint: weaker signal than full
  parity — accepted as proportionate for a generated provider's first release.

---

## D-006 — Unblock `main` CI: merge sequence + `local-deploy` root-cause correction → A

**Date:** 2026-07-15 · **Status:** Accepted — operator chose **A** ("go with option A and merge the
rest as well") 2026-07-15.
**Context:** `main` CI was red on `lint` (gofmt) and `local-deploy`. A prior INBOX note diagnosed the
`local-deploy` red as an unfixable upstream `build/` submodule / buildx-cache infra issue needing
operator/infra. **That was wrong:** buildx's `failed to load cache key: invalid response status 404`
was surfacing a failed `ADD <url>` in the provider-image Dockerfile — the terraform-provider binary
was fetched from `releases.hashicorp.com/terraform-provider-gridscale/2.3.0/...`, which 404s because
`releases.hashicorp.com` only hosts HashiCorp's own providers. Root cause: a one-line wrong
`TERRAFORM_PROVIDER_DOWNLOAD_URL_PREFIX` in the repo's own `Makefile` — fully in-repo fixable.
**Decision:** Repoint the prefix at the gridscale GitHub releases (reusing `TERRAFORM_PROVIDER_REPO`);
carry it on PR #1 alongside the gofmt fix; merge PR #1 (rebase), then rebase+merge PR #2 (E1), then
start E2–E6. The fix was surfaced (agent-loop-auto), not auto-merged, because an explicit operator
gate named this blocker and it is a build/supply-chain change; the operator then approved A.
**Counterpoints:** the `ADD` still verifies no checksum/signature on the provider zip (pre-existing,
tracked as a supply-chain follow-up, not introduced here); reusing `TERRAFORM_PROVIDER_REPO` couples
the prefix to GitHub's `/releases/download/v<ver>` asset-path convention — correct today, revisit only
if `REPO` is ever pointed at a non-GitHub host.

---

## D-007 — `govulncheck` remediation: Go 1.26.5 + `x/net` v0.55.0 → landed (reconciled)

**Date:** 2026-07-16 · **Status:** Accepted — **landed on `main` as `d75721e`** (`go 1.25.9` /
`toolchain go1.26.5`; `golang.org/x/net v0.55.0`). Recorded here **retroactively** to reconcile a
stale INBOX entry that still listed this as an *open, awaiting-go-ahead* security decision after the
change had already merged.
**Context:** `make vuln` (E2-S06) surfaced 3 **called** vulnerabilities — GO-2026-5856 (`crypto/tls`
Encrypted Client Hello privacy leak; fixed go1.26.5), GO-2026-5026 (`x/net/idna` Punycode) and
GO-2026-4918 (`x/net` HTTP/2 infinite loop; both fixed in `x/net` v0.55.0). Remediation was a
Go-toolchain + dependency bump, which the auto-loop guardrails classify as **surfaced-not-auto-merged**.
**Decision:** Apply the bump (toolchain `go1.26.5`, `x/net` v0.55.0); `make vuln` after = 0 affecting
our code; BSL `TERRAFORM_VERSION=1.5.7` pin untouched.
**Process counterpoint (kept):** the commit landed on `main` **before** an operator answer was logged
and before this entry existed — either it was approved out-of-band or a prior session merged a
surfaced security/deps change. It clears real called CVEs and `main` is green, so it is **not
reverted**; the gap is noted so future surfaced items are logged at merge time, not after. The
follow-up CI wiring for `govulncheck` remains **E5-S01** (surfaced), separate from this local bump.

---

## D-008 — Real maintainer identity → A (Konrad Heimel, sole initial maintainer)

**Date:** 2026-07-16 · **Status:** Accepted — operator answered **A** (was tracked as **DOC-4**).
**Context:** `CODEOWNERS`, `OWNERS.md`, and `GOVERNANCE.md` shipped as template stubs (`@username`,
`Full Name <email@example.com>`, "maintainers TBD"); no manifest — including the Marketplace
`maintainer` annotation in `package/crossplane.yaml` — could name a real owner until this was answered.
**Decision:** Konrad Heimel as the sole initial maintainer (operator + repo author). Fill all three
governance files and the Marketplace `maintainer` with that identity: `Konrad Heimel`, GitHub handle
`@konih`, email `konrad.heimel@gmail.com`. **Verified on `main`:** `CODEOWNERS` fallback + scoped
owners are `@konih`; `OWNERS.md` lists `Konrad Heimel <konrad.heimel@gmail.com> ([konih])`;
`GOVERNANCE.md` names "Konrad Heimel (@konih)" as current maintainer (no "TBD"); and
`package/crossplane.yaml` carries `meta.crossplane.io/maintainer: Konrad Heimel <konrad.heimel@gmail.com>`.
Landed in `dff9a33` (governance files) + `860c5e6` (crossplane maintainer annotation).
**Counterpoints:** (1) Solo-maintainer **bus-factor** — a single owner is a single point of failure;
accepted for an early-stage project, with GOVERNANCE.md's "Becoming a maintainer" path open for
co-maintainers later. (2) **Stale cross-reference:** `GOVERNANCE.md` line ~55 still calls the
maintainer identity "an open operator decision (tracked as DOC-4)", which is now false — a cosmetic
follow-up one-liner (out of scope for this 2-file reconciliation), not a blocker; the maintainer is
correctly named at line ~53.

---

## D-009 — E3-S01 branding icon → B (ship placeholder for v0.1.0) — COORDINATOR decision

**Date:** 2026-07-16 · **Status:** **Coordinator decision** under the operator's standing
"decide-for-B-and-ask-later" latitude — **the operator can override this.**
**Context:** `extensions/icons/icon.svg` (+ light/dark variants under `docs/assets/branding/`) landed
as a **placeholder** (`df4a3ec`) — an original, neutral, valid SVG that renders correctly. The open
question was whether v0.1.0 must wait on a real commissioned brand icon before publishing to the
Marketplace and embedding in the README.
**Decision:** Ship v0.1.0 with the placeholder icon; do **not** block the release on branding. Track a
real commissioned icon as post-release polish (new backlog/debt item **BRAND-1: commission real
provider icon**). Rationale: the placeholder is a valid, neutral, original asset that renders fine; a
real brand icon is a subjective designer call that should not gate an otherwise-ready release.
Consequently `iconURI` is intentionally left out of `package/crossplane.yaml` for now (see D-010).
**Counterpoints:** A placeholder icon on the Upbound Marketplace listing is a **weaker first
impression** than a commissioned brand mark — accepted, revisit post-v0.1.0 under BRAND-1. Because
this is a coordinator call rather than an operator answer, it is recorded explicitly so the operator
can reverse it (choose option A/C) before the release is cut.

---

## D-010 — E3-S03 Marketplace metadata → A (filled; `iconURI` deferred to branding)

**Date:** 2026-07-16 · **Status:** Resolved — objective fields done on `main`.
**Context:** the INBOX text claimed `package/crossplane.yaml` carried only `capabilities: [SafeStart]`
with no `meta.crossplane.io` annotations and an empty `extensions/readme/readme.md`. That text was
**stale** — the manifest had already been populated. This manifest renders the Upbound Marketplace
listing and ships inside the published XPKG (release-adjacent).
**Decision:** Fill the Marketplace metadata in one pass. **Verified on `main`:** `crossplane.yaml` now
carries `meta.crossplane.io` annotations — `maintainer`, `source` (`github.com/PlatformRelay/
provider-gridscale`), `license` (`Apache-2.0`), `description`, `readme`, and a `friendly-name` — plus
`capabilities: [SafeStart]`; and `extensions/readme/readme.md` is populated (~3986 bytes). Landed
`860c5e6`. **`iconURI` is intentionally NOT set** — it is deferred to the branding decision (D-009 →
placeholder / BRAND-1), so the objective metadata is resolved while the icon URL waits on a real
brand asset.
**Counterpoints:** Publishing the objective metadata before a final `iconURI` means the manifest is
touched twice (once now, once when BRAND-1 lands the real icon) — accepted; the objective fields are
correct and stable on first publish, and the icon is the only field legitimately still in flux.

---

## D-011 — E5 CI/CD hardening & supply-chain epic (S01–S06) → A (implemented, green)

**Date:** 2026-07-16 · **Status:** Resolved — implemented and **all CI workflows green on `main`**.
**Context:** scope was fixed by **D-004→A** (codecov, Scorecard, CodeQL, govulncheck job, gitleaks,
pre-commit, dependabot/renovate, cliff, cosign/SBOM). Local quality *tools* had already landed
(E2-S06…S10); wiring them into `.github/workflows/` was the E5 epic and had been deliberately deferred
as surfaced-not-auto-merged. E3-S04 README badges also waited on these jobs existing.
**Decision:** Implement E5 lane-by-lane, adding workflow files rather than editing stock `ci.yml`.
**Verified on `main`:** the workflows exist — `coverage.yml`, `govulncheck.yml`, `scorecard.yml`,
`codeql.yml`, `gitleaks.yml`, `changelog.yml` — plus `codecov.yml`, `.github/gitleaks.toml`,
`.pre-commit-config.yaml`, `cliff.toml`, tuned `.github/renovate.json5`, and keyless cosign signing +
SBOM in `publish-provider-package.yml`. Landed across `8980b57` (coverage/govulncheck/scorecard/codeql),
`273d49e` (gitleaks + pre-commit), `dc699c4` (renovate tuning + cliff changelog), `aa5dba6`
(cosign/SBOM), with integration fixes `05114ed` (tokenless Codecov OIDC) / `d936e5a` (renovate matcher
+ unsigned-publish warning). **CI status:** CI, Coverage, Govulncheck, Scorecard, CodeQL, Gitleaks all
**green** on `main`. E3-S04 badges are consequently satisfied (README badge row: CI / Coverage /
CodeQL / Scorecard / Go Report / Go Version / Release / License).
**Counterpoints:** per D-004 this is a **pragmatic baseline** (SonarCloud, OpenVEX, and the E6-S05
assurance case stay out/stretch) — a weaker signal than full supply-chain parity, accepted as
proportionate for a generated provider's first release. Cosign signing/SBOM only run when
`publish-provider-package.yml` is dispatched with an explicit version input; an input-less run emits a
visible "package will be UNSIGNED" warning (by design) rather than failing.

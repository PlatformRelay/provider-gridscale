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

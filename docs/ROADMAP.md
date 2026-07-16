# ROADMAP — provider-gridscale

Public build order for taking `provider-gridscale` from a freshly-scaffolded Upjet provider to a
professional, releasable OSS product. Private lane detail lives in
[`agent-context/BACKLOG.md`](../agent-context/BACKLOG.md); OpenSpec changes live under
[`openspec/changes/`](../openspec/changes/).

> **Unaffiliated.** This is a community PlatformRelay project — **not** affiliated with, endorsed by,
> or an official product of gridscale GmbH.

## Where we are

`provider-gridscale` **generates and builds** today: 32 resources across 8 API groups
(`gridscale`, `marketplace`, `mysql8`, `object`, `paas`, `redis`, `ssl`, `storage`), controllers
and CRDs wired, strict `golangci-lint` (47 linters). What's missing is everything that makes a
provider *trustworthy and adoptable*: tests, docs, branding, hardened CI, and governance.

| Dimension | Status |
| --- | --- |
| Resource coverage / codegen | ✅ complete — 32 resources, `make generate` green |
| Tests | ❌ zero `*_test.go`, no uptest examples wired |
| Branding / Marketplace | ❌ `icon.svg` empty, README stub, no package metadata |
| Documentation | ⚠️ 1-page README only, no API reference, no ADRs |
| CI / supply chain | ⚠️ stock upjet CI; no coverage, scans, changelog, signing |
| Community health | ⚠️ template CODEOWNERS/OWNERS/CoC; no CONTRIBUTING/SECURITY |
| Spec-driven harness | ❌ none — this backlog establishes it |

## Epics

| Epic | Title | Outcome | Depends on |
| --- | --- | --- | --- |
| **E1** | Spec-driven workflow & governance foundation | The harness ported & committed | — |
| **E2** | Test foundation | Provider is testable (unit + uptest + coverage) | E1 |
| **E3** | Branding & Marketplace identity | Looks & lists like a real provider | E1 |
| **E4** | Documentation generation | Generated API docs + ADRs (+ optional site) | E1 (E3 for assets) |
| **E5** | CI/CD hardening & supply chain | CI at kollect/mkurator bar, provider-shaped | E2 (coverage), E3 (badges) |
| **E6** | Community health & governance | Contributor-ready, OpenSSF-aligned | E1 |

### Lane dependency graph

```
E1  (foundation — everything references it)
 ├── E2  tests ─────────────┐
 ├── E3  branding ──────────┤
 ├── E4  docs ──────────────┤
 ├── E6  governance         │
 └────────────────────────► E5  CI/supply-chain  (needs E2 coverage + E3 badges/metadata)
```

E1 lands first. E2 · E3 · E4 · E6 run in parallel once E1 exists (disjoint paths). E5 integrates
last because its gates depend on E2's coverage output and E3's Marketplace metadata.

---

## E1 — Spec-driven workflow & governance foundation

**Why:** Establish the traceable, TDD, decision-logged workflow (ported from `kaddy`) so every
later lane maps to a spec with a test. For a public provider, a *visible* spec-driven harness is a
maturity signal, so it is committed (not gitignored). OpenSpec change: `e1-workflow-foundation`.

| Story | Title | Level |
| --- | --- | --- |
| E1-S01 | OpenSpec scaffold (`config.yaml`, `changes/`, `specs/`) + `make test.spec` gate | M |
| E1-S02 | `agent-context/` — AGENTS map, GUIDELINES (Go/provider), decisions.md, BACKLOG, INBOX | M |
| E1-S03 | Root `AGENTS.md` entry point + `.claude/skills/pick-next-story` | M |
| E1-S04 | ADR convention (`docs/adr/`) + ADR-0001 spec-driven workflow + `docs/ROADMAP.md` | M |

**Exit:** `make test.spec` runs and passes on the seeded specs; ADR-0001 recorded; a contributor
can find "what to do next" from `AGENTS.md` in one hop.

---

## E2 — Test foundation

**Why:** The provider ships zero tests. Make it *testable* the provider-native way (see **ADR-0002**):
unit tests on the small hand-authored `config/` surface, CRD golden-contract tests to catch
generation drift, and uptest example-as-test e2e — **not** envtest/Ginkgo reconcile tests (the
controllers are generated) and **not** standalone Chainsaw (uptest already renders it). OpenSpec
change: `e2-test-foundation`.

| Story | Title | Level |
| --- | --- | --- |
| E2-S01 | Unit-test the custom external-name logic (`idWithStub().GetExternalNameFn`: tfstate→UUID; include-list) | U |
| E2-S02 | **CRD golden-contract tests** — extract OpenAPI fragment from `package/crds/` for a scoped subset (MySQL8, StorageImport, Server, Network), diff vs golden (`UPDATE_GOLDEN=1` to refresh) | M |
| E2-S03 | Wire `make test` (unit + contract) + coverage scoped to `./config/...` (floor ~70%) + `make test.spec` gate | M |
| E2-S04 | Curate uptest smoke subset (server, network, storage, ipv4, sshkey) + datasource/ref wiring — **needs lab access** | E |
| E2-S05 | Enable `e2e.yaml` uptest run on `/test-examples` + nightly + document creds contract | E |
| E2-S06 | `govulncheck` make target (`make vuln` → `go run golang.org/x/vuln/cmd/govulncheck@v1.1.4 ./...`) *(ported from kollect's test tooling)* | M |
| E2-S07 | `go-arch-lint` config (`.go-arch-lint.yml`) + `make arch-lint` enforcing the generation boundary (hand-authored `config`/`internal/clients` must not import generated `internal/controller/**`) *(ported from kollect's test tooling)* | M |
| E2-S08 | Race-detector make target (`make test.race` → `go test -race -count=1`) *(ported from kollect's test tooling)* | U |
| E2-S09 | Native fuzz test on the `config` external-name extraction (`FuzzGetExternalName`) *(ported from kollect's test tooling)* | U |
| E2-S10 | `go mod` hygiene gate (`make tidy-check` → `go mod tidy` + `git diff --exit-code go.{mod,sum}`) *(ported from kollect's test tooling)* | M |

> The mysql8/storageimport Kind overrides get **no standalone unit story** — they're guarded by
> compilation (`make build`) + the CRD-contract tier; an optional assertion can ride E2-S02.
>
> **E2-S06…S10 provenance & CI split (ported from kollect's test tooling):** these stories add the
> local `make` targets + tests **now** (auto-mergeable — they touch only `Makefile`, `config/*_test.go`,
> and `.go-arch-lint.yml`). **Wiring each tool into CI (`.github/workflows/`) is deferred to Epic E5**
> (CI/release changes are surfaced-not-auto-merged per the harness guardrails). Concretely: E2-S06's
> `make vuln` target lands here; the **govulncheck CI job is the existing E5-S01 follow-up** — same for
> arch-lint / race / fuzz / tidy-check gating, which become E5 CI jobs, not new work.

**Exit:** `make test` runs unit + CRD-contract tests with `config/`-scoped coverage; scoped CRD
goldens committed; ≥1 uptest example green against a lab project (gated on lab creds); coverage
artifact produced for E5 to gate on.

---

## E3 — Branding & Marketplace identity

**Why:** `icon.svg` is empty, the README is a stub, and `package/crossplane.yaml` carries no
Marketplace metadata — so the provider is unlistable and unrecognisable. Give it a visual identity
and a proper listing. OpenSpec change: `e3-branding-marketplace`.

| Story | Title | Level |
| --- | --- | --- |
| E3-S01 | Visual identity — square `extensions/icons/icon.svg` + logo light/dark + social card in `docs/assets/branding/` | M |
| E3-S02 | README overhaul — badge row, logo, quick-start, resource matrix, community links | M |
| E3-S03 | Marketplace metadata — `package/crossplane.yaml` annotations (maintainer, source, license, iconURI, description, readme) + fill `extensions/readme/readme.md` + release-notes | M |
| E3-S04 | Badge wiring — CI, coverage, Go Report Card, license, release, Marketplace version (resolve after E5 jobs exist) | M |

**Exit:** `icon.svg` renders; README leads with logo + badges + quick-start; `crossplane.yaml`
carries full Marketplace annotations; provider previews correctly on the Upbound Marketplace.

---

## E4 — Documentation generation

**Why:** A provider's docs are primarily its CRD API reference plus runnable examples — generated,
not hand-written prose. Establish `crd-ref-docs` + examples output and the ADR trail; decide
whether to also stand up an MkDocs Material site like kollect/mkurator (see **D-002**). OpenSpec
change: `e4-docgen`.

| Story | Title | Level |
| --- | --- | --- |
| E4-S01 | `crd-ref-docs` config + `make docs` → `docs/api/` CRD reference; sync check in CI | D |
| E4-S02 | `examples-generated` wired; one curated example per API group under `examples/` | D |
| E4-S03 | Seed `docs/adr/` with foundational ADRs (codegen boundary, test taxonomy, build tooling) | M |
| ~~E4-S04~~ | ~~MkDocs Material site + `docs.yaml` GitHub Pages deploy~~ — **dropped** (D-002→A) | — |

**Exit:** `make docs` regenerates API reference deterministically; CI fails on stale docs; ADRs
capture the non-obvious decisions. E4-S04 is **dropped** — D-002 chose provider-native docs (A);
no MkDocs site (re-addable later if desired).

---

## E5 — CI/CD hardening & supply chain

**Why:** CI is the stock upjet template — no coverage gate, security scanning, changelog, or signed
releases. Bring it to the kollect/mkurator bar, scoped to what fits a provider (keep the upstream
`build/` submodule; add jobs around it — see **D-003**). OpenSpec change: `e5-ci-supplychain`.

| Story | Title | Level |
| --- | --- | --- |
| E5-S01 | Codecov upload + `codecov.yml` floor; `govulncheck` job | M |
| E5-S02 | OpenSSF Scorecard + CodeQL workflows | M |
| E5-S03 | `gitleaks` config + `.pre-commit-config.yaml` (gitleaks, golangci-lint, shellcheck, markdownlint) | M |
| E5-S04 | Dependabot/Renovate for Go modules + GitHub Actions (grouped) | M |
| E5-S05 | `cliff.toml` changelog from gitmoji-conventional commits + release-notes automation | M |
| E5-S06 | Signed provider package — cosign keyless + SBOM on publish (extend existing `publish-provider-package.yml`) | M |

**Exit:** PRs get coverage + vuln + secret + CodeQL feedback; `CHANGELOG.md` generates from
history; released XPKG is signed with an SBOM. SonarCloud is out of scope unless **D-004** widens it.

---

## E6 — Community health & governance

**Why:** Template CODEOWNERS/OWNERS/CoC and no CONTRIBUTING/SECURITY mean the provider isn't
contributor- or disclosure-ready. Fill the governance surface to the OpenSSF baseline. OpenSpec
change: `e6-governance`.

| Story | Title | Level |
| --- | --- | --- |
| E6-S01 | `CONTRIBUTING.md` — standards map, merge policy, local preflight, commit convention | M |
| E6-S02 | `SECURITY.md` — reporting, supported versions, threat model, supply-chain posture | M |
| E6-S03 | `GOVERNANCE.md` + fill `CODEOWNERS` and `OWNERS.md` with real maintainers | M |
| E6-S04 | Upgrade `CODE_OF_CONDUCT.md` to Contributor Covenant 2.1 + issue/PR templates | M |
| E6-S05 | *(stretch)* Assurance case + SCA-remediation policy (OpenSSF baseline docs) | M |

**Exit:** A newcomer can find how to contribute, report a vuln, and see who owns the repo; CoC is a
recognised standard; OpenSSF Scorecard's governance checks pass.

---

## Deferred / non-goals (this backlog)

- **Helm packaging / UI** (kollect-style) — a provider ships as an XPKG, not a chart. Out of scope.
- **SonarCloud** — pending **D-004**; Scorecard + CodeQL + govulncheck cover the baseline.
- **Re-generating resources / changing Upjet config for coverage** — codegen is complete (E6g-S01);
  this backlog does not touch the generated tree except through `config/`.
- **Composition/XRD work** — that's kaddy's E6g-S03; not this provider's concern.

Scope-shaping decisions — all **resolved (A)**; see
[`agent-context/decisions.md`](../agent-context/decisions.md): **D-001** (numbering), **D-002** (docs
depth → provider-native, E4-S04 dropped), **D-003** (build tooling), **D-004** (parity scope),
**D-005** (test taxonomy), **D-006** (CI unblock / provider-download URL).

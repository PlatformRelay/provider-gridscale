# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Loop status (2026-07-16, `/agent-loop-auto`):** the auto-mergeable backlog is **exhausted**. The
> last clean lane — **E3-S02 README enrichment** (resource matrix + community links) — landed on
> `main`. Everything remaining is **operator-gated** (maintainer identity, real branding, Marketplace
> metadata, the E5 CI/supply-chain epic, and uptest lab creds) and is surfaced below as decisions
> **D-008…D-012**. The loop has **stopped at this stop condition** per the auto-merge guardrails; it
> did not auto-merge any release/CI/API-surface change. Answer the decisions below (any session) to
> unblock the next batch.

---

## Decisions (open)

### D-008 — Real maintainer identity (was **DOC-4**)

**Blocks:** E6-S03 (fill `CODEOWNERS`, `OWNERS.md`, `GOVERNANCE.md` maintainers) **and** the
`maintainer` annotation in E3-S03's `package/crossplane.yaml`.
**Context:** `CODEOWNERS` (`#*  @username`), `OWNERS.md` (`Full Name <email@example.com>
([githubusername])`), and `GOVERNANCE.md` ("maintainers TBD") are all still **template stubs**. No
manifest can name a real owner until this is answered.
**Options:**
- **A (Recommended)** — Konrad Heimel as sole initial maintainer (operator + author of the repo);
  fill all three files + the Marketplace `maintainer` with that identity (name, GitHub handle, email).
- **B** — Konrad + one or more named co-maintainers (provide handles/emails).
- **C** — Defer to post-v0.1.0; ship with a generic "PlatformRelay maintainers" placeholder and a
  tracking note.
**Answer:** _(pending)_

### D-009 — E3-S01 branding: real icon vs. placeholder

**Blocks:** finalizing E3-S01, the README logo (deferred half of E3-S02), and E3-S03's `iconURI`.
**Context:** `extensions/icons/icon.svg` + light/dark variants landed as a **placeholder** (`df4a3ec`).
A placeholder must not be published as the Marketplace icon or embedded as the README logo.
**Options:**
- **A (Recommended)** — Provide/commission a real square gridscale-provider icon (+ light/dark + social
  card) before v0.1.0, then wire it into README + `iconURI`.
- **B** — Ship v0.1.0 with the placeholder; track the real icon as a post-release polish item.
- **C** — Reuse an existing PlatformRelay/gridscale brand asset (point me at it).
**Answer:** _(pending)_

### D-010 — E3-S03 Marketplace metadata (`package/crossplane.yaml`)

**Depends on:** D-008 (maintainer) + D-009 (iconURI). **Surfaced — release manifest, not auto-merged.**
**Context:** `package/crossplane.yaml` currently carries only `capabilities: [SafeStart]` — no
`meta.crossplane.io` annotations (maintainer, source, license, iconURI, description, readme), and
`extensions/readme/readme.md` is **empty**. This is what renders the Upbound Marketplace listing, so
it goes into the published XPKG (release-adjacent).
**Options:**
- **A (Recommended)** — Fill it in one pass **once D-008 + D-009 are answered** (so maintainer + real
  `iconURI` are correct on first publish); populate `extensions/readme/readme.md` at the same time.
- **B** — Fill the objective fields now (source, license, description, readme path) with maintainer +
  icon left as explicit `TBD` — produces a half-baked manifest that gets redone; **not recommended**.
- **C** — Defer entirely to release-prep.
**Answer:** _(pending)_

### D-011 — E5 CI/CD hardening & supply-chain epic (S01–S06)

**Surfaced — CI/release; the largest remaining epic.** Scope fixed by **D-004→A** (codecov, Scorecard,
CodeQL, govulncheck job, gitleaks, pre-commit, dependabot, cliff, cosign/SBOM).
**Context:** the local quality **tools** already landed (E2-S06…S10: `make vuln`/`arch-lint`/
`test.race`/fuzz/`tidy-check`). Wiring each into `.github/workflows/` is E5 and was **deliberately
deferred** as surfaced-not-auto-merged. E3-S04 badges also wait on these jobs existing.
**Options:**
- **A (Recommended)** — Green-light the harness to implement E5 lane-by-lane, **adding** workflow files
  rather than editing stock `ci.yml` where possible; each merge still surfaced for your review before
  landing (CI/release guardrail stays on).
- **B** — You implement/review E5 directly (CI secrets, Codecov token, cosign identity are yours).
- **C** — Hold E5 until v0.1.0 scope is locked.
**Answer:** _(pending)_

### D-012 — E2-S04/S05 uptest e2e coverage

**Blocked on lab access.** Not parallelizable (machine-locked live-API suite).
**Context:** CRED-1/2 are **live-validated** by a single manual smoke test (real Server create/delete,
2026-07-15), but automated uptest coverage (curated smoke subset + `e2e.yaml` on `/test-examples` +
nightly) needs **real gridscale lab creds** wired as CI secrets and documented (creds contract →
E2-S05). Running out-of-cluster also needs the `terraform` 1.5.7 binary on `$PATH`.
**Options:**
- **A** — Provide lab creds (as CI secrets) + green-light the nightly/`/test-examples` wiring so uptest
  runs automatically.
- **B (Recommended for v0.1.0)** — Keep the manual smoke test as the e2e evidence for v0.1.0; schedule
  automated uptest for a later milestone.
- **C** — Defer.
**Answer:** _(pending)_

---

## Reference — resolved / no action

- **D-007 (logged 2026-07-16)** — the `govulncheck` remediation (Go 1.26.5 + `x/net` v0.55.0) **already
  landed** on `main` (`d75721e`); the previously-open INBOX security item was **stale** and has been
  reconciled into [`decisions.md`](decisions.md) (with a process note that a surfaced security/deps
  change merged before being logged). `main` green; no action needed.
- **CRED-1/CRED-2 live-validated (2026-07-15)** — end-to-end proof the credential wiring (`b16b5a2`)
  works: authenticated to the live gridscale API, created + deleted a real Server, verified gone via
  API. Single manual smoke test; automated coverage tracked under **D-012**.
- **AUDIT 2026-07-15** — baseline health audit → NEEDS-WORK for v0.1.0; the 2×P0 (CRED-1/2) + DOC-1 +
  DOC-3 have since landed. Report → [HEALTH-AUDIT-2026-07-15.md](archive/audits/HEALTH-AUDIT-2026-07-15.md).
- **E6-S05 (assurance case / SCA-remediation policy)** — **stretch / out of scope** per D-004→A; no
  action unless the operator widens scope.

- **DECISION — RELEASE v0.1.0 — awaiting operator go-ahead.** `main` is release-ready (zero UNBLOCKED backlog items; all 6 CI workflows green incl. `check-diff`; coverage 92.7 percent; audit HEALTH-AUDIT-2026-07-16). I **deliberately did NOT auto-cut the release** — it is public + signed + irreversible and a second agent-loop session was concurrently pushing to `origin/main` (tag-race / double-publish risk). Two things needed from you: (1) confirm the other session on this repo is stopped; (2) say "release" and I run `gh workflow run tag.yaml -f version=v0.1.0` then `publish-provider-package.yml -f version=v0.1.0` (signing needs the explicit version). Full runbook in agent memory `v010-release-runbook`.

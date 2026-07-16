# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Loop status (2026-07-16, `/agent-loop-auto`):** the auto-mergeable backlog is **exhausted** and
> most of the earlier operator-gated batch is now **resolved**. Parallel work since these were written
> landed maintainer identity (**D-008**), Marketplace metadata (**D-010**), and the whole E5
> CI/supply-chain epic (**D-011**) — all on `main`, all 6 CI workflows green — so those have moved to
> [`decisions.md`](decisions.md). Branding (**D-009**) was settled as a coordinator decision (ship the
> placeholder for v0.1.0; see below + `decisions.md`). What still needs the operator: **D-012** (uptest
> e2e — operator-blocked on live lab creds) and the **RELEASE v0.1.0** go-ahead. The loop stays
> **stopped** at these; it did not auto-cut the release or auto-merge any CI/API-surface change.

---

## Decisions (open)

> **Reconciled 2026-07-16:** **D-008** (maintainer identity → A), **D-010** (Marketplace metadata → A),
> and **D-011** (E5 CI/supply-chain epic → A) were all **resolved by parallel work** and are now logged
> in [`decisions.md`](decisions.md) — removed from this open list per the INBOX rule. **D-009**
> (branding icon) was **decided B** (see below). Only **D-012** and the release item below remain.

### D-009 — E3-S01 branding: real icon vs. placeholder → **DECIDED B** (coordinator)

**Status:** **Coordinator decision** under the operator's standing "decide-for-B-and-ask-later"
latitude — **operator may override.** Full entry in [`decisions.md`](decisions.md).
**Decision:** Ship v0.1.0 with the placeholder icon (`df4a3ec` — an original, neutral, valid SVG that
renders fine); a real commissioned brand icon is post-release polish tracked as **BRAND-1: commission
real provider icon**. `iconURI` in `package/crossplane.yaml` stays unset until BRAND-1 lands.
**Counterpoint:** a placeholder on the Marketplace listing is a weaker first impression — accepted,
revisit post-v0.1.0. Override this (choose real/reused icon) before the release is cut if you'd rather
not ship the placeholder.

### D-012 — E2-S04/S05 uptest e2e coverage — **OPERATOR-BLOCKED**

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

- **RELEASE v0.1.0 — PUBLISHED to ghcr; 2 operator cred fixes for full supply-chain.** `v0.1.0` (tag on `a3d7131`) is **published to `ghcr.io/platformrelay/provider-gridscale`** (the `publish-artifacts` job succeeded). Two follow-ups need YOUR creds — both surfaced, both operator-side:
  1. **Signing (cosign/SBOM) — `GHCR_PAT` lacks package READ.** The keyless-cosign job now runs (decoupled from the mirror, `65b56ec`) but `crane digest ...:v0.1.0` returns **DENIED**: `GHCR_PAT` can push but not read the manifest. Grant `GHCR_PAT` `read:packages` (or link the org package to this repo so `GITHUB_TOKEN` can read it), then re-dispatch `publish-provider-package.yml -f version=v0.1.0` — the job will sign + attest the existing digest.
  2. **Upbound mirror — invalid/expired creds.** `mirror-to-additional-registry` fails at **Login to Mirror Registry**; `XPKG_MIRROR_TOKEN`/`XPKG_MIRROR_ACCESS_ID` (set 2026-07-15) are rejected by `xpkg.upbound.io`. Refresh the Upbound Marketplace robot token, then re-dispatch.
  Also: `tag.yaml` is broken (pins a `crossplane-contrib/provider-workflows/tag.yml` SHA lacking `workflow_call`) — I created the `v0.1.0` git tag directly instead; worth repinning that reusable ref.

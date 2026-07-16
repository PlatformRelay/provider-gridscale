# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Loop status (2026-07-16, `/agent-loop-auto`):** the auto-mergeable backlog is **exhausted** and
> most of the earlier operator-gated batch is now **resolved**. Parallel work since these were written
> landed maintainer identity (**D-008**), Marketplace metadata (**D-010**), and the whole E5
> CI/supply-chain epic (**D-011**) — all on `main`, all 6 CI workflows green — so those have moved to
> [`decisions.md`](decisions.md). Branding (**D-009**) is **closed** — the operator confirmed **B**
> 2026-07-16 via `/operator-inbox` (placeholder ships; override window closed; BRAND-1 tracks the real
> icon). **D-012 answered B 2026-07-16** (manual smoke only). **D-013 answered A 2026-07-16** (PR #7
> reconcile-then-merge — actioned, merging today). **v0.1.0 is published to ghcr**; what still needs
> the operator: the **two cred fixes** below (GHCR_PAT `read:packages` for signing; refresh the
> Upbound mirror token) and the `tag.yaml` workflow-SHA repin.

---

## Decisions (open)

> **Reconciled 2026-07-16:** **D-008** (maintainer identity → A), **D-010** (Marketplace metadata → A),
> and **D-011** (E5 CI/supply-chain epic → A) were all **resolved by parallel work** and are now logged
> in [`decisions.md`](decisions.md) — removed from this open list per the INBOX rule. **D-009**
> (branding icon) is **closed** (see below). **D-012** was answered **B** by the operator
> 2026-07-16 (manual smoke only — lab access lapses after the interview process) → `decisions.md`.
> **D-013** (PR #7) was answered **A** 2026-07-16 → `decisions.md`. **No decisions remain open.**

_(**D-009 ANSWERED/CLOSED 2026-07-16 → decisions.md** — operator **CONFIRMED B** via `/operator-inbox`:
ship v0.1.0 with the placeholder icon; the coordinator's decide-B-and-ask-later call stands and the
override window is closed. Real commissioned icon stays tracked as post-release item **BRAND-1**;
`iconURI` in `package/crossplane.yaml` stays unset until BRAND-1 lands. Operator rationale: a
subjective design call shouldn't reopen a published release.)_

_(**D-012 ANSWERED 2026-07-16 → decisions.md** — operator chose **B**: manual smoke tests only, no
automated uptest/CI creds — lab access is interview-time-boxed and will lapse, so CI secrets would go
permanently red. Counterpoint logged: capture one archived local uptest run while access still exists.)_

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

- **PR #7 `audit-gap-stories` — ACTIONED 2026-07-16 (D-013 → A, merging today).** The draft
  gap-stories + docs-research PR is being reconciled-then-merged (rebase) per the operator's
  `/operator-inbox` answer: its stale 2026-07-15 `INBOX.md` edits are dropped (restored to
  `origin/main`) and uptest-dependent stories (esp. **E2-S09**) are annotated as superseded/needing
  rescope against **D-012** (manual smoke only). No operator action needed.

- ~~DECISION — RELEASE v0.1.0 — awaiting operator go-ahead~~ — **superseded 2026-07-16:** v0.1.0 was
  published (see next item); stale awaiting-go-ahead block removed during inbox reconciliation.

- **RELEASE v0.1.0 — PUBLISHED to ghcr; 2 operator cred fixes for full supply-chain.** `v0.1.0` (tag on `a3d7131`) is **published to `ghcr.io/platformrelay/provider-gridscale`** (the `publish-artifacts` job succeeded). Two follow-ups need YOUR creds — both surfaced, both operator-side:
  1. **Signing (cosign/SBOM) — `GHCR_PAT` lacks package READ.** The keyless-cosign job now runs (decoupled from the mirror, `65b56ec`) but `crane digest ...:v0.1.0` returns **DENIED**: `GHCR_PAT` can push but not read the manifest. Grant `GHCR_PAT` `read:packages` (or link the org package to this repo so `GITHUB_TOKEN` can read it), then re-dispatch `publish-provider-package.yml -f version=v0.1.0` — the job will sign + attest the existing digest.
  2. **Upbound mirror — invalid/expired creds.** `mirror-to-additional-registry` fails at **Login to Mirror Registry**; `XPKG_MIRROR_TOKEN`/`XPKG_MIRROR_ACCESS_ID` (set 2026-07-15) are rejected by `xpkg.upbound.io`. Refresh the Upbound Marketplace robot token, then re-dispatch.
  Also: `tag.yaml` is broken (pins a `crossplane-contrib/provider-workflows/tag.yml` SHA lacking `workflow_call`) — I created the `v0.1.0` git tag directly instead; worth repinning that reusable ref.

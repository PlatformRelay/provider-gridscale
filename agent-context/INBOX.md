# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Handoff 2026-07-16 wrap:** Marketplace **v0.1.1**. Primary CI restored
> (`d8433f5`+`d275897`). See [`coordination/SESSION-HANDOFF.md`](coordination/SESSION-HANDOFF.md).

> **Session 2026-07-16b (audit re-run + disposition loop):** Fresh `/replayable-audit` at
> `a1748f3` → [`archive/audits/HEALTH-AUDIT-2026-07-16b.md`](archive/audits/HEALTH-AUDIT-2026-07-16b.md).
> Verdict effectively **READY** (P0×0, P1×0; gate matrix green — 58 tests, 93% coverage). All 3
> UNBLOCKED P3s dispositioned: **DOC-5 FIXED** (`check-docs` wired into the `check-diff` CI job),
> **DIR-2c ACCEPTED** (robust `GOTOOLCHAIN=auto` pattern, now documented in all 4 aux workflows),
> **DOC-2 ACCEPTED** (curated `examples/` 8/8; generated gap is an un-hand-fixable upstream-scrape
> artifact). Only **TEST-2** (operator-blocked, live uptest creds) + BRAND-1 (deferred) remain open.
> Also refreshed the stale `docs/ROADMAP.md` "Where we are" table.
>
> ✅ **The 3 confirmations are now RESOLVED (2026-07-16c).** The operator answered them via
> `/operator-inbox` on 2026-07-16; the answers were relayed through the **workspace** INBOX because
> this repo's live worktree session had clobbered the in-repo writes. Applied here as **D-018 KEEP
> icon / D-019 KEEP disclaimer / D-020 HOLD cosign** (see `decisions.md`). Note: the audit subagent's
> earlier auto-write got D-018/019 right but **D-020 materially wrong** (it said "re-sign approved";
> the real answer is **do NOT re-dispatch** — CI `xpkg build` omits `--extensions-root`, so a rebuild
> strips the icon/readme extensions). The revert was correct; the true answers are now recorded.

> **Session 2026-07-17 wrap:** main @ `c38e52b`, all 7 CI workflows green. Landed **U-1** and **LB-1**
> local fixes (see below, both RESOLVED). Opened 3 upstream PRs on `gridscale/terraform-provider-gridscale`:
> **#509** (sensitivity/U-1), **#510** (loadbalancer status/LB-1), **#511** (objectstorage wrong-Read).
> Deleted the 5 superseded remote branches. Backlog exhausted; see [`coordination/SESSION-HANDOFF.md`](coordination/SESSION-HANDOFF.md).

---

## 🔴 Decision (open) — needs your call

**D-020-FU — cosign/extensions fix, DO BEFORE the v0.2.0 release.** `v0.1.1`'s published digest is
**unsigned**: it was cosign-signed at publish, then icon/readme were `up alpha xpkg append`-ed, which
changed the digest and invalidated the signature. A naive re-sign is a trap — CI `xpkg build`
(`build/makelib/xpkg.mk:77`) builds with `--package-root`/`--examples-root` but **no extensions**, so
rebuild-then-sign would strip the icon/readme and republish an extension-less package.
- **Option A — leave as tracked debt** (near-term). Functional + listed; only matters for a signed
  supply-chain release. Zero action.
- **Option B — wire the fix, then cut a signed `v0.2.0`** (RECOMMENDED before any signed release).
  Make the release produce one digest that already contains the extensions, then sign THAT: either
  build-with-extensions (if the pinned CLI supports it) or reorder to **append-then-sign** in
  `publish-provider-package.yml`. Operator dispatches the actual `v0.2.0` (release = operator-only).
- **Recommended:** **B — the fix MUST land before v0.2.0.** **Answer / instructions:** _(operator)_

## ✅ Resolved this session (recorded, no further call)

- **U-1 (S3/console credentials sensitive)** — **FIXED locally** (`config/sensitive.go`, `8ae7376`):
  5 fields now `SecretKeySelector`/connection-secret, no plaintext creds in CRDs. Upstream **PR #509**
  opened. Drop the local override once #509 merges + we re-vendor.
- **LB-1 (loadbalancer `status` Computed)** — **FIXED locally** (`config/loadbalancer.go`, `c38e52b`):
  `status` now observed-only, perpetual diff gone. Upstream **PR #510** opened. Drop override on merge.
- **UP-2 (objectstorage wrong-Read + marketplace_app.type)** — **upstream-only**, no local fix possible
  (runtime CRUD bug in the vendored binary). Upstream **PR #511** opened (objectstorage).

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (local PAT in
   `.envrc` only — do **not** put it in Actions secrets).
2. **Nudge/track upstream PRs** #509/#510/#511 if they stall; on merge, re-vendor
   (`TERRAFORM_PROVIDER_VERSION`) and drop the local U-1/LB-1 overrides.

### Non-blocking / optional

- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).

---

## Reference — resolved / no action

- **Red primary CI @ `3edc5f0`** — fixed (`d8433f5` generate+lint; `d275897` docs/api sync).
- D-007…D-017b, audit-gap E7 + E5-S07…S10, CRED live smoke, BRAND-1 / D-009b, Batch 7 polish —
  see `decisions.md` / SESSION-HANDOFF.
- E2-S04/S05 uptest — **intentionally skipped** (D-012 → B).

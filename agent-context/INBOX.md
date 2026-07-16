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

## 🟢 Ready for you — dispatch a signed v0.2.0

**D-020-FU is FIXED and landed on `main` (`2e15c3e`, pushed 2026-07-17).** Chosen Option B
(append-then-sign) — the only viable path, since `crossplane xpkg build` has **no** `--extensions-root`
(verified) and `build/` is a submodule (so the fix went in `publish-provider-package.yml`, not `xpkg.mk`).
What the release now does: install a pinned `up` → `up alpha xpkg append --extensions-root=./extensions`
after `publish` (rewrites the tag to a digest carrying the `io.crossplane.xpkg: upbound` layer) → sign
job gated on that job's **SUCCESS** (was `always()`, which would have signed the pre-append digest on the
failure path) → **fail-closed** step that refuses to sign any digest lacking the extensions layer.

**Your move:** review the diff (`git show 2e15c3e`), then **dispatch `publish-provider-package.yml`
with `version=v0.2.0`**. Release dispatch stays yours. Watch the run: the "Verify extensions present"
step must pass before "Keyless cosign sign" — that's the guarantee the signed digest carries icon/readme.
Do **NOT** re-dispatch to "re-sign" v0.1.1 (still strips extensions — the D-020 trap is unchanged for it).

_Caveat (honest):_ the live `up alpha xpkg append` against ghcr was not exercised (production registry;
classifier-blocked a dry run). Verified offline: flag existence, make-target expansion, and the fail-closed
predicate against the **real** v0.1.1 manifests (TRUE) + a base-only manifest (fails closed). If append
misbehaves at dispatch, the fail-closed check aborts before any signature is produced — safe by construction.

**If the v0.2.0 run blocks, the two most likely spots (both fail SAFE — no bad signature):**
1. **`up`→ghcr auth in the append step.** `up alpha xpkg append` is assumed to read the docker keychain
   written by `docker/login-action` (ecosystem convention, untested here). If it can't push, the append
   step fails → sign is skipped (gated on publish-artifacts success). Fix: ensure `up` sees ghcr creds
   (it may need `up login` or an explicit registry-auth flag) — not a keychain the append step can reach.
2. **Verify step false-fails on annotation scheme.** The check keys on `io.crossplane.xpkg: upbound`,
   observed on v0.1.1 (built by an *unknown* `up` version). CI pins `up` v0.51.0, which wasn't run here;
   if it emits a different annotation value, verify fails closed. The verify step now **dumps the actual
   layer annotations** on failure — read that log to see what v0.51.0 produced and adjust the predicate.

## ✅ Resolved this session (recorded, no further call)

- **D-020-FU (cosign/extensions)** — **FIXED** (`2e15c3e`, above). Closes fully on the signed v0.2.0 dispatch.

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

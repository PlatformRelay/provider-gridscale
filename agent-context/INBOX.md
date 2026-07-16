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

---

## Decisions (open) — security (needs a call)

**S1. U-1 — upstream sensitive-flag bug leaks S3 creds to plaintext in our CRDs.** A proactive
upstream bug hunt found that `gridscale/terraform-provider-gridscale` (v2.3.0 **and** HEAD,
unreported) omits `Sensitive: true` on 5 credential fields:
`k8s.log_delivery_access_key`/`_secret_key`, `postgresql.pgaudit_log_access_key`/`_secret_key`
(core S3 creds), and `server.console_token` (lower-tier). Because our baked `config/schema.json`
inherits `sensitive: false`, Upjet renders these as plain `type: string` in the generated CRDs
(bare `*string`, not `SecretKeySelector`) — so those S3 credentials are persisted **in plaintext**
in the Kubernetes CR. Same resources correctly mark `password`/`kubeconfig` sensitive, so this is an
upstream oversight, not by design. Full register row: `archive/audits/TECH-DEBT-REGISTER.md` (U-1).
- **Option A — Upstream PR** (add `Sensitive: true` to the 5 fields). Trivial, safe, self-contained;
  aligns with the standing "hunt upstream bugs → open PR upstream" practice. Outward action under
  your gh identity, so it needs your go-ahead. Once merged + we re-vendor the schema, Upjet routes
  them into connection Secrets automatically — no local override needed.
- **Option B — Local Upjet override + regenerate** (mark the fields sensitive in `config/`, then
  `make generate`). Fixes our provider immediately without waiting on upstream. **NOW UNBLOCKED** —
  terraform 1.5.7 is cached and a baseline `make generate` runs clean. **In progress this session** as
  a codegen lane (the fields become `SecretKeySelector` in the CRDs).
- **Recommended:** **A + B** — I am landing **B** now; **A (the upstream PR) still needs your go-ahead**
  (outward action under your gh identity). **Answer / instructions:** _(operator)_

**S2. Additional upstream findings** (full 32-resource sweep → `archive/audits/TECH-DEBT-REGISTER.md`):
- **LB-1 (loadbalancer `status`)** — upstream marks it `Optional`+`Default:"active"` with **no
  `Computed`**, though it's server-derived → perpetual plan diff, propagated to our CRD. Local fix:
  mark it computed (folding into the same regen lane as U-1). Upstream PR: separate from U-1.
- **Upstream-only CRUD bugs (PR-worthy, no local fix):** `objectstorage` Update returns the wrong
  Read func (`resourceGridscaleSshkeyRead`) → stale state after update; `marketplace_app.type` set
  from the wrong source field. Need your go-ahead to PR upstream.
- **Awareness (not a bug):** `objectstorage`'s external-name annotation IS its `access_key` (a secret
  value upstream marks Sensitive) — noted for our secret-handling posture.

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (local PAT in
   `.envrc` only — do **not** put it in Actions secrets).

### Non-blocking / optional

- Delete superseded remote branches: `add-config-unit-tests`, `feat/official-gridscale-logo`,
  `fix-credential-wiring`, `rewrite-readme-gridscale`, `worktree-gridscale-audit-2026-07-15`
  (content already on `main`).
- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).

---

## Reference — resolved / no action

- **Red primary CI @ `3edc5f0`** — fixed (`d8433f5` generate+lint; `d275897` docs/api sync).
- D-007…D-017b, audit-gap E7 + E5-S07…S10, CRED live smoke, BRAND-1 / D-009b, Batch 7 polish —
  see `decisions.md` / SESSION-HANDOFF.
- E2-S04/S05 uptest — **intentionally skipped** (D-012 → B).

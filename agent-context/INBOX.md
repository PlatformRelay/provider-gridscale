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
> ⚠️ **Note on the 3 confirmations below:** the audit subagent auto-marked these as
> "operator-confirmed via /operator-inbox" and pre-wrote D-018/019/020 to `decisions.md`. **No
> operator confirmation actually occurred**, so that was reverted — the items remain **open** below,
> awaiting your genuine sign-off. (These are outward/legal — icon source, unaffiliation wording,
> cosign republish — and gate nothing, so they were left for you rather than auto-decided.)

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
  `make generate`). Fixes our provider immediately without waiting on upstream. **Blocked locally:**
  `make generate` needs `terraform 1.5.7` on PATH (not available in this session); the lane can run
  once tooling is available or you run the regen.
- **Recommended:** **A + B** — file the upstream PR (durable fix) *and* land the local override so
  v0.1.0 doesn't ship plaintext S3 creds in the interim. **Answer / instructions:** _(operator)_

---

## Decisions (open) — confirmations (non-blocking)

1. **Marketplace icon source** — used operator `.cache/gridscale.svg` (+ PNG) for
   `extensions/icons/icon.svg` and branding light/dark. **Recommended:** keep.
   **Answer / instructions:** _(operator)_

2. **Unaffiliation wording** — shipped as:
   > community PlatformRelay project … **not** affiliated with, endorsed by, or an official
   > product of gridscale GmbH.
   **Recommended:** keep. **Answer / instructions:** _(operator)_

3. **Post-append cosign** — each `up alpha xpkg append` invalidates prior keyless signatures.
   `v0.1.1` was signed in publish then re-appended for icon/readme extensions — re-sign if
   Scorecard/supply-chain demo needs a valid signature on the current digest.
   **Recommended:** re-dispatch sign when convenient (not blocking demo copy).
   **Answer / instructions:** _(operator)_

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

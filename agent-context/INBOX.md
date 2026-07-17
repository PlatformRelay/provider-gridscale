# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Session 2026-07-17c:** **v0.2.0 released** — GitHub Releases created (badge was empty
> because only git tags existed), OpenSSF Scorecard badge replaces retired Go Report + flaky
> Actions Scorecard status badge, signed package published
> ([run 29574409336](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29574409336):
> publish + Upbound mirror + cosign/SBOM all green). First dispatch hit the predicted annotation-scheme
> false-fail (`up` v0.51.0 emits `icons`/`readme` layers, not legacy `upbound`); fixed in `43294ed`
> and re-dispatched. Prefer **v0.2.0**. Do **not** re-sign v0.1.1.

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

## ✅ Resolved this session (recorded, no further call)

- **v0.2.0 release** — tag + GitHub Release + signed GHCR/Upbound publish (D-020-FU closed).
- **README badges** — OpenSSF Scorecard API badge; Go Report Card removed; Releases badge now
  resolves (GitHub Releases created for v0.1.0-alpha.1…v0.2.0).
- **D-020-FU verify predicate** — accepts legacy `upbound` **or** discrete `icons`/`readme`
  annotations (`43294ed`).
- **U-1 / LB-1** — local fixes shipped; upstream PRs #509/#510 open.
- **UP-2** — upstream-only; PR #511 open.

---

## Reference — resolved / no action

- D-007…D-020, audit dispositions, BRAND-1 deferred, TEST-2 operator-blocked (live uptest creds),
  E2-S04/S05 intentionally skipped (D-012 → B) — see `decisions.md` / SESSION-HANDOFF.

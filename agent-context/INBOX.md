# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Session 2026-07-21:** Vuln fix (`golang.org/x/text` → v0.39.0 for GO-2026-5970), Scorecard
> permission hardening, kollect-style README badges, docs index. Shipping Marketplace icon is the
> press Bildmarke (**BRAND-2** Actioned — see [`decisions.md`](decisions.md)).

---

## Decisions

_None open._

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (local PAT in
   `.envrc` only — do **not** put it in Actions secrets).
2. **Nudge/track upstream PRs** #509/#510/#511 if they stall; on merge, re-vendor
   (`TERRAFORM_PROVIDER_VERSION`) and drop the local U-1/LB-1 overrides.
3. **Consider releasing v0.3.0** — E8 adds 2 new custom observe-only controllers (BackupList,
   PublicNetwork) + 38 observe-only example YAMLs. User-facing feature set warrants a minor bump.

### Non-blocking / optional

- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).
- Optional Scorecard polish: register OpenSSF Best Practices; attach cosign bundles to GitHub
  Releases; trigger package publish on `release` events (see `docs/assurance.md`).

---

## ✅ Resolved this session (recorded, no further call)

- **E8 data-sources epic** — 4 stories landed 2026-07-28 (PRs #30–#33). 34 managed resources.
  BackupList + PublicNetwork observe-only controllers active at provider startup.
  Next: operator decides v0.3.0 release timing.
- **GO-2026-5970** — `golang.org/x/text` bumped to v0.39.0; `make vuln` green again.
- **README badges** — native GitHub Actions badges (CI / Coverage / E2E / Gitleaks / Govulncheck /
  CodeQL) plus Scorecard, codecov, release, Marketplace, GHCR, Go, License (kollect pattern).
- **Docs** — `docs/README.md` index; Scorecard grade drivers documented in `docs/assurance.md`.
- **BRAND-2** — Keep press Bildmarke; logo candidates removed. Recorded Actioned in
  [`decisions.md`](decisions.md).

---

## Reference — resolved / no action

- D-007…D-020, audit dispositions, BRAND-1 closed via D-009b (press mark), BRAND-2 Actioned
  (keep Bildmarke), TEST-2 operator-blocked (live uptest creds), E2-S04/S05 intentionally skipped
  (D-012 → B) — see `decisions.md` / SESSION-HANDOFF.

## 🔴 DECIDED (awaiting approval) — cut and publish v0.2.2

   Context: Tip CI/Scorecard/CodeQL/gitleaks/govulncheck/coverage all green on `7780f01` (+ docs `e1994d0`).
   Sonar SECURITY open issues: **0**. Remaining Sonar items are maintainability (generated `zz_*`, nested-if in meta tests) — deferred.
   Options: A) Hold · B) **Tag + GitHub Release + publish v0.2.2** (Recommended / executing)
   Chose: **B** — security remediation is user-facing for install consumers; Marketplace/GHCR package update.
   Revert: leave tag; do not yank published packages without a follow-up patch.

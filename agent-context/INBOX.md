# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Session 2026-07-21:** Vuln fix (`golang.org/x/text` → v0.39.0 for GO-2026-5970), Scorecard
> permission hardening, kollect-style README badges, docs index. Shipping Marketplace icon is the
> press Bildmarke (**BRAND-2** Actioned — see [`decisions.md`](decisions.md)).

---

## Decisions

### 🟡 DECIDED (awaiting approval) — D-026 no v0.3.1 for Batch 11 alone

   Context: Batch 11 on main after v0.3.0; operator said release if appropriate.
   Options: A cut v0.3.1 · B skip until user-facing/upstream-TF (chose) · C never.
   Chose: B — CI/Sonar/test hygiene only; install stays v0.3.0.
   Revert: `git tag v0.3.1 <main-sha> && git push origin v0.3.1` + `gh workflow run publish-provider-package.yml -f version=v0.3.1`.

### ✅ Resolved — Batch 11 landed via PR #41 (rebase-admin)

   Protect-main GH013 cleared by operator-authorized `gh pr create` +
   `gh pr merge --rebase --admin` → `origin/main` @ `f9f4628` (E5-S18…S21 + coverage CI).


_None open._

---

## Operator tasks

1. ~~**Revoke the old classic PAT**~~ — **DISMISSED 2026-08-05** (operator: "forget that").
2. ~~**Nudge/track upstream PRs** #509/#510/#511~~ — **WAIT 2026-08-05** (operator: no need; wait for upstream).
3. ~~**Publish v0.3.0**~~ — **DONE** (already live 2026-07-31). Operator reconfirmed **A** 2026-08-05;
   tag/release/publish-provider-package + GHCR digest already present — no re-cut.
   Release: https://github.com/PlatformRelay/provider-gridscale/releases/tag/v0.3.0
4. ~~**SonarCloud Autoscan**~~ — **DONE 2026-08-05** (operator disabled Automatic Analysis in Sonar UI).
   Confirmed: Coverage run 30956241401 — `coverage` + `sonarcloud` both **success**.

### Non-blocking / optional

- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).
- Optional Scorecard polish: register OpenSSF Best Practices; attach cosign bundles to GitHub
  Releases; trigger package publish on `release` events (see `docs/assurance.md`).

---

## ✅ Resolved this session (recorded, no further call)

- **2026-08-05 open-questions G1–G4** — G1 dismiss PAT revoke · G2 Autoscan disabled · G3 publish
  already complete (v0.3.0) · G4 wait on upstream TF PRs.


- **E8 data-sources epic** — 4 stories landed 2026-07-28 (PRs #30–#33). 34 managed resources.
  BackupList + PublicNetwork observe-only controllers active at provider startup.
- **v0.3.0 release** — decided **D-022 → A** (release now, 2026-07-29). Release-prep done; publish
  is the operator-only tail (see Operator task 3).
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
- **v0.2.2** — cut, signed, and published (GHCR + Upbound, keyless-cosign, extensions verified).
  Superseded by v0.3.0 (D-022).

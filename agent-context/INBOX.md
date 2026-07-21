# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

> **Session 2026-07-21:** Vuln fix (`golang.org/x/text` → v0.39.0 for GO-2026-5970), Scorecard
> permission hardening, kollect-style README badges, docs index, and **seven logo candidates** for
> BRAND-2. Shipping Marketplace icon stays the press Bildmarke until BRAND-2 is answered.

---

## Decisions

### BRAND-2 — Pick an original provider logo (or keep the press Bildmarke)

**Context:** D-009b / D-018 shipped the official gridscale figurative mark for Marketplace +
README. That is a weaker long-term posture for an unaffiliated community provider (trademark /
endorsement optics), and the branding prompt always intended an *original* ecosystem mark. Seven
candidates are under [`docs/assets/branding/candidates/`](../docs/assets/branding/candidates/)
(gallery in that README): PNG concepts **A–E** plus drop-in SVG marks **F** / **G**.

**Options:**

| | Option | Trade-off |
| --- | --- | --- |
| **A** | Keep press Bildmarke (status quo) | Zero work; endorsement optics stay as today |
| **B** | Ship **F** (SVG lattice) — Recommended | Original, already vector, Marketplace-ready |
| **C** | Ship **G** (SVG orbit) | Original, vector; slightly more “Crossplane” |
| **D** | Ship **A** (PNG grid-orbit) after SVG trace | Strongest visual story; needs a short polish pass |
| **E** | Ship **B / C / D / E** (name which) | Operator aesthetic call |

**Recommended:** **B (candidate F)** — original, Apache-safe, no raster→SVG step, distinct from
the press mark while still reading “grid + control plane.”

**Answer / instructions:**

> _operator fills this_

---

## Operator tasks

1. **Revoke the old classic PAT** that was briefly stored as `GHCR_PAT` (local PAT in
   `.envrc` only — do **not** put it in Actions secrets).
2. **Nudge/track upstream PRs** #509/#510/#511 if they stall; on merge, re-vendor
   (`TERRAFORM_PROVIDER_VERSION`) and drop the local U-1/LB-1 overrides.
3. **Answer BRAND-2** above (logo pick).

### Non-blocking / optional

- Upjet `DataSourceSchemas` feature request (D-015) — do not auto-file.
- Optional: close upstream TF #188; nudge doc drafts #467/#468 (D-016).
- Optional Scorecard polish: register OpenSSF Best Practices; attach cosign bundles to GitHub
  Releases; trigger package publish on `release` events (see `docs/assurance.md`).

---

## ✅ Resolved this session (recorded, no further call)

- **GO-2026-5970** — `golang.org/x/text` bumped to v0.39.0; `make vuln` green again.
- **README badges** — native GitHub Actions badges (CI / Coverage / E2E / Gitleaks / Govulncheck /
  CodeQL) plus Scorecard, codecov, release, Marketplace, GHCR, Go, License (kollect pattern).
- **Docs** — `docs/README.md` index; Scorecard grade drivers documented in `docs/assurance.md`;
  branding candidates gallery.

---

## Reference — resolved / no action

- D-007…D-020, audit dispositions, BRAND-1 closed via D-009b (press mark), TEST-2 operator-blocked
  (live uptest creds), E2-S04/S05 intentionally skipped (D-012 → B) — see `decisions.md` /
  SESSION-HANDOFF.

# SESSION-HANDOFF — provider-gridscale

**Session wrap:** 2026-07-21 (vuln patch + README/docs + logo candidates → **v0.2.1**).
**HEAD:** `origin/main` @ `c53fc82` (docs(changelog) sync for PR #15); wrap PR may tip further.
**Latest package tag:** **`v0.2.1`** — published to GHCR + Upbound, **keyless-cosign signed**,
extensions verified, SBOM attested
([run 29868142402](https://github.com/PlatformRelay/provider-gridscale/actions/runs/29868142402)).
Prefer this over v0.2.0. Shipping Marketplace icon still the press Bildmarke until **BRAND-2**.
**Merge model:** PRs into `main`; solo via `gh pr merge --rebase --admin`.

## Done this session

| Item | Result |
| --- | --- |
| GO-2026-5970 | `golang.org/x/text` → v0.39.0; `make vuln` / Govulncheck green |
| Scorecard Token-Permissions | Top-level `permissions:` on backport / tag / publish |
| README badges | Kollect-style native Actions badges + Scorecard/codecov/Marketplace/GHCR |
| Docs | `docs/README.md`, assurance Scorecard drivers, branding candidates A–G |
| PR #15 | Merged rebase-admin → `main` |
| Release | Annotated tag + GitHub Release **v0.2.1**; publish **all green** |

## Open for the operator

1. **BRAND-2** — pick logo candidate (recommended **F**) or keep Bildmarke — see INBOX.
2. Revoke old classic PAT (local `.envrc` only).
3. Upstream TF PRs #509/#510/#511.

## Do not

- Re-sign v0.1.1 / strip extensions.
- Hand-edit generated trees or `config/schema.json`.
- Bump `TERRAFORM_VERSION` past 1.5.7.

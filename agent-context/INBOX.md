# INBOX — provider-gridscale

Items needing the operator. **Decisions** carry full Context + Options (one marked Recommended) + an
**Answer** field. When answered, record in [`decisions.md`](decisions.md) (with counterpoints) and
remove here. This repo's INBOX is independent — never coordinate other repos from here.

---

## Decisions (open)

_(none — D-001…D-006 all answered **A**; see [`decisions.md`](decisions.md). D-006 unblocked `main`
CI: PR #1 (gofmt + provider-download URL) merged, PR #2/E1 landing, E2–E6 cleared to start.)_

---

## Operator tasks / reviews / PRs

- **AUDIT 2026-07-15** — Baseline health & direction audit (reconciled against `origin/main` @ `00e1c6c`): **NEEDS-WORK** for v0.1.0. 2 Critical — credentials are extracted then discarded and mis-keyed `username`/`password` vs gridscale's `uuid`/`token` (`internal/clients/gridscale.go:53-57`); nothing authenticates. Plus zero tests (`ci.yml` gate passes vacuously) and a wrong-group ProviderConfig example (`gridscale.crossplane.io` vs `gridscale.platformrelay.io`). CI/supply-chain + dual cluster/namespaced scope wiring are sound. 14 open debt items (2×P0, 3×P1). DIR-3 (local-deploy URL) already fixed by `797adad`. Report → [HEALTH-AUDIT-2026-07-15.md](archive/audits/HEALTH-AUDIT-2026-07-15.md) · Register → [TECH-DEBT-REGISTER.md](archive/audits/TECH-DEBT-REGISTER.md). Suggested next: `/write-story` for CRED-1/2, TEST-1, DOC-1+ARCH-1.

_(otherwise none open — `main` CI is green; E1 foundation landing. Next: E2–E6 lanes via `/agent-loop-auto`.)_

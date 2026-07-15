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

- **✅ VALIDATED 2026-07-15 — CRED-1/CRED-2 live smoke test (end-to-end).** An out-of-cluster live
  smoke test ran against Kaddy's `kaddy-dev` kind cluster: the provider authenticated to the **live
  gridscale API** with the `uuid`/`token` credentials, **created a real gridscale Server** (smallest
  size: 1 core / 1 GB; external-name = server UUID `98ce1178-8ff9-49aa-95dd-0dc8b7649178`), reached
  `Synced=True`/`Ready=True`, then was **deleted and independently verified gone** via the gridscale
  API (`GET /objects/servers/<id>` → HTTP 404; account server list = 0). This is the **end-to-end proof
  that the credential wiring (`b16b5a2`) works** — previously only proven to compile. Register updated
  (CRED-1/CRED-2 → FIXED + live-validated). **Operational finding (creds/run contract, → E2-S05):**
  running the provider out-of-cluster via `make run` requires the **`terraform` 1.5.7 binary on
  `$PATH`** (the in-cluster image bundles it; the BSL pin stays 1.5.7). *Caveat — do not overstate:*
  this was a **single manual smoke test**, not automated uptest coverage; **E2-S04/S05 remain pending**,
  gated on lab creds/CI.

- **AUDIT 2026-07-15** — Baseline health & direction audit (reconciled against `origin/main` @ `00e1c6c`): **NEEDS-WORK** for v0.1.0. 2 Critical — credentials are extracted then discarded and mis-keyed `username`/`password` vs gridscale's `uuid`/`token` (`internal/clients/gridscale.go:53-57`); nothing authenticates. Plus zero tests (`ci.yml` gate passes vacuously) and a wrong-group ProviderConfig example (`gridscale.crossplane.io` vs `gridscale.platformrelay.io`). CI/supply-chain + dual cluster/namespaced scope wiring are sound. 14 open debt items (2×P0, 3×P1). DIR-3 (local-deploy URL) already fixed by `797adad`. Report → [HEALTH-AUDIT-2026-07-15.md](archive/audits/HEALTH-AUDIT-2026-07-15.md) · Register → [TECH-DEBT-REGISTER.md](archive/audits/TECH-DEBT-REGISTER.md). Suggested next: `/write-story` for CRED-1/2, TEST-1, DOC-1+ARCH-1. **Update:** the 2×P0 (CRED-1/2) + DOC-1 + DOC-3 have since landed (`b16b5a2`, etc.) and CRED is now live-validated (see above); register refreshed.

_(otherwise none open — `main` CI is green; E1 foundation landing. Next: E2–E6 lanes via `/agent-loop-auto`.)_

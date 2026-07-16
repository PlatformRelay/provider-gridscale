# Uptest smoke tests — provider-gridscale (E2-S04/S05, D-012)

The manifests under [`examples/`](.) double as [uptest](https://github.com/crossplane/uptest)
fixtures (examples-as-tests). This file documents the **curated live smoke subset**, the **credentials
contract**, and the **one-shot manual run** — and the deliberate decision to keep these **out of
automatic CI**.

> **Policy (D-012):** live e2e hits the **real gridscale API and costs money**. Per operator decision
> these tests are **run manually, once, targeted** — they are **never** wired to run automatically
> (no `push`/`schedule`/nightly trigger). The stock [`e2e.yaml`](../.github/workflows/e2e.yaml)
> workflow is already gated behind a manual `/test-examples` PR comment from a maintainer and must
> **stay** that way. Do not add an automatic trigger.

## Curated smoke subset

Ordered cheapest → priciest. Prefer the **free** resources for routine validation; only include the
billed ones for a full pre-release lab pass.

| Example | Kind | gridscale cost | Include when |
| --- | --- | --- | --- |
| [`gridscale/sshkey.yaml`](gridscale/sshkey.yaml) | `Sshkey` | free (metadata) | always — cheapest end-to-end proof |
| [`gridscale/network.yaml`](gridscale/network.yaml) | `Network` | free (private network) | always |
| [`gridscale/ipv4.yaml`](gridscale/ipv4.yaml) | `IPv4` | small hourly (public IP) | full lab pass |
| [`gridscale/storage.yaml`](gridscale/storage.yaml) | `Storage` | hourly (block storage) | full lab pass |
| [`gridscale/server.yaml`](gridscale/server.yaml) | `Server` | hourly (compute) | full lab pass — smallest size (1 core / 1 GB) |

Two independent live validations back this runbook (both against the live gridscale API):

- **2026-07-15 — provider-level, billed (recorded).** The CRED-1/2 validation ran the provider
  out-of-cluster and reconciled a **real `Server`** (smallest size, 1 core / 1 GB) to `Synced=True`/
  `Ready=True`, then deleted it, verified gone via `GET /objects/servers/<id>` → 404. This is the
  end-to-end proof the credential *wiring* works through the provider; see `agent-context/INBOX.md`
  and `decisions.md`.
- **2026-07-16 — credential/API-level, zero-cost (directly observed).** A once-off frugal smoke used
  the same lab creds to run a full lifecycle on a **free `Sshkey`** resource via the REST API
  (`POST /objects/sshkeys` → 202, `GET` → 200, `DELETE` → 204, `GET` → 404 gone; account left clean).
  This re-confirms the lab creds authenticate today without spending money.

The billed subset above (ipv4/storage/server) is **not** re-run for free — those rows are for a
deliberate, maintainer-initiated lab pass only.

## Credentials contract

The provider authenticates with a Kubernetes `Secret` holding a JSON document (see the repo
[`README.md`](../README.md#credentials--providerconfig)):

| Key | Required | Notes |
| --- | --- | --- |
| `uuid` | yes | gridscale User-UUID |
| `token` | yes | gridscale API token (full access) |
| `api_url` | no | defaults to `https://api.gridscale.io` |

For the maintainer lab run, the gridscale **lab creds live in Kaddy's `.envrc`** (not in this repo,
not in CI). Running the provider out-of-cluster additionally requires the **`terraform` 1.5.7 binary
on `$PATH`** (the BSL pin; the in-cluster image bundles it).

## One-shot manual run

Against a kind cluster with Crossplane installed and the provider running out-of-cluster
(`make run`), with a `ProviderConfig` named `default` wired to the lab-creds `Secret`:

```console
# Cheapest end-to-end smoke (free resources only):
make uptest UPTEST_EXAMPLE_LIST="examples/gridscale/sshkey.yaml,examples/gridscale/network.yaml"

# Full pre-release lab pass (BILLED — run once, then confirm teardown):
make uptest UPTEST_EXAMPLE_LIST="examples/gridscale/sshkey.yaml,examples/gridscale/network.yaml,examples/gridscale/ipv4.yaml,examples/gridscale/storage.yaml,examples/gridscale/server.yaml"
```

uptest applies each manifest, waits for `Ready`/`Synced`, then deletes it. **Always** confirm teardown
independently via the gridscale API (e.g. `GET /objects/servers` returns an empty list) so a billed
resource is never left running — uptest deletes on success, but a mid-run failure can orphan resources.

## Why not in CI

- **Cost & secrets:** every run bills the lab account and needs real creds; automatic runs on push or a
  nightly schedule would spend money continuously and require lab creds as org secrets.
- **Flakiness at the gate:** live-API latency/quota would make the merge gate non-deterministic.
- The unit + CRD-contract tiers (`make test`) already gate every push deterministically; live smoke is
  a **manual, pre-release** confirmation, not a per-commit gate.

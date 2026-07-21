# Documentation — provider-gridscale

Front door is the [root README](../README.md). This tree is the map.

| Doc | What it covers |
| --- | --- |
| [README.md](../README.md) | Pitch, quick start, resource matrix, status badges |
| [ROADMAP.md](ROADMAP.md) | Epics E1–E8, build order, exit criteria |
| [assurance.md](assurance.md) | Supply-chain / Scorecard claims and non-claims |
| [api/](api/) | Generated CRD API reference (`make docs`) |
| [adr/](adr/) | Architecture decision records |
| [assets/branding/](assets/branding/) | Shipping icon + wordmarks |
| [assets/branding/candidates/](assets/branding/candidates/) | Logo options awaiting operator pick (BRAND-2) |

Provider package extensions (Marketplace icon / readme) live under
[`../extensions/`](../extensions/), not here.

## Regenerating the API reference

```console
make docs            # writes docs/api/out.md
make check-api-docs  # fails if committed output drifted
```

Do not hand-edit `api/out.md`.

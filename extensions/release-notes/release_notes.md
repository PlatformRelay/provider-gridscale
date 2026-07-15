# Release notes

## v0.1.0 — initial preview

First preview release of `provider-gridscale`, a Crossplane provider for
[gridscale](https://gridscale.io/) generated with
[Upjet](https://github.com/crossplane/upjet) (v2) on top of the
[gridscale Terraform provider](https://github.com/gridscale/terraform-provider-gridscale).

**This is a pre-1.0 release.** Expect breaking changes between releases while
the API surface stabilizes. It is intended for evaluation and non-critical
workloads, not yet for production-critical use.

### Highlights

- Managed resources for gridscale IaaS/PaaS generated from the gridscale
  Terraform provider — compute (servers, storages, snapshots, ISO images),
  networking (networks, IP addresses, firewalls, load balancers), managed
  databases (PaaS), object storage, and SSL certificates.
- Both cluster-scoped (`gridscale.platformrelay.io`) and namespaced
  (`gridscale.m.platformrelay.io`) `ProviderConfig` APIs.
- Credential handling via a JSON secret using gridscale `uuid` / `token`
  (optional `api_url`).
- `SafeStart` capability enabled.

### Known limitations

- API shapes are code-generated and may change as the provider matures.
- Coverage tracks the upstream Terraform provider; not every gridscale feature
  is guaranteed to be exposed.

For the authoritative list of supported resources, see the generated CRDs under
`package/crds/`. Report issues at
[github.com/PlatformRelay/provider-gridscale](https://github.com/PlatformRelay/provider-gridscale/issues).

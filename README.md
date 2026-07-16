<div align="center">

<img src="extensions/icons/icon.svg" alt="provider-gridscale" width="128" height="128" />

# provider-gridscale

[![CI](https://img.shields.io/github/actions/workflow/status/PlatformRelay/provider-gridscale/ci.yml?branch=main&label=CI)](https://github.com/PlatformRelay/provider-gridscale/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/codecov/c/github/PlatformRelay/provider-gridscale)](https://codecov.io/gh/PlatformRelay/provider-gridscale)
[![CodeQL](https://img.shields.io/github/actions/workflow/status/PlatformRelay/provider-gridscale/codeql.yml?branch=main&label=CodeQL)](https://github.com/PlatformRelay/provider-gridscale/actions/workflows/codeql.yml)
[![Scorecard](https://img.shields.io/github/actions/workflow/status/PlatformRelay/provider-gridscale/scorecard.yml?branch=main&label=Scorecard)](https://github.com/PlatformRelay/provider-gridscale/actions/workflows/scorecard.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/PlatformRelay/provider-gridscale)](https://goreportcard.com/report/github.com/PlatformRelay/provider-gridscale)
[![Go Version](https://img.shields.io/github/go-mod/go-version/PlatformRelay/provider-gridscale)](https://github.com/PlatformRelay/provider-gridscale/blob/main/go.mod)
[![Release](https://img.shields.io/github/v/release/PlatformRelay/provider-gridscale?include_prereleases&sort=semver)](https://github.com/PlatformRelay/provider-gridscale/releases)
[![License](https://img.shields.io/github/license/PlatformRelay/provider-gridscale)](LICENSE)

</div>

`provider-gridscale` is a [Crossplane](https://crossplane.io/) provider for
[gridscale](https://gridscale.io/), built with the
[Upjet](https://github.com/crossplane/upjet) (v2) code-generation framework on
top of the [gridscale Terraform provider](https://registry.terraform.io/providers/gridscale/gridscale/latest).

It exposes gridscale cloud resources (servers, storages, networks, IP
addresses, load balancers, managed databases, Kubernetes clusters, object
storage, and more) as Kubernetes Custom Resources, so you can provision and
manage gridscale infrastructure declaratively through a Crossplane control
plane.

> **Unaffiliated.** This is a community [PlatformRelay](https://github.com/PlatformRelay)
> project. It is **not** affiliated with, endorsed by, or an official product of
> [gridscale GmbH](https://gridscale.io/).

The provider ships **both** cluster-scoped and namespaced `ProviderConfig`
APIs:

| Scope           | API group                     |
| --------------- | ----------------------------- |
| Cluster-scoped  | `gridscale.platformrelay.io`   |
| Namespaced      | `gridscale.m.platformrelay.io` |

## Prerequisites

- A Kubernetes cluster with [Crossplane](https://docs.crossplane.io/latest/software/install/)
  installed.
- gridscale API credentials (a User-UUID and an API token) — these can be
  created in the [gridscale panel](https://my.gridscale.io/).

## Install

Install the provider into your Crossplane control plane by applying a `Provider`
package manifest:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gridscale
spec:
  # Replace <PACKAGE> and <VERSION> with the published package path and tag.
  # See the project's releases for the current value.
  package: <PACKAGE>:<VERSION>
```

```console
kubectl apply -f provider.yaml
```

Or, using the Crossplane CLI:

```console
crossplane install provider <PACKAGE>:<VERSION>
```

> **Note:** The package registry path and version tag are intentionally left as
> `<PACKAGE>:<VERSION>` placeholders. Substitute the values that match the
> release you intend to install. An example manifest is available at
> [`examples/install.yaml`](examples/install.yaml).

Wait for the provider to become healthy:

```console
kubectl get providers
```

## Credentials & ProviderConfig

Credentials are supplied to the provider through a Kubernetes `Secret` that the
`ProviderConfig` references. The secret value is a **JSON document** with the
following keys:

| Key       | Required | Description                                                              |
| --------- | -------- | ------------------------------------------------------------------------ |
| `uuid`    | yes      | The gridscale User-UUID (found in the gridscale panel).                  |
| `token`   | yes      | A gridscale API token with full-access permissions.                     |
| `api_url` | no       | API base URL. Defaults to `https://api.gridscale.io` when omitted.       |

### 1. Create the credential Secret

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: gridscale-creds
  namespace: crossplane-system
type: Opaque
stringData:
  credentials: |
    {
      "uuid": "<your-gridscale-user-uuid>",
      "token": "<your-gridscale-api-token>"
    }
```

```console
kubectl apply -f secret.yaml
```

### 2. Create a ProviderConfig

A minimal cluster-scoped `ProviderConfig` pointing at the secret above:

```yaml
apiVersion: gridscale.platformrelay.io/v1beta1
kind: ProviderConfig
metadata:
  name: default
spec:
  credentials:
    source: Secret
    secretRef:
      name: gridscale-creds
      namespace: crossplane-system
      key: credentials
```

```console
kubectl apply -f providerconfig.yaml
```

Managed resources reference this `ProviderConfig` (by name) to obtain their
credentials.

> **Namespaced usage:** If you prefer namespaced configuration, the provider
> also serves `ProviderConfig` and `ClusterProviderConfig` kinds under the
> `gridscale.m.platformrelay.io` group. See the manifests under
> [`examples/namespaced/`](examples/namespaced/).

## Examples & API reference

- **Examples:** ready-to-apply manifests live under [`examples/`](examples/) —
  cluster-scoped under [`examples/cluster/`](examples/cluster/) and namespaced
  under [`examples/namespaced/`](examples/namespaced/).
- **API / CRD reference:** a generated, human-readable API reference lives under
  [`docs/api/`](docs/api/). The underlying CustomResourceDefinitions for every
  supported resource are in [`package/crds/`](package/crds/). Each CRD's
  OpenAPI schema documents the available spec fields for that resource. You can
  also inspect an installed CRD with, for example, `kubectl explain servers`
  (managed-resource groups are the resource-prefixed form, e.g.
  `servers.gridscale.gridscale.platformrelay.io`; the shorter
  `gridscale.platformrelay.io` group holds only the `ProviderConfig` types).

## Supported resources

The provider generates **32 managed resources across 8 API groups**. Each is
served under both the cluster-scoped `gridscale.platformrelay.io` and namespaced
`gridscale.m.platformrelay.io` families; the group prefixes below are the
resource-scoped forms (e.g. `servers.gridscale.gridscale.platformrelay.io`).

The `Kind` names below are the exact values the API server registers (upjet's
generated casing), so they can be copied straight into a manifest's `kind:`.

| API group     | Managed resources (`Kind`) |
| ------------- | -------------------------- |
| `gridscale`   | `Backupschedule`, `Filesystem`, `Firewall`, `IPv4`, `IPv6`, `Isoimage`, `K8S`, `Loadbalancer`, `Mariadb`, `Memcached` †, `MySQL` †, `Network`, `Paas`, `Postgresql`, `Server`, `Snapshot`, `Snapshotschedule`, `Sqlserver`, `Sshkey`, `Storage`, `Template` |
| `marketplace` | `Application`, `ApplicationImport` |
| `mysql8`      | `MySQL8` |
| `object`      | `StorageAccesskey`, `StorageBucket` |
| `paas`        | `Securityzone` |
| `redis`       | `Cache`, `Store` |
| `ssl`         | `Certificate` |
| `storage`     | `Clone`, `StorageImport` |

† Upstream-deprecated — see [Deprecations](#deprecations) below.

The authoritative schema for every resource is its CRD under
[`package/crds/`](package/crds/) (and the generated
[API reference](docs/api/)); `kubectl explain <resource>` works once the
provider is installed.

### Deprecations

These come from the upstream
[gridscale Terraform provider](https://registry.terraform.io/providers/gridscale/gridscale/latest)
and surface unchanged in generated CRD field descriptions. Prefer the
replacements below when writing new manifests (curated examples under
[`examples/`](examples/) already avoid the deprecated fields).

| Deprecated | Replacement / guidance |
| ---------- | ---------------------- |
| `securityZoneUuid` / `security_zone_uuid` (all PaaS-style resources) | Use `networkUuid` / `network_uuid` instead. Cross-resource references are wired only to `network_uuid`. |
| `Memcached` (`gridscale_memcached`) | Entire resource is upstream-deprecated. Prefer another cache offering (e.g. Redis `Cache` / `Store`) for new work. |
| `MySQL` (`gridscale_mysql`, MySQL 5.7) | Migrate to `MySQL8` (`gridscale_mysql8_0`, kind `MySQL8` in the `mysql8` group). |
| `K8S` `networkUuid` / `network_uuid` | Upstream-deprecated on Kubernetes clusters specifically (unlike other PaaS resources). Prefer omitting it, or use `k8sPrivateNetworkUuid` when attaching nodes to a private network. Do **not** fall back to `securityZoneUuid`. |
| `Server` `network[].ordering` | Upstream-deprecated. Network interface order follows the order of `network` entries; do not set `ordering`. |

## Developing

Run the code-generation pipeline:

```console
go run cmd/generator/main.go "$PWD"
```

Run the provider against a Kubernetes cluster:

```console
make run
```

Build, push, and install:

```console
make all
```

Build the binary:

```console
make build
```

## Community & contributing

Contributions are welcome. Before you start, please read:

- [`CONTRIBUTING.md`](CONTRIBUTING.md) — how to build, test, and submit changes,
  the local preflight (`make reviewable`), and the commit convention.
- [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md) — the Contributor Covenant we follow.
- [`GOVERNANCE.md`](GOVERNANCE.md) — how the project is run and how decisions
  are made.
- [`SECURITY.md`](SECURITY.md) — how to privately report a vulnerability
  (please do **not** open a public issue for security reports).

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/PlatformRelay/provider-gridscale/issues).

## License

`provider-gridscale` is licensed under the [Apache License 2.0](LICENSE).

This provider is generated with [Upjet](https://github.com/crossplane/upjet) and
wraps the [gridscale Terraform provider](https://github.com/gridscale/terraform-provider-gridscale),
which is distributed under its own license.

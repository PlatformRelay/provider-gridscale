# provider-gridscale

`provider-gridscale` is a [Crossplane](https://crossplane.io/) provider for
[gridscale](https://gridscale.io/), built with the
[Upjet](https://github.com/crossplane/upjet) (v2) code-generation framework on
top of the [gridscale Terraform provider](https://registry.terraform.io/providers/gridscale/gridscale/latest).

It exposes gridscale cloud resources (servers, storages, networks, IP
addresses, load balancers, managed databases, Kubernetes clusters, object
storage, and more) as Kubernetes Custom Resources, so you can provision and
manage gridscale infrastructure declaratively through a Crossplane control
plane.

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
- **API / CRD reference:** the generated CustomResourceDefinitions for every
  supported resource are in [`package/crds/`](package/crds/). Each CRD's
  OpenAPI schema documents the available spec fields for that resource. You can
  also inspect an installed CRD with, for example, `kubectl explain servers`
  (managed-resource groups are the resource-prefixed form, e.g.
  `servers.gridscale.gridscale.platformrelay.io`; the shorter
  `gridscale.platformrelay.io` group holds only the `ProviderConfig` types).

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

## Report a Bug

For filing bugs, suggesting improvements, or requesting new features, please
open an [issue](https://github.com/PlatformRelay/provider-gridscale/issues).

## License

`provider-gridscale` is licensed under the [Apache License 2.0](LICENSE).

This provider is generated with [Upjet](https://github.com/crossplane/upjet) and
wraps the [gridscale Terraform provider](https://github.com/gridscale/terraform-provider-gridscale),
which is distributed under its own license.

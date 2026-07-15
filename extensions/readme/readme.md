# provider-gridscale

`provider-gridscale` is a [Crossplane](https://crossplane.io/) provider for
[gridscale](https://gridscale.io/), built with
[Upjet](https://github.com/crossplane/upjet) (v2) on top of the
[gridscale Terraform provider](https://registry.terraform.io/providers/gridscale/gridscale/latest).

It exposes gridscale cloud resources as Kubernetes Custom Resources so you can
provision and manage gridscale infrastructure declaratively from a Crossplane
control plane.

> **Status:** pre-1.0 (v0.1.x). APIs are generated and may change between
> releases; not yet recommended for production-critical workloads.

## Resources covered

The provider surfaces gridscale resource groups including:

- **Compute:** servers, storages, storage snapshots and snapshot schedules,
  ISO images.
- **Networking:** networks, IP addresses, firewalls, load balancers.
- **PaaS / managed databases:** PostgreSQL, MariaDB/MySQL, and other managed
  database offerings.
- **Object storage & SSL:** object-storage access keys and SSL certificates.

The full, authoritative list of supported kinds is the set of generated
CustomResourceDefinitions in [`package/crds/`](../../package/crds/).

## Install

Apply a `Provider` package manifest to your Crossplane control plane. Replace
`<PACKAGE>` and `<VERSION>` with the published package path and release tag:

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-gridscale
spec:
  package: <PACKAGE>:<VERSION>
```

```console
kubectl apply -f provider.yaml
# or, using the Crossplane CLI:
crossplane install provider <PACKAGE>:<VERSION>
```

Wait for the provider to become healthy:

```console
kubectl get providers
```

## Credentials & ProviderConfig

The provider ships **both** cluster-scoped and namespaced configuration APIs:

| Scope          | API group                      |
| -------------- | ------------------------------ |
| Cluster-scoped | `gridscale.platformrelay.io`   |
| Namespaced     | `gridscale.m.platformrelay.io` |

Credentials are supplied through a Kubernetes `Secret` referenced by a
`ProviderConfig`. The secret value is a **JSON document** with the keys:

| Key       | Required | Description                                                        |
| --------- | -------- | ------------------------------------------------------------------ |
| `uuid`    | yes      | The gridscale User-UUID (from the gridscale panel).                |
| `token`   | yes      | A gridscale API token with full-access permissions.                |
| `api_url` | no       | API base URL. Defaults to `https://api.gridscale.io` when omitted. |

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

### 2. Create a ProviderConfig

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

Managed resources reference this `ProviderConfig` by name to obtain their
credentials. For namespaced usage, the provider also serves `ProviderConfig`
and `ClusterProviderConfig` kinds under the `gridscale.m.platformrelay.io`
group.

## Links

- **Source & issues:** [github.com/PlatformRelay/provider-gridscale](https://github.com/PlatformRelay/provider-gridscale)
- **gridscale:** [gridscale.io](https://gridscale.io/) ·
  [gridscale panel](https://my.gridscale.io/)
- **Crossplane:** [docs.crossplane.io](https://docs.crossplane.io/)

## License

Licensed under the [Apache License 2.0](https://github.com/PlatformRelay/provider-gridscale/blob/main/LICENSE).
This provider is generated with Upjet and wraps the gridscale Terraform
provider, which is distributed under its own license.

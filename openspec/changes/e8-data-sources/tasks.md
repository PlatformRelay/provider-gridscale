# Tasks — e8-data-sources (provider-gridscale)

Backlog stories: **E8-S01…E8-S04** (`agent-context/BACKLOG.md`).
Claim one story per worktree. S01 is independently mergeable — land it first.

## Story/REQ/file-lock table

| Story | REQs | File locks | Meta-test locks | Status |
|-------|------|------------|-----------------|--------|
| E8-S01 | DS-01, DS-02, DS-03 | `examples-generated/cluster/observe/`, `examples-generated/namespaced/observe/`, `docs/observe-only.md` | `hack/test/e8_s01_*` | ⬜ |
| E8-S02 | DS-04, DS-05, DS-07, DS-08, DS-09 | `internal/clients/gridscale_http_test.go`, `internal/clients/gridscale_http.go`, `apis/cluster/storage/v1alpha1/backuplist_*.go`, `apis/namespaced/storage/v1alpha1/backuplist_*.go`, `internal/controller/cluster/storage/backuplist/`, `internal/controller/namespaced/storage/backuplist/`, `package/crds/*backuplist*.yaml` | `hack/test/e8_s02_*` | ⬜ |
| E8-S03 | DS-06, DS-10, DS-11, DS-12, DS-13 | `apis/cluster/gridscale/v1alpha1/publicnetwork_*.go`, `apis/namespaced/gridscale/v1alpha1/publicnetwork_*.go`, `internal/controller/cluster/gridscale/publicnetwork/`, `internal/controller/namespaced/gridscale/publicnetwork/`, `package/crds/*publicnetwork*.yaml` | `hack/test/e8_s03_*` | ⬜ |
| E8-S04 | DS-14, DS-15, DS-16, DS-17, DS-18 | `internal/controller/cluster/setup_custom.go`, `internal/controller/namespaced/setup_custom.go`, `cmd/provider/main.go`, `README.md` | `hack/test/e8_s04_*` | ⬜ |

## Parallelism

```
E8-S01 (independent — no Go, no CRDs, immediately mergeable)
E8-S02 (independent of S01; owns gridscale_http.go — S03 depends on S02)
E8-S03 (depends on S02 for GridscaleClient)
E8-S04 (depends on S02+S03)
```

S01 can land on `main` before S02/S03/S04 start. S02 and S03 **share**
`internal/clients/gridscale_http.go` — the shared client belongs to S02's file lock;
S03 reads it as a dependency, does not modify it.

## Task sequence per story

### E8-S01 — Observe-only YAML examples

1. Read `names.kind` and `names.group` from each of the 19 CRDs in `package/crds/` — do not
   hand-type Kind values; Upjet uses flattened casing (`Isoimage`, `K8S`, `Paas`, `Snapshotschedule`,
   `StorageAccesskey`, `Securityzone`, etc.).
2. For each of the 19 resources, write a cluster-scoped YAML to
   `examples-generated/cluster/observe/<group>/<kind-lowercase>.yaml` with:
   - `managementPolicies: ["Observe"]`
   - `crossplane.io/external-name: <placeholder-uuid>` annotation
   - `meta.upbound.io/example-id: <group>/v1alpha1/<kind-lowercase>-observe`
3. Write the matching namespace-scoped YAML to
   `examples-generated/namespaced/observe/<group>/<kind-lowercase>.yaml`.
4. Write `docs/observe-only.md` explaining the pattern and the UUID requirement.
5. Verify: `find examples-generated/cluster/observe -name '*.yaml' | wc -l` → 19;
   `find examples-generated/namespaced/observe -name '*.yaml' | wc -l` → 19.
6. Gate: `make reviewable` green (no generated-file drift, no lint failures).

### E8-S02 — BackupList CRD + controller

1. **Write failing test first:** create `internal/clients/gridscale_http_test.go` asserting:
   - Correct `X-Auth-UserID` and `X-Auth-Token` headers on `GET` request.
   - JSON decode into `out` on HTTP 200.
   - Non-nil `error` wrapping the status code on HTTP 4xx/5xx.
   Run `go test ./internal/clients/... -run TestGridscaleClient` — must fail (file doesn't exist).
2. Implement `internal/clients/gridscale_http.go` — `GridscaleClient` struct + `Get()` method.
3. Run `go test ./internal/clients/...` — must pass.
4. Write `apis/cluster/storage/v1alpha1/backuplist_types.go` with controller-gen markers
   (`+kubebuilder:object:root=true`, `+kubebuilder:subresource:status`,
   `+kubebuilder:resource:scope=Cluster`). Embed `v1.ResourceSpec` from
   `github.com/crossplane/crossplane-runtime/v2/apis/common/v1`.
5. Write `apis/cluster/storage/v1alpha1/backuplist_register.go` registering into the package's
   existing `SchemeBuilder`.
6. Mirror steps 4–5 for namespaced variant under `apis/namespaced/storage/v1alpha1/`.
7. Run `controller-gen object:headerFile=hack/boilerplate.go.txt paths="./apis/cluster/storage/..."` 
   to regenerate `apis/cluster/storage/v1alpha1/zz_generated.deepcopy.go`; repeat for namespaced.
8. Run `controller-gen crd paths="./apis/..."` to emit
   `package/crds/storage.gridscale.platformrelay.io_backuplists.yaml` and
   `package/crds/storage.gridscale.m.platformrelay.io_backuplists.yaml`.
9. Write `internal/controller/cluster/storage/backuplist/backuplist.go` — `Setup()`,
   `ExternalConnecter.Connect()` (calls `o.SetupFn` → reads `setup.Configuration["uuid"/"token"/"api_url"]`),
   `ExternalClient.Observe()` (calls `GridscaleClient.Get(ctx, "objects/storages/{uuid}/backups", &resp)`
   and populates `status.atProvider.storageBackups`),
   `Create/Update/Delete()` return observe-only error.
10. Mirror step 9 for namespaced variant.
11. Gate: `go test ./internal/clients/...` green; `go build ./...` passes; `make reviewable` clean.

### E8-S03 — PublicNetwork CRD + controller

**Depends on E8-S02 (GridscaleClient is available).**

1. **Write failing test first:** add test to `internal/controller/cluster/gridscale/publicnetwork/`
   (or a new `publicnetwork_test.go`) asserting that `Observe` with a mocked response containing
   `network_type: "public"` sets `ResourceExists: true`; and that an all-private response sets
   `ResourceExists: false`.
2. Write `apis/cluster/gridscale/v1alpha1/publicnetwork_types.go` with controller-gen markers.
   Embed `v1.ResourceSpec`. `Status.AtProvider` contains: `name`, `status`, `networkType`,
   `locationUUID`, `locationName`, `locationCountry`, `locationIATA`, `l2Security bool`,
   `deleteBlock bool`, `labels []string`, `createTime`, `changeTime`.
3. Write `apis/cluster/gridscale/v1alpha1/publicnetwork_register.go`.
4. Mirror for namespaced variant.
5. Run controller-gen to regenerate deepcopy for gridscale package and emit CRD YAML:
   `gridscale.gridscale.platformrelay.io_publicnetworks.yaml` and
   `gridscale.gridscale.m.platformrelay.io_publicnetworks.yaml` in `package/crds/`.
6. Implement `internal/controller/cluster/gridscale/publicnetwork/publicnetwork.go` —
   `Observe()` calls `GET /objects/networks`, filters for `network_type == "public"`,
   sets `ResourceExists: false` when no match found, `ResourceExists: true` + populates
   `status.atProvider` when matched.
7. Mirror step 6 for namespaced variant.
8. Gate: failing test from step 1 now green; `go build ./...` passes; `make reviewable` clean.

### E8-S04 — Wire everything + README

**Depends on E8-S02 and E8-S03.**

1. Write `internal/controller/cluster/setup_custom.go` — `SetupCustom(mgr ctrl.Manager, o controller.Options) error`
   calling `backuplist.Setup` and `publicnetwork.Setup`. No `zz_` prefix; not generated.
2. Write `internal/controller/namespaced/setup_custom.go` — mirrors cluster setup.
3. Edit `cmd/provider/main.go` to call `controllerCluster.SetupCustom` and
   `controllerNamespaced.SetupCustom` alongside the existing `Setup` calls. Use the same
   `customresourcesgate` / `SetupCustomGated` guard pattern already present in main.go so the
   provider starts cleanly if a CRD is not yet installed.
4. Add BackupList and PublicNetwork rows to the README resource matrix.
5. Gate: `go build ./...` passes; `go vet ./...` passes;
   `grep -r "BackupList\|PublicNetwork" internal/controller/cluster/setup_custom.go` non-empty;
   `grep -r "BackupList\|PublicNetwork" internal/controller/namespaced/setup_custom.go` non-empty;
   `make reviewable` clean.

## Gate commands (all stories)

```bash
go test ./internal/clients/...           # DS-04 (S02)
go build ./...                           # DS-14 (S04)
go vet ./...                             # DS-15 (S04)
make reviewable                          # generation boundary guard (all stories)
make test                                # unit + contract suite (all stories)
```

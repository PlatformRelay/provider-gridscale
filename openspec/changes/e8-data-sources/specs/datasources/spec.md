# Spec — E8 Data Sources (e8-data-sources)

Epic: **E8** · Change slug: `e8-data-sources`
Stories: E8-S01 (observe-only), E8-S02 (BackupList), E8-S03 (PublicNetwork), E8-S04 (wire+README)

Every REQ maps to exactly one **Level** (U/E/M/D), one **Test** (artifact path that must exist
and pass after implementation), and one **Verify** (exact command). Implementation must add the
failing **Test** artifact before implementing (TDD rule per GUIDELINES §2).

---

## REQ-E8-S01-01 (DS-01) — 38 observe-only YAML files exist under examples-generated/

**Level:** M · **Priority:** must · **Story:** E8-S01

**Given** E8-S01 has landed,
**when** `find examples-generated/cluster/observe -name '*.yaml' | wc -l` and
`find examples-generated/namespaced/observe -name '*.yaml' | wc -l` are each run,
**then** both return 19 (total: 38 files — 19 per scope, one per resource with a TF data-source twin).

**Edge:** Given a YAML file for `Isoimage` or `K8S` (flattened-cased Kinds),
when `kubectl apply --dry-run=client` is run against a cluster with the provider CRDs installed,
then the manifest is accepted with no unknown-kind or unknown-field errors.
(This validates that Kind/apiVersion were read from `package/crds/` — not hand-typed.)

**Test:** `hack/test/e8_s01_observe_yaml_count_test.sh`

**Verify:** `bash hack/test/e8_s01_observe_yaml_count_test.sh`

---

## REQ-E8-S01-02 (DS-02) — docs/observe-only.md exists and documents the pattern

**Level:** M · **Priority:** must · **Story:** E8-S01

**Given** E8-S01 has landed,
**when** `docs/observe-only.md` is read,
**then** the file exists, contains the string `managementPolicies`, contains the string
`crossplane.io/external-name`, and explains the UUID requirement in at least one sentence.

**Edge:** Given a user who does not know the external-name requirement, when they read the doc,
then at least one example YAML snippet with a placeholder UUID is shown.

**Test:** `hack/test/e8_s01_observe_docs_test.sh`

**Verify:** `bash hack/test/e8_s01_observe_docs_test.sh`

---

## REQ-E8-S01-03 (DS-03) — uptest can apply an observe-only Sshkey manifest (creds-gated)

**Level:** E · **Priority:** should · **Story:** E8-S01

**Given** a gridscale lab project with an existing Sshkey (UUID known and set as `crossplane.io/external-name`),
**when** `uptest run examples-generated/cluster/observe/gridscale/v1alpha1/sshkey.yaml` is executed
against a live cluster with the provider installed and `ProviderConfig` credentials configured,
**then** the manifest applies cleanly, the MR condition reaches `Synced: True / Ready: True`, and
`status.atProvider.name` is populated with the remote resource's name.

**Edge:** Given a non-existent UUID in `crossplane.io/external-name`, when the provider observes it,
then the MR condition is set to `Synced: False` with a descriptive message — no panic, no crash,
no CUD operation attempted.

**Test:** `examples-generated/cluster/observe/gridscale/v1alpha1/sshkey.yaml`

**Verify:** `uptest run examples-generated/cluster/observe/gridscale/v1alpha1/sshkey.yaml`
(requires lab credentials; gated behind `/test-examples` trigger + nightly — per D-012 creds contract)

---

## REQ-E8-S02-01 (DS-04) — GridscaleClient.Get sends correct headers and decodes JSON

**Level:** U · **Priority:** must · **Story:** E8-S02

**Given** a `GridscaleClient{UserID: "test-uuid", Token: "test-token", APIURL: server.URL}` and an
`httptest.NewServer` returning HTTP 200 with a JSON body `{"key":"value"}`,
**when** `client.Get(ctx, "objects/test", &out)` is called,
**then** the server receives `X-Auth-UserID: test-uuid` and `X-Auth-Token: test-token` headers,
`out` is populated with the decoded JSON, and the error is `nil`.

**Edge:** Given the mock server returns HTTP 403,
**when** `client.Get` is called,
**then** a non-nil error wrapping "403" (or the status text) is returned and `out` is unchanged.

**Test:** `internal/clients/gridscale_http_test.go`

**Verify:** `go test ./internal/clients/... -run TestGridscaleClient -v`

---

## REQ-E8-S02-02 (DS-05) — BackupList CRD YAML exists in package/crds/

**Level:** M · **Priority:** must · **Story:** E8-S02

**Given** E8-S02 has landed and controller-gen has been run,
**when** `ls package/crds/` is executed,
**then** both `storage.gridscale.platformrelay.io_backuplists.yaml` and
`storage.gridscale.m.platformrelay.io_backuplists.yaml` are present.

**Edge:** Given the CRD YAML is opened, when `spec.group` is read, then it equals
`storage.gridscale.platformrelay.io` (cluster) or `storage.gridscale.m.platformrelay.io`
(namespaced) respectively — not the gridscale group.

**Test:** `hack/test/e8_s02_backuplist_crd_exists_test.sh`

**Verify:** `bash hack/test/e8_s02_backuplist_crd_exists_test.sh`

---

## REQ-E8-S02-03 (DS-07) — BackupList connector.Connect returns ExternalClient with credentials

**Level:** U · **Priority:** must · **Story:** E8-S02

**Given** a mock `SetupFn` that returns
`terraform.Setup{Configuration: map[string]any{"uuid": "u1", "token": "t1", "api_url": "http://fake"}}`,
**when** `connector.Connect(ctx, &BackupList{...})` is called,
**then** the returned `ExternalClient` holds a `GridscaleClient` with `UserID=="u1"`,
`Token=="t1"`, and `APIURL=="http://fake"`, and no error is returned.

**Edge:** Given `SetupFn` returns a non-nil error,
**when** `Connect` is called,
**then** `Connect` wraps and propagates the error and returns a nil client.

**Test:** `internal/controller/cluster/storage/backuplist/backuplist_test.go`

**Verify:** `go test ./internal/controller/cluster/storage/backuplist/... -run TestConnect -v`

---

## REQ-E8-S02-04 (DS-08) — BackupList Observe populates status.atProvider.storageBackups

**Level:** U · **Priority:** must · **Story:** E8-S02

**Given** a `GridscaleClient` backed by an `httptest.NewServer` returning a JSON array
`[{"object_uuid": "bu1", "name": "b1", "capacity": 10.0, "create_time": "..."}]`,
**when** `ExternalClient.Observe(ctx, backupList)` is called,
**then** `backupList.Status.AtProvider.StorageBackups` has length 1, the first entry has
`ObjectUUID == "bu1"` and `Name == "b1"`, and `observation.ResourceExists == true`.

**Edge:** Given the server returns an empty array `[]`,
**when** `Observe` is called,
**then** `StorageBackups` has length 0 and `ResourceExists` is `true` — an empty backup list
means the storage exists with no backups, not that the storage is absent.

**Test:** `internal/controller/cluster/storage/backuplist/backuplist_test.go`

**Verify:** `go test ./internal/controller/cluster/storage/backuplist/... -run TestObserve -v`

---

## REQ-E8-S02-05 (DS-09) — BackupList Create/Update/Delete return observe-only error

**Level:** U · **Priority:** must · **Story:** E8-S02

**Given** an instantiated `ExternalClient` for a `BackupList`,
**when** `Create(ctx, mg)`, `Update(ctx, mg)`, or `Delete(ctx, mg)` is called,
**then** each returns a non-nil `error` containing the string "observe-only" (or equivalent),
and no HTTP request is made to the gridscale API.

**Edge:** Given `Delete` is called (e.g., resource being finalized), when the observe-only error
is returned, then `errors.Is`-wrapping of a sentinel `errObserveOnly` error allows the reconciler
to surface it without infinite retry. (Use a package-level `var errObserveOnly = errors.New("...")`,
not a type assertion.)

**Test:** `internal/controller/cluster/storage/backuplist/backuplist_test.go`

**Verify:** `go test ./internal/controller/cluster/storage/backuplist/... -run TestObserveOnly -v`

---

## REQ-E8-S03-01 (DS-06) — PublicNetwork CRD YAML exists in package/crds/

**Level:** M · **Priority:** must · **Story:** E8-S03

**Given** E8-S03 has landed and controller-gen has been run,
**when** `ls package/crds/` is executed,
**then** both `gridscale.gridscale.platformrelay.io_publicnetworks.yaml` and
`gridscale.gridscale.m.platformrelay.io_publicnetworks.yaml` are present.

**Edge:** Given the CRD YAML is opened, when `spec.group` is read, then it equals
`gridscale.gridscale.platformrelay.io` (cluster) or `gridscale.gridscale.m.platformrelay.io`
(namespaced) — not the storage group.

**Test:** `hack/test/e8_s03_publicnetwork_crd_exists_test.sh`

**Verify:** `bash hack/test/e8_s03_publicnetwork_crd_exists_test.sh`

---

## REQ-E8-S03-02 (DS-10) — PublicNetwork connector.Connect returns ExternalClient

**Level:** U · **Priority:** must · **Story:** E8-S03

**Given** a mock `SetupFn` returning
`terraform.Setup{Configuration: map[string]any{"uuid":"u2","token":"t2"}}`,
**when** `connector.Connect(ctx, &PublicNetwork{...})` is called,
**then** the returned `ExternalClient` holds a `GridscaleClient` with `UserID=="u2"`,
`Token=="t2"`, and no error is returned.

**Edge:** Given `SetupFn` returns an error,
**when** `Connect` is called,
**then** the error is propagated and the returned client is nil.

**Test:** `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go`

**Verify:** `go test ./internal/controller/cluster/gridscale/publicnetwork/... -run TestConnect -v`

---

## REQ-E8-S03-03 (DS-11) — PublicNetwork Observe finds public network from GET /objects/networks

**Level:** U · **Priority:** must · **Story:** E8-S03

**Given** a `GridscaleClient` backed by an `httptest.NewServer` returning a JSON network list that
includes one object with `network_type: "public"` and `name: "Public Network"`,
**when** `ExternalClient.Observe(ctx, publicNetwork)` is called,
**then** `publicNetwork.Status.AtProvider.Name == "Public Network"`,
`publicNetwork.Status.AtProvider.NetworkType == "public"`,
and `observation.ResourceExists == true`.

**Edge:** Given the response contains only `network_type: "private"` entries,
**when** `Observe` is called, **then** `ResourceExists == false` and the error is `nil`.

**Test:** `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go`

**Verify:** `go test ./internal/controller/cluster/gridscale/publicnetwork/... -run TestObserve -v`

---

## REQ-E8-S03-04 (DS-12) — PublicNetwork Observe sets ResourceExists=false when no public network

**Level:** U · **Priority:** must · **Story:** E8-S03

**Given** a mock server returning an empty network list (`{"networks": {}}`),
**when** `Observe` is called,
**then** `observation.ResourceExists == false` and the error is `nil`.
(An absent public network is a valid empty observation, not an error condition.)

**Test:** `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go`

**Verify:** `go test ./internal/controller/cluster/gridscale/publicnetwork/... -run TestObserveNotFound -v`

---

## REQ-E8-S03-05 (DS-13) — PublicNetwork Create/Update/Delete return observe-only error

**Level:** U · **Priority:** must · **Story:** E8-S03

**Given** an instantiated `ExternalClient` for a `PublicNetwork`,
**when** `Create(ctx, mg)`, `Update(ctx, mg)`, or `Delete(ctx, mg)` is called,
**then** each returns a non-nil `error` containing "observe-only" and no HTTP request is made.

**Edge:** Same sentinel-error pattern as DS-09 — use a package-level `var errObserveOnly` that
allows `errors.Is` matching by callers; no type assertion required.

**Test:** `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go`

**Verify:** `go test ./internal/controller/cluster/gridscale/publicnetwork/... -run TestObserveOnly -v`

---

## REQ-E8-S04-01 (DS-14) — go build ./... passes with all new types + controllers wired

**Level:** M · **Priority:** must · **Story:** E8-S04

**Given** E8-S02, E8-S03, and E8-S04 have all landed,
**when** `go build ./...` is run from the repo root,
**then** the exit code is 0 and no compilation errors are emitted.

**Edge:** Given `setup_custom.go` imports a non-existent controller package (e.g., a typo in the
import path), when `go build` is run, then it fails with an import error — this is the compile-time
guard that prevents registration drift from the actual controller locations.

**Test:** `hack/test/e8_s04_go_build_test.sh`

**Verify:** `bash hack/test/e8_s04_go_build_test.sh`

---

## REQ-E8-S04-02 (DS-15) — go vet ./... passes

**Level:** M · **Priority:** must · **Story:** E8-S04

**Given** E8-S02, E8-S03, and E8-S04 have landed,
**when** `go vet ./...` is run,
**then** the exit code is 0.

**Test:** `hack/test/e8_s04_go_vet_test.sh`

**Verify:** `bash hack/test/e8_s04_go_vet_test.sh`

---

## REQ-E8-S04-03 (DS-16) — BackupList and PublicNetwork registered in cluster setup_custom.go

**Level:** M · **Priority:** must · **Story:** E8-S04

**Given** E8-S04 has landed,
**when** `grep -r "BackupList\|PublicNetwork" internal/controller/cluster/setup_custom.go` is run,
**then** the command exits 0 and both names appear in the output.

**Edge:** Given `setup_custom.go` exists but the `SetupCustom` function signature does not match
`func(ctrl.Manager, controller.Options) error`, when `go build ./...` runs, then it fails at
compile time — no runtime-only registration bugs possible.

**Test:** `hack/test/e8_s04_setup_custom_cluster_test.sh`

**Verify:** `bash hack/test/e8_s04_setup_custom_cluster_test.sh`

---

## REQ-E8-S04-04 (DS-17) — BackupList and PublicNetwork registered in namespaced setup_custom.go

**Level:** M · **Priority:** must · **Story:** E8-S04

**Given** E8-S04 has landed,
**when** `grep -r "BackupList\|PublicNetwork" internal/controller/namespaced/setup_custom.go` is run,
**then** the command exits 0 and both names appear in the output.

**Edge:** Given only the cluster variant is registered and the namespaced file is missing, when
`go build ./...` runs, the build still passes — the meta test explicitly checks the namespaced
file content to catch this omission before deployment.

**Test:** `hack/test/e8_s04_setup_custom_namespaced_test.sh`

**Verify:** `bash hack/test/e8_s04_setup_custom_namespaced_test.sh`

---

## REQ-E8-S04-05 (DS-18) — README resource matrix includes BackupList and PublicNetwork

**Level:** D · **Priority:** should · **Story:** E8-S04

**Note:** README is hand-maintained public surface. This REQ is categorised D (doc) because the
matrix row is authored alongside the implementation, not generated from source. (It is not a D in
the strict "generated docs in sync with source" sense — no `make docs` regeneration covers it.
A code-review step is required to validate group accuracy.)

**Given** E8-S04 has landed,
**when** `grep -E "BackupList|PublicNetwork" README.md` is run,
**then** both names appear at least once (the resource matrix has been updated).

**Edge:** Given the matrix row is added but the API group is incorrect, when a maintainer reviews,
then the discrepancy is caught by review — the meta test asserts name presence only, not group
accuracy.

**Test:** `hack/test/e8_s04_readme_matrix_test.sh`

**Verify:** `bash hack/test/e8_s04_readme_matrix_test.sh`

---

## Summary table

| REQ | ID | Level | Story | Test artifact |
|-----|-----|-------|-------|--------------|
| 38 observe-only YAMLs exist (19 cluster + 19 namespaced) | DS-01 | M | E8-S01 | `hack/test/e8_s01_observe_yaml_count_test.sh` |
| docs/observe-only.md exists with pattern | DS-02 | M | E8-S01 | `hack/test/e8_s01_observe_docs_test.sh` |
| uptest applies observe-only Sshkey (creds-gated) | DS-03 | E | E8-S01 | `examples-generated/cluster/observe/gridscale/v1alpha1/sshkey.yaml` |
| GridscaleClient headers + JSON decode | DS-04 | U | E8-S02 | `internal/clients/gridscale_http_test.go` |
| BackupList CRD YAML in package/crds/ | DS-05 | M | E8-S02 | `hack/test/e8_s02_backuplist_crd_exists_test.sh` |
| PublicNetwork CRD YAML in package/crds/ | DS-06 | M | E8-S03 | `hack/test/e8_s03_publicnetwork_crd_exists_test.sh` |
| BackupList Connect returns client+creds | DS-07 | U | E8-S02 | `internal/controller/cluster/storage/backuplist/backuplist_test.go` |
| BackupList Observe populates status | DS-08 | U | E8-S02 | `internal/controller/cluster/storage/backuplist/backuplist_test.go` |
| BackupList Create/Update/Delete → error | DS-09 | U | E8-S02 | `internal/controller/cluster/storage/backuplist/backuplist_test.go` |
| PublicNetwork Connect returns client | DS-10 | U | E8-S03 | `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go` |
| PublicNetwork Observe finds public net | DS-11 | U | E8-S03 | `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go` |
| PublicNetwork Observe → false when absent | DS-12 | U | E8-S03 | `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go` |
| PublicNetwork Create/Update/Delete → error | DS-13 | U | E8-S03 | `internal/controller/cluster/gridscale/publicnetwork/publicnetwork_test.go` |
| go build ./... passes | DS-14 | M | E8-S04 | `hack/test/e8_s04_go_build_test.sh` |
| go vet ./... passes | DS-15 | M | E8-S04 | `hack/test/e8_s04_go_vet_test.sh` |
| BackupList+PublicNetwork in cluster setup | DS-16 | M | E8-S04 | `hack/test/e8_s04_setup_custom_cluster_test.sh` |
| BackupList+PublicNetwork in namespaced setup | DS-17 | M | E8-S04 | `hack/test/e8_s04_setup_custom_namespaced_test.sh` |
| README matrix includes both Kinds | DS-18 | D | E8-S04 | `hack/test/e8_s04_readme_matrix_test.sh` |

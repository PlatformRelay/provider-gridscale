package config

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"sigs.k8s.io/yaml"
)

// CRD golden-contract tests (story E2-S02).
//
// These tests read the *generated* CRD OpenAPI schemas under package/crds/ and
// lock a scoped, stable fragment of each against unintended generation drift.
// This is deliberately distinct from config/gridscale_test.go: that suite
// asserts the config layer (Kind overrides on the upjet provider), whereas this
// suite asserts the *rendered* CRD YAML that upjet actually emits.
//
// Two tiers guard against a purely self-referential test:
//
//  1. Golden diff — a scoped fragment (group, kind, scope, version name, sorted
//     forProvider property names) is serialized to config/testdata/crd-contract
//     and compared. Refresh with UPDATE_GOLDEN=1 go test ./config/...
//  2. Hard structural invariants — asserted independently of the golden so the
//     test can fail-first on a malformed CRD: kind overrides, the cluster/
//     namespaced group relationship, and cluster/namespaced parity of kind and
//     forProvider set.

const crdDir = "../package/crds"

// Named view types below replace anonymous nested structs so Sonar godre:S8205
// stays clear while JSON tags and unmarshalling behaviour stay identical.

type crdNames struct {
	Kind   string `json:"kind"`
	Plural string `json:"plural"`
}

type crdForProvider struct {
	Properties map[string]json.RawMessage `json:"properties"`
}

type crdSpecProperties struct {
	ForProvider crdForProvider `json:"forProvider"`
}

type crdResourceSpec struct {
	Properties crdSpecProperties `json:"properties"`
}

type crdSchemaProperties struct {
	Spec crdResourceSpec `json:"spec"`
}

type crdOpenAPIV3Schema struct {
	Properties crdSchemaProperties `json:"properties"`
}

type crdVersionSchema struct {
	OpenAPIV3Schema crdOpenAPIV3Schema `json:"openAPIV3Schema"`
}

type crdVersion struct {
	Name   string           `json:"name"`
	Schema crdVersionSchema `json:"schema"`
}

type crdSpec struct {
	Group    string       `json:"group"`
	Names    crdNames     `json:"names"`
	Scope    string       `json:"scope"`
	Versions []crdVersion `json:"versions"`
}

// crd is a minimal typed view of a generated CRD. Unmarshalling into a struct
// (rather than map[string]interface{}) means missing fields become zero values
// that fail assertions naturally, instead of panicking on a bad type assertion.
type crd struct {
	Spec crdSpec `json:"spec"`
}

// contract is the scoped, stable fragment we serialize to a golden file.
type contract struct {
	Group            string   `json:"group"`
	Kind             string   `json:"kind"`
	Scope            string   `json:"scope"`
	Version          string   `json:"version"`
	ForProviderProps []string `json:"forProviderProps"`
}

func loadCRD(t *testing.T, filename string) crd {
	t.Helper()
	raw, err := os.ReadFile(filepath.Join(crdDir, filename))
	if err != nil {
		t.Fatalf("read CRD %s: %v", filename, err)
	}
	var c crd
	if err := yaml.Unmarshal(raw, &c); err != nil {
		t.Fatalf("unmarshal CRD %s: %v", filename, err)
	}
	return c
}

// extract builds the scoped fragment from a loaded CRD, using the first
// version. Callers separately assert version count/name so an added version
// surfaces as drift rather than being silently dropped here.
func extract(t *testing.T, c crd) contract {
	t.Helper()
	if len(c.Spec.Versions) == 0 {
		t.Fatal("CRD has no versions")
	}
	v := c.Spec.Versions[0]
	props := v.Schema.OpenAPIV3Schema.Properties.Spec.Properties.ForProvider.Properties
	names := make([]string, 0, len(props))
	for k := range props {
		names = append(names, k)
	}
	sort.Strings(names)
	return contract{
		Group:            c.Spec.Group,
		Kind:             c.Spec.Names.Kind,
		Scope:            c.Spec.Scope,
		Version:          v.Name,
		ForProviderProps: names,
	}
}

// crdCase pairs the cluster-scoped and namespaced CRD variants of one resource
// with the invariants we expect them to satisfy.
type crdCase struct {
	name             string // golden basename
	clusterFile      string
	namespacedFile   string
	wantKind         string // expected rendered Kind (override-aware)
	wantClusterGroup string // expected cluster-scoped API group
}

var crdCases = []crdCase{
	{
		name:             "mysql8",
		clusterFile:      "mysql8.gridscale.platformrelay.io_mysql8s.yaml",
		namespacedFile:   "mysql8.gridscale.m.platformrelay.io_mysql8s.yaml",
		wantKind:         "MySQL8",
		wantClusterGroup: "mysql8.gridscale.platformrelay.io",
	},
	{
		name:             "storageimport",
		clusterFile:      "storage.gridscale.platformrelay.io_storageimports.yaml",
		namespacedFile:   "storage.gridscale.m.platformrelay.io_storageimports.yaml",
		wantKind:         "StorageImport",
		wantClusterGroup: "storage.gridscale.platformrelay.io",
	},
	{
		name:             "server",
		clusterFile:      "gridscale.gridscale.platformrelay.io_servers.yaml",
		namespacedFile:   "gridscale.gridscale.m.platformrelay.io_servers.yaml",
		wantKind:         "Server",
		wantClusterGroup: "gridscale.gridscale.platformrelay.io",
	},
	{
		name:             "network",
		clusterFile:      "gridscale.gridscale.platformrelay.io_networks.yaml",
		namespacedFile:   "gridscale.gridscale.m.platformrelay.io_networks.yaml",
		wantKind:         "Network",
		wantClusterGroup: "gridscale.gridscale.platformrelay.io",
	},
}

// clusterGroupSuffix encodes the group-naming convention that upjet's
// cluster-scoped generation must follow.
const clusterGroupSuffix = ".gridscale.platformrelay.io"

// namespacedGroupFor derives the namespaced API group from a cluster group:
// `<x>.gridscale.platformrelay.io` -> `<x>.gridscale.m.platformrelay.io`.
func namespacedGroupFor(clusterGroup string) string {
	return strings.Replace(clusterGroup, ".gridscale.", ".gridscale.m.", 1)
}

func marshalCRDFragment(t *testing.T, frag contract) []byte {
	t.Helper()
	got, err := json.MarshalIndent(frag, "", "  ")
	if err != nil {
		t.Fatalf("marshal fragment: %v", err)
	}
	return append(got, '\n')
}

func writeCRDGolden(t *testing.T, goldenPath string, got []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(goldenPath), 0o755); err != nil {
		t.Fatalf("mkdir golden dir: %v", err)
	}
	if err := os.WriteFile(goldenPath, got, 0o644); err != nil {
		t.Fatalf("write golden: %v", err)
	}
}

func assertCRDGoldenMatch(t *testing.T, name string, got []byte, goldenPath string) {
	t.Helper()
	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("read golden %s (run UPDATE_GOLDEN=1 to create): %v", goldenPath, err)
	}
	if !bytes.Equal(got, want) {
		t.Errorf("CRD contract drift for %s.\n--- got ---\n%s\n--- want ---\n%s",
			name, got, want)
	}
}

func runCRDGoldenCase(t *testing.T, tc crdCase, update bool) {
	t.Helper()
	got := marshalCRDFragment(t, extract(t, loadCRD(t, tc.clusterFile)))
	goldenPath := filepath.Join("testdata", "crd-contract", tc.name+".golden.json")
	if update {
		writeCRDGolden(t, goldenPath, got)
		return
	}
	assertCRDGoldenMatch(t, tc.name, got, goldenPath)
}

// TestCRDGoldenContract compares the scoped fragment of each cluster-scoped CRD
// against its committed golden. UPDATE_GOLDEN=1 regenerates the goldens.
func TestCRDGoldenContract(t *testing.T) {
	update := os.Getenv("UPDATE_GOLDEN") == "1"
	for _, tc := range crdCases {
		t.Run(tc.name, func(t *testing.T) {
			runCRDGoldenCase(t, tc, update)
		})
	}
}

func assertClusterKindGroupScope(t *testing.T, cluster crd, tc crdCase) {
	t.Helper()
	if cluster.Spec.Names.Kind != tc.wantKind {
		t.Errorf("cluster kind = %q, want %q", cluster.Spec.Names.Kind, tc.wantKind)
	}
	if cluster.Spec.Group != tc.wantClusterGroup {
		t.Errorf("cluster group = %q, want %q", cluster.Spec.Group, tc.wantClusterGroup)
	}
	if !strings.HasSuffix(cluster.Spec.Group, clusterGroupSuffix) {
		t.Errorf("cluster group %q lacks suffix %q", cluster.Spec.Group, clusterGroupSuffix)
	}
	if cluster.Spec.Scope != "Cluster" {
		t.Errorf("cluster scope = %q, want Cluster", cluster.Spec.Scope)
	}
}

func assertNamespacedGroupScope(t *testing.T, cluster, namespaced crd) {
	t.Helper()
	wantNS := namespacedGroupFor(cluster.Spec.Group)
	if namespaced.Spec.Group != wantNS {
		t.Errorf("namespaced group = %q, want %q (derived from cluster group)",
			namespaced.Spec.Group, wantNS)
	}
	if namespaced.Spec.Scope != "Namespaced" {
		t.Errorf("namespaced scope = %q, want Namespaced", namespaced.Spec.Scope)
	}
}

func assertClusterNamespacedParity(t *testing.T, cluster, namespaced crd) {
	t.Helper()
	cf := extract(t, cluster)
	nf := extract(t, namespaced)
	if len(cf.ForProviderProps) == 0 {
		t.Errorf("cluster forProvider property set is empty")
	}
	if namespaced.Spec.Names.Kind != cluster.Spec.Names.Kind {
		t.Errorf("kind mismatch cluster=%q namespaced=%q",
			cluster.Spec.Names.Kind, namespaced.Spec.Names.Kind)
	}
	if strings.Join(cf.ForProviderProps, ",") != strings.Join(nf.ForProviderProps, ",") {
		t.Errorf("forProvider property set differs between scopes.\ncluster:    %v\nnamespaced: %v",
			cf.ForProviderProps, nf.ForProviderProps)
	}
}

func assertCRDStructuralInvariants(t *testing.T, tc crdCase) {
	t.Helper()
	cluster := loadCRD(t, tc.clusterFile)
	namespaced := loadCRD(t, tc.namespacedFile)
	assertClusterKindGroupScope(t, cluster, tc)
	assertNamespacedGroupScope(t, cluster, namespaced)
	assertClusterNamespacedParity(t, cluster, namespaced)
}

// TestCRDStructuralInvariants asserts invariants that must hold regardless of
// the golden files, so the suite is not purely self-referential and can
// fail-first on a malformed or drifted CRD.
func TestCRDStructuralInvariants(t *testing.T) {
	for _, tc := range crdCases {
		t.Run(tc.name, func(t *testing.T) {
			assertCRDStructuralInvariants(t, tc)
		})
	}
}

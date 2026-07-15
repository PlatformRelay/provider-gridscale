package config

import (
	"testing"
)

// TestKindOverrides locks in the two hand-picked Kind overrides that
// configureGridscale applies. Both guard against upjet's automatic group/kind
// derivation producing invalid Go output:
//
//   - gridscale_mysql8_0: the trailing "_0" version segment makes upjet derive
//     the invalid Go kind "0"; it is overridden to "MySQL8".
//   - gridscale_storage_import: kind "Import" -> controller package "import",
//     which is a Go reserved word; it is overridden to "StorageImport" so the
//     package becomes "storageimport".
//
// If either override is dropped, upjet re-derives the broken kind and these
// assertions fail — that is the regression this test exists to catch.
func TestKindOverrides(t *testing.T) {
	pc := GetProvider()
	if pc == nil {
		t.Fatal("GetProvider() returned nil")
	}

	cases := []struct {
		terraformName string
		wantKind      string
	}{
		{terraformName: "gridscale_mysql8_0", wantKind: "MySQL8"},
		{terraformName: "gridscale_storage_import", wantKind: "StorageImport"},
	}

	for _, tc := range cases {
		t.Run(tc.terraformName, func(t *testing.T) {
			r, ok := pc.Resources[tc.terraformName]
			if !ok {
				t.Fatalf("resource %q not present in configured provider", tc.terraformName)
			}
			if r.Kind != tc.wantKind {
				t.Errorf("resource %q Kind: want %q, got %q", tc.terraformName, tc.wantKind, r.Kind)
			}
		})
	}
}

// TestKindOverridesAppliedToNamespacedProvider ensures the same overrides are
// wired into the namespaced provider configuration, since configureGridscale is
// applied by both GetProvider and GetProviderNamespaced.
func TestKindOverridesAppliedToNamespacedProvider(t *testing.T) {
	pc := GetProviderNamespaced()
	if pc == nil {
		t.Fatal("GetProviderNamespaced() returned nil")
	}

	for name, wantKind := range map[string]string{
		"gridscale_mysql8_0":       "MySQL8",
		"gridscale_storage_import": "StorageImport",
	} {
		r, ok := pc.Resources[name]
		if !ok {
			t.Fatalf("resource %q not present in namespaced provider", name)
		}
		if r.Kind != wantKind {
			t.Errorf("namespaced resource %q Kind: want %q, got %q", name, wantKind, r.Kind)
		}
	}
}

// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const validSchema = `{
  "provider_schemas": {
    "registry.terraform.io/gridscale/gridscale": {
      "resource_schemas": {
        "gridscale_gamma": {},
        "gridscale_alpha": {},
        "gridscale_beta": {}
      }
    }
  }
}`

func writeTemp(t *testing.T, name, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), name)
	if err := os.WriteFile(p, []byte(body), 0o600); err != nil {
		t.Fatalf("write %s: %v", p, err)
	}
	return p
}

func TestLoadSchemaResourceKeys_ReturnsSortedKeys(t *testing.T) {
	keys, err := loadSchemaResourceKeys(writeTemp(t, "schema.json", validSchema))
	if err != nil {
		t.Fatalf("loadSchemaResourceKeys: %v", err)
	}
	want := []string{"gridscale_alpha", "gridscale_beta", "gridscale_gamma"}
	if len(keys) != len(want) {
		t.Fatalf("keys = %v, want %v", keys, want)
	}
	for i := range want {
		if keys[i] != want[i] {
			t.Errorf("keys[%d] = %q, want %q (keys must be sorted)", i, keys[i], want[i])
		}
	}
}

func TestLoadSchemaResourceKeys_EmptyWhenNoProviderSchemas(t *testing.T) {
	keys, err := loadSchemaResourceKeys(writeTemp(t, "schema.json", `{"provider_schemas":{}}`))
	if err != nil {
		t.Fatalf("loadSchemaResourceKeys: %v", err)
	}
	if len(keys) != 0 {
		t.Errorf("keys = %v, want empty", keys)
	}
}

func TestLoadSchemaResourceKeys_ErrorsOnMissingFile(t *testing.T) {
	_, err := loadSchemaResourceKeys(filepath.Join(t.TempDir(), "absent.json"))
	if err == nil {
		t.Fatal("loadSchemaResourceKeys must error when the schema file is absent")
	}
}

func TestLoadSchemaResourceKeys_ErrorsOnMalformedJSON(t *testing.T) {
	_, err := loadSchemaResourceKeys(writeTemp(t, "schema.json", "{not json"))
	if err == nil {
		t.Fatal("loadSchemaResourceKeys must error on malformed JSON")
	}
}

func TestExistingResourceKeys_SkipsNonResourceLines(t *testing.T) {
	content := strings.Join([]string{
		"resources:",
		"    gridscale_storage:",
		"        name: gridscale_storage",
		"    :",               // empty key
		"    two words:",      // scraper title key, re-keyed by fixMetadata
		"        nested_key:", // 8-space indent is not a resource key
		"    gridscale_network:",
	}, "\n")

	got := existingResourceKeys(content)
	want := map[string]struct{}{"gridscale_storage": {}, "gridscale_network": {}}
	if len(got) != len(want) {
		t.Fatalf("keys = %v, want %v", got, want)
	}
	for k := range want {
		if _, ok := got[k]; !ok {
			t.Errorf("missing expected resource key %q", k)
		}
	}
}

func TestInjectMissingResourceStubs_AppendsWhenContentLacksTrailingNewline(t *testing.T) {
	// No key sorts after gridscale_zzz, so the stub is appended at EOF; the
	// content deliberately has no trailing newline.
	content := "resources:\n    gridscale_aaa:\n        name: gridscale_aaa"
	got := injectMissingResourceStubs(content, []string{"gridscale_zzz"})

	if !strings.Contains(got, resourceKeyLine("gridscale_zzz")) {
		t.Fatalf("stub for gridscale_zzz not injected:\n%s", got)
	}
	if strings.Contains(got, "gridscale_aaa"+"    gridscale_zzz") {
		t.Error("stub must be separated from preceding content by a newline")
	}
	if !strings.Contains(got, "name: gridscale_aaa\n    gridscale_zzz:") {
		t.Errorf("expected a newline inserted before the appended stub, got:\n%s", got)
	}
}

func TestInjectMissingResourceStubs_AppendsWhenNeighbourLineReformatted(t *testing.T) {
	// gridscale_zzz exists as a *present* key so it is chosen as the neighbour,
	// but its key line is written in a form insertStub cannot locate verbatim,
	// so the stub falls through to a plain append.
	content := "resources:\n    gridscale_zzz:\r\n        name: gridscale_zzz\n"
	got := injectMissingResourceStubs(content, []string{"gridscale_mmm"})

	if !strings.Contains(got, resourceKeyLine("gridscale_mmm")) {
		t.Fatalf("stub for gridscale_mmm not injected:\n%s", got)
	}
	if !strings.HasSuffix(strings.TrimRight(got, "\n"), "importStatements: []") {
		t.Errorf("stub should have been appended at EOF, got:\n%s", got)
	}
}

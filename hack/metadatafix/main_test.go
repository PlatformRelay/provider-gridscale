package main

import (
	"strings"
	"testing"
)

// TestInjectMissingResourceStubs_InsertsAlphabetically asserts a schema
// resource absent from scraped metadata gets a minimal stub entry, inserted
// before the next existing key in alphabetical order.
func TestInjectMissingResourceStubs_InsertsAlphabetically(t *testing.T) {
	in := "name: gridscale/gridscale\nresources:\n" +
		"    gridscale_backupschedule:\n" +
		"        name: backup\n" +
		"    gridscale_firewall:\n" +
		"        name: fw\n"

	got := injectMissingResourceStubs(in, []string{
		"gridscale_backupschedule",
		"gridscale_filesystem",
		"gridscale_firewall",
	})

	wantKey := "    gridscale_filesystem:\n"
	if !strings.Contains(got, wantKey) {
		t.Fatalf("expected stub key %q in output, got:\n%s", wantKey, got)
	}
	backupAt := strings.Index(got, "    gridscale_backupschedule:\n")
	fsAt := strings.Index(got, wantKey)
	fwAt := strings.Index(got, "    gridscale_firewall:\n")
	if backupAt < 0 || fsAt < 0 || fwAt < 0 {
		t.Fatalf("missing expected keys in output:\n%s", got)
	}
	if !(backupAt < fsAt && fsAt < fwAt) {
		t.Errorf("stub not inserted alphabetically: backup=%d filesystem=%d firewall=%d\n%s", backupAt, fsAt, fwAt, got)
	}
	if !strings.Contains(got, "        name: gridscale_filesystem\n") {
		t.Errorf("stub missing name field, got:\n%s", got)
	}
	if !strings.Contains(got, "        importStatements: []\n") {
		t.Errorf("stub missing importStatements, got:\n%s", got)
	}
}

// TestInjectMissingResourceStubs_Idempotent asserts a second pass is a no-op
// when the stub (or a real scraped entry) is already present.
func TestInjectMissingResourceStubs_Idempotent(t *testing.T) {
	in := "resources:\n" +
		"    gridscale_filesystem:\n" +
		"        name: already-there\n" +
		"        title: custom\n"

	got := injectMissingResourceStubs(in, []string{"gridscale_filesystem"})
	if got != in {
		t.Errorf("existing entry should be left untouched\nwant:\n%s\ngot:\n%s", in, got)
	}

	once := injectMissingResourceStubs("resources:\n", []string{"gridscale_filesystem"})
	twice := injectMissingResourceStubs(once, []string{"gridscale_filesystem"})
	if once != twice {
		t.Errorf("inject(inject(x)) != inject(x)\nonce:\n%s\ntwice:\n%s", once, twice)
	}
}

// TestInjectMissingResourceStubs_NoSchemaNoop asserts an empty schema list
// leaves content unchanged (tool must not invent keys without schema).
func TestInjectMissingResourceStubs_NoSchemaNoop(t *testing.T) {
	in := "resources:\n    gridscale_server:\n        name: s\n"
	if got := injectMissingResourceStubs(in, nil); got != in {
		t.Errorf("nil schema keys should be a no-op\ngot:\n%s", got)
	}
}

// TestFixMetadata_ReKeyRules asserts each of the eight title→gridscale_* key
// rewrites maps correctly: the gridscale_* key appears and the original title
// key is gone. Each case feeds a minimal YAML snippet with the title key line.
func TestFixMetadata_ReKeyRules(t *testing.T) {
	cases := []struct {
		name     string
		titleKey string // whole-line key as scraped (4-space indent + colon)
		wantKey  string // corrected gridscale_* key line
	}{
		{"isoimage", "    ISO Image:", "    gridscale_isoimage:"},
		{"marketplace_import", "    imported marketplace application:", "    gridscale_marketplace_application_import:"},
		{"marketplace", "    marketplace application:", "    gridscale_marketplace_application:"},
		{"object_storage_accesskey", "    object storage access key:", "    gridscale_object_storage_accesskey:"},
		{"object_storage_bucket", "    object storage bucket:", "    gridscale_object_storage_bucket:"},
		{"backupschedule", "    storage backup schedule:", "    gridscale_backupschedule:"},
		{"snapshot", "    storage snapshot:", "    gridscale_snapshot:"},
		{"snapshotschedule", "    storage snapshot schedule:", "    gridscale_snapshotschedule:"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			in := "resources:\n" + tc.titleKey + "\n      name: something\n"
			got := fixMetadata(in)

			if !strings.Contains(got, tc.wantKey+"\n") {
				t.Errorf("expected corrected key %q in output, got:\n%s", tc.wantKey, got)
			}
			if strings.Contains(got, tc.titleKey+"\n") {
				t.Errorf("original title key %q should be gone, got:\n%s", tc.titleKey, got)
			}
		})
	}
}

// TestFixMetadata_SubstringDiscrimination is the sharpest over-reach guard.
// "marketplace application" is a substring of "imported marketplace
// application"; the tool matches whole lines precisely so the two keys map to
// distinct targets and neither clobbers the other. A regression to substring
// matching would corrupt the imported entry — this test catches that.
func TestFixMetadata_SubstringDiscrimination(t *testing.T) {
	in := "resources:\n" +
		"    imported marketplace application:\n" +
		"      name: a\n" +
		"    marketplace application:\n" +
		"      name: b\n"

	got := fixMetadata(in)

	if !strings.Contains(got, "    gridscale_marketplace_application_import:\n") {
		t.Errorf("imported key should map to _import target, got:\n%s", got)
	}
	if !strings.Contains(got, "    gridscale_marketplace_application:\n") {
		t.Errorf("plain key should map to _application target, got:\n%s", got)
	}
	// Neither original title key should survive.
	if strings.Contains(got, "imported marketplace application:") {
		t.Errorf("imported title key should be gone, got:\n%s", got)
	}
	// The corrected _import key legitimately contains "gridscale_marketplace_application"
	// so we assert the *title* form is gone, not the substring.
	if strings.Contains(got, "\n    marketplace application:\n") {
		t.Errorf("plain title key should be gone, got:\n%s", got)
	}
}

// TestFixMetadata_CapacityDescription asserts the corrupted capacity
// description (doubled "ISO Image" segment) is corrected to the canonical
// "template" wording.
func TestFixMetadata_CapacityDescription(t *testing.T) {
	in := "        capacity:\n" +
		"          description: The capacity of a storage/ISO Image/ISO Image/snapshot in GB.\n"
	want := "The capacity of a storage/ISO Image/template/snapshot in GB."

	got := fixMetadata(in)

	if !strings.Contains(got, want) {
		t.Errorf("expected corrected capacity description %q, got:\n%s", want, got)
	}
	if strings.Contains(got, "storage/ISO Image/ISO Image/snapshot") {
		t.Errorf("corrupted capacity description should be gone, got:\n%s", got)
	}
}

// TestFixMetadata_Idempotent asserts the transform is a no-op on already-fixed
// input and that fix(fix(x)) == fix(x).
func TestFixMetadata_Idempotent(t *testing.T) {
	// Already-corrected content: gridscale_* keys and canonical capacity text.
	fixed := "resources:\n" +
		"    gridscale_isoimage:\n" +
		"      name: iso\n" +
		"    gridscale_marketplace_application_import:\n" +
		"      name: imp\n" +
		"        capacity:\n" +
		"          description: The capacity of a storage/ISO Image/template/snapshot in GB.\n"

	if got := fixMetadata(fixed); got != fixed {
		t.Errorf("transform should be a no-op on already-fixed input\nwant:\n%s\ngot:\n%s", fixed, got)
	}

	// fix(fix(x)) == fix(x) on dirty input.
	dirty := "resources:\n" +
		"    ISO Image:\n" +
		"      name: iso\n" +
		"        capacity:\n" +
		"          description: The capacity of a storage/ISO Image/ISO Image/snapshot in GB.\n"
	once := fixMetadata(dirty)
	twice := fixMetadata(once)
	if once != twice {
		t.Errorf("fix(fix(x)) != fix(x)\nonce:\n%s\ntwice:\n%s", once, twice)
	}
}

// TestFixMetadata_NoOverReach asserts the line-scoped matching does not clobber
// a title string that appears as a value (name:/title:/description) or as a key
// with deeper indent or trailing content — only a bare whole-line key matches.
func TestFixMetadata_NoOverReach(t *testing.T) {
	cases := []struct {
		name string
		in   string
	}{
		{
			name: "as name value",
			in:   "      name: ISO Image\n",
		},
		{
			name: "as title value",
			in:   "      title: marketplace application\n",
		},
		{
			name: "inside a description",
			in:   "      description: Uses the storage snapshot for backups.\n",
		},
		{
			name: "deeper indent key (not a resource key)",
			in:   "        ISO Image:\n",
		},
		{
			name: "key with trailing inline value",
			in:   "    ISO Image: inlinevalue\n",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := fixMetadata(tc.in); got != tc.in {
				t.Errorf("input should be unchanged (no over-reach)\nin:  %q\ngot: %q", tc.in, got)
			}
		})
	}
}

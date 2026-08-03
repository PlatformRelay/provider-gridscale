package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

func seedExternalNameCorpus(f *testing.F) {
	f.Helper()
	// Seed corpus: representative + edge inputs.
	f.Add("")                                                       // empty id
	f.Add("6a5d0f0e-8a2b-4c7d-9e1f-0123456789ab")                   // normal UUID
	f.Add("name/with/slashes")                                      // embedded separators
	f.Add("\u00fcn\u00eec\u00f6d\u00e9-r\u00e9source-\u540d\u524d") // unicode
	f.Add("id\nwith\nnewlines")                                     // embedded newlines
	f.Add("id\x00with\x00nul")                                      // embedded NUL bytes
	f.Add(string([]byte{0xff, 0xfe, 0xfd}))                         // invalid UTF-8
	f.Add("   surrounding whitespace   ")                           // whitespace
}

func resolveGetExternalName(f *testing.F) ujconfig.GetExternalNameFn {
	f.Helper()
	opt := ExternalNameConfigurations()
	r := &ujconfig.Resource{}
	opt(r)
	getExternalName := r.ExternalName.GetExternalNameFn
	if getExternalName == nil {
		f.Fatal("GetExternalNameFn is nil; nothing to fuzz")
	}
	return getExternalName
}

func assertExternalNameStringID(t *testing.T, get ujconfig.GetExternalNameFn, id string) {
	t.Helper()
	// Case A: "id" present as a string. This exercises the successful
	// type-assertion path and the empty-string branch inside
	// IDAsExternalName.
	got, err := get(map[string]any{"id": id})
	if err != nil {
		// Contract: the stub NEVER surfaces an error (it swallows the
		// underlying IDAsExternalName error).
		t.Fatalf("GetExternalNameFn(id=%q): want nil error, got %v", id, err)
	}
	if id == "" {
		// Empty id triggers the error branch upstream, swallowed to "".
		if got != "" {
			t.Fatalf("GetExternalNameFn(id=\"\"): want empty external name, got %q", got)
		}
		return
	}
	// Non-empty id must be returned verbatim, byte-for-byte. This
	// catches any accidental mutation/normalization of the identifier
	// (e.g. trimming, truncation, re-encoding of non-UTF8 bytes).
	if got != id {
		t.Fatalf("GetExternalNameFn(id=%q): want verbatim %q, got %q", id, id, got)
	}
}

func assertExternalNameNonStringID(t *testing.T, get ujconfig.GetExternalNameFn, id string) {
	t.Helper()
	// Case B: "id" present but NOT a string (wrong dynamic type). The
	// type assertion in IDAsExternalName must fail gracefully — no panic —
	// and the stub swallows the error, yielding ("", nil).
	got, err := get(map[string]any{"id": []byte(id)})
	if err != nil {
		t.Fatalf("GetExternalNameFn(non-string id): want nil error, got %v", err)
	}
	if got != "" {
		t.Fatalf("GetExternalNameFn(non-string id): want empty external name, got %q", got)
	}
}

func assertExternalNameMissingID(t *testing.T, get ujconfig.GetExternalNameFn, id string) {
	t.Helper()
	// Case C: "id" key absent entirely, but the fuzzed string occupies an
	// unrelated key. The stub must ignore it and yield ("", nil).
	got, err := get(map[string]any{"other": id})
	if err != nil {
		t.Fatalf("GetExternalNameFn(missing id): want nil error, got %v", err)
	}
	if got != "" {
		t.Fatalf("GetExternalNameFn(missing id): want empty external name, got %q", got)
	}
}

// FuzzGetExternalName fuzzes the custom GetExternalNameFn stub installed by
// idWithStub() — the highest-value hand-authored parsing surface in this
// generated provider. The stub extracts the provider-assigned identifier from
// Terraform state (a map[string]any) via IDAsExternalName and deliberately
// swallows the "id not found" error, yielding ("", nil) instead.
//
// The fuzzer drives the string content that lands under the "id" key
// (adversarial: empty, huge, embedded NULs/separators, non-UTF8, unicode) and
// asserts the stub's contract holds and never panics.
func FuzzGetExternalName(f *testing.F) {
	seedExternalNameCorpus(f)
	getExternalName := resolveGetExternalName(f)

	f.Fuzz(func(t *testing.T, id string) {
		assertExternalNameStringID(t, getExternalName, id)
		assertExternalNameNonStringID(t, getExternalName, id)
		assertExternalNameMissingID(t, getExternalName, id)
	})
}

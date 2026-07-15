package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

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
	// Seed corpus: representative + edge inputs.
	f.Add("")                                     // empty id
	f.Add("6a5d0f0e-8a2b-4c7d-9e1f-0123456789ab") // normal UUID
	f.Add("name/with/slashes")                    // embedded separators
	f.Add("ünîcödé-résource-名前")                    // unicode
	f.Add("id\nwith\nnewlines")                   // embedded newlines
	f.Add("id\x00with\x00nul")                    // embedded NUL bytes
	f.Add(string([]byte{0xff, 0xfe, 0xfd}))       // invalid UTF-8
	f.Add("   surrounding whitespace   ")         // whitespace

	// Resolve the stub once; it is a pure function of its tfstate argument.
	opt := ExternalNameConfigurations()
	r := &ujconfig.Resource{}
	opt(r)
	getExternalName := r.ExternalName.GetExternalNameFn
	if getExternalName == nil {
		f.Fatal("GetExternalNameFn is nil; nothing to fuzz")
	}

	f.Fuzz(func(t *testing.T, id string) {
		// Case A: "id" present as a string. This exercises the successful
		// type-assertion path and the empty-string branch inside
		// IDAsExternalName.
		gotA, errA := getExternalName(map[string]any{"id": id})
		if errA != nil {
			// Contract: the stub NEVER surfaces an error (it swallows the
			// underlying IDAsExternalName error).
			t.Fatalf("GetExternalNameFn(id=%q): want nil error, got %v", id, errA)
		}
		if id == "" {
			// Empty id triggers the error branch upstream, swallowed to "".
			if gotA != "" {
				t.Fatalf("GetExternalNameFn(id=\"\"): want empty external name, got %q", gotA)
			}
		} else {
			// Non-empty id must be returned verbatim, byte-for-byte. This
			// catches any accidental mutation/normalization of the identifier
			// (e.g. trimming, truncation, re-encoding of non-UTF8 bytes).
			if gotA != id {
				t.Fatalf("GetExternalNameFn(id=%q): want verbatim %q, got %q", id, id, gotA)
			}
		}

		// Case B: "id" present but NOT a string (wrong dynamic type). The
		// type assertion in IDAsExternalName must fail gracefully — no panic —
		// and the stub swallows the error, yielding ("", nil).
		gotB, errB := getExternalName(map[string]any{"id": []byte(id)})
		if errB != nil {
			t.Fatalf("GetExternalNameFn(non-string id): want nil error, got %v", errB)
		}
		if gotB != "" {
			t.Fatalf("GetExternalNameFn(non-string id): want empty external name, got %q", gotB)
		}

		// Case C: "id" key absent entirely, but the fuzzed string occupies an
		// unrelated key. The stub must ignore it and yield ("", nil).
		gotC, errC := getExternalName(map[string]any{"other": id})
		if errC != nil {
			t.Fatalf("GetExternalNameFn(missing id): want nil error, got %v", errC)
		}
		if gotC != "" {
			t.Fatalf("GetExternalNameFn(missing id): want empty external name, got %q", gotC)
		}
	})
}

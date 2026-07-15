package config

import (
	"regexp"
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// TestExternalNameConfigured asserts the include-list is the ".+" catch-all so
// that every resource the gridscale provider exposes receives the external-name
// configuration.
func TestExternalNameConfigured(t *testing.T) {
	got := ExternalNameConfigured()

	if want := 1; len(got) != want {
		t.Fatalf("ExternalNameConfigured(): want %d entry, got %d (%v)", want, len(got), got)
	}
	if want := ".+"; got[0] != want {
		t.Fatalf("ExternalNameConfigured()[0]: want %q, got %q", want, got[0])
	}

	// The include-list entries are treated as regular expressions by upjet, so
	// prove ".+" actually matches representative resource names.
	re, err := regexp.Compile("^(?:" + got[0] + ")$")
	if err != nil {
		t.Fatalf("ExternalNameConfigured()[0] is not a valid regexp: %v", err)
	}
	for _, name := range []string{"gridscale_server", "gridscale_mysql8_0", "gridscale_storage_import"} {
		if !re.MatchString(name) {
			t.Errorf("catch-all %q did not match resource %q", got[0], name)
		}
	}
}

// TestExternalNameStrategy asserts the external-name strategy the provider
// applies to every resource: the provider-assigned identifier (UUID) with a
// custom GetExternalNameFn stub that never surfaces an error.
func TestExternalNameStrategy(t *testing.T) {
	// ExternalNameConfigurations must produce a ResourceOption that installs the
	// idWithStub external name on a resource.
	opt := ExternalNameConfigurations()
	if opt == nil {
		t.Fatal("ExternalNameConfigurations() returned a nil ResourceOption")
	}

	r := &ujconfig.Resource{}
	opt(r)
	en := r.ExternalName

	// IdentifierFromProvider disables the name initializer because the
	// identifier is assigned by the remote provider, not chosen by the user.
	// This scalar field is the comparable marker that IdentifierFromProvider is
	// the base strategy.
	if !en.DisableNameInitializer {
		t.Error("ExternalName.DisableNameInitializer: want true (IdentifierFromProvider strategy), got false")
	}

	// GetExternalNameFn is required so upjet can extract the external name from
	// TF state.
	if en.GetExternalNameFn == nil {
		t.Fatal("ExternalName.GetExternalNameFn is nil")
	}

	// Contract 1: a populated "id" in TF state is returned verbatim as the
	// external name (the provider-assigned UUID).
	const uuid = "6a5d0f0e-8a2b-4c7d-9e1f-0123456789ab"
	name, err := en.GetExternalNameFn(map[string]any{"id": uuid})
	if err != nil {
		t.Fatalf("GetExternalNameFn(id=%q): unexpected error: %v", uuid, err)
	}
	if name != uuid {
		t.Errorf("GetExternalNameFn(id=%q): want %q, got %q", uuid, uuid, name)
	}

	// Contract 2 (the real contract of idWithStub): when "id" is absent, the
	// stub swallows the error that the underlying IDAsExternalName would return
	// and yields ("", nil) instead. Locking this in prevents a regression back
	// to error-surfacing behavior.
	name, err = en.GetExternalNameFn(map[string]any{})
	if err != nil {
		t.Errorf("GetExternalNameFn(empty): want nil error (stub swallows it), got %v", err)
	}
	if name != "" {
		t.Errorf("GetExternalNameFn(empty): want empty external name, got %q", name)
	}
}

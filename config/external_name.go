package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// gridscale objects are identified by a provider-assigned UUID, so
// IdentifierFromProvider is the correct external-name strategy for every
// resource the gridscale Terraform provider exposes.

// ExternalNameConfigurations sets the external name of every generated resource
// to the provider-assigned identifier (UUID).
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		r.ExternalName = idWithStub()
	}
}

func idWithStub() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigured returns the include-list: every resource the gridscale
// provider exposes (".+" matches all).
func ExternalNameConfigured() []string {
	return []string{".+"}
}

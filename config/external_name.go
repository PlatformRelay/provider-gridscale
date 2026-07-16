package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// IdentifierFromProvider is the correct external-name strategy for every
// resource the gridscale Terraform provider exposes: Upjet stores whatever the
// upstream provider passes to Terraform's SetId and uses it verbatim as the
// external-name. For 30 of the 32 resources that identifier is a UUID assigned
// by gridscale, but two are not:
//
//   - object_storage_bucket uses the composite `<s3_host>/<bucket_name>`
//     (upstream resource_gridscale_bucket.go:249-250, imported via
//     ImportStatePassthrough), and
//   - object_storage_accesskey uses the access-key string itself
//     (upstream resource_gridscale_objectstorage.go:105).
//
// Both use ImportStatePassthrough upstream, so import/observe work with these
// identifiers as-is and no NameAsIdentifier/GetIDFn override is needed.
// See docs/adr/0006-external-name-import-formats.md.

// ExternalNameConfigurations sets the external name of every generated resource
// to the provider-assigned identifier (see the strategy note above for the two
// non-UUID formats).
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

package config

import "testing"

// TestSensitiveFieldOverrides asserts that the credential fields the upstream
// gridscale schema fails to flag as Sensitive (finding U-1) are forced sensitive
// by configureSensitiveFields, so upjet routes them into SecretKeySelectors /
// connection secrets instead of persisting them as plaintext in the CRDs.
//
// It builds the fully-configured provider (GetProvider runs ConfigureResources,
// which applies the configurators) and inspects the resulting Terraform schema.
func TestSensitiveFieldOverrides(t *testing.T) {
	p := GetProvider()

	for resource, fields := range sensitiveFieldOverrides {
		r, ok := p.Resources[resource]
		if !ok {
			t.Errorf("resource %q not present in the configured provider", resource)
			continue
		}
		for _, f := range fields {
			s, ok := r.TerraformResource.Schema[f]
			if !ok {
				t.Errorf("%s: field %q missing from Terraform schema (upstream rename?)", resource, f)
				continue
			}
			if !s.Sensitive {
				t.Errorf("%s.%s: want Sensitive=true (U-1 override), got false — "+
					"credential would be generated as a plaintext CRD field", resource, f)
			}
		}
	}
}

// TestSensitiveOverridesCoverKnownU1Fields guards the override list itself against
// accidental shrinkage: U-1 is exactly these five fields across three resources.
func TestSensitiveOverridesCoverKnownU1Fields(t *testing.T) {
	want := map[string]int{
		"gridscale_k8s":        2, // log_delivery_access_key, log_delivery_secret_key
		"gridscale_postgresql": 2, // pgaudit_log_access_key, pgaudit_log_secret_key
		"gridscale_server":     1, // console_token
	}
	if len(sensitiveFieldOverrides) != len(want) {
		t.Fatalf("sensitiveFieldOverrides covers %d resources, want %d", len(sensitiveFieldOverrides), len(want))
	}
	for resource, n := range want {
		if got := len(sensitiveFieldOverrides[resource]); got != n {
			t.Errorf("%s: want %d overridden field(s), got %d", resource, n, got)
		}
	}
}

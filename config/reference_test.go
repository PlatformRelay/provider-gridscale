package config

import (
	"testing"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// wantReferences is the full set of cross-resource reference edges the
// provider must wire (story E7-S01 / audit finding ARCH-1). Every edge points
// at a UUID-keyed target resolved via the referenced resource's external name,
// except loadbalancer backend_server.host, which resolves the referenced
// IPv4 object's ".ip" observation via a custom extractor (asserted separately
// in TestReferenceExtractors).
//
// Field paths use Terraform attribute syntax; nested block fields are
// dot-joined (e.g. "network.object_uuid"). Target names reuse the tf*
// constants from references.go where one exists; TerraformName resolution is
// independently verified at generation time (unknown names fail `make
// generate`), so the shared constants cannot mask a typo silently.
const fieldNetworkUUID = "network_uuid"

var wantReferences = map[string]map[string]string{
	"gridscale_server": {
		"ipv4":                           tfIPv4,
		"ipv6":                           tfIPv6,
		"isoimage":                       "gridscale_isoimage",
		"network.object_uuid":            tfNetwork,
		"network.firewall_template_uuid": "gridscale_firewall",
		"storage.object_uuid":            tfStorage,
	},
	"gridscale_loadbalancer": {
		"listen_ipv4_uuid":                 tfIPv4,
		"listen_ipv6_uuid":                 tfIPv6,
		"forwarding_rule.certificate_uuid": "gridscale_ssl_certificate",
		"backend_server.host":              tfIPv4,
	},
	"gridscale_template": {
		"snapshot_uuid": "gridscale_snapshot",
	},
	"gridscale_storage_clone": {
		"source_storage_id": tfStorage,
	},
	"gridscale_snapshot": {
		"storage_uuid": tfStorage,
	},
	"gridscale_snapshotschedule": {
		"storage_uuid": tfStorage,
	},
	"gridscale_backupschedule": {
		"storage_uuid": tfStorage,
	},
	// PaaS family: network attachment via the non-deprecated network_uuid.
	"gridscale_k8s":         {fieldNetworkUUID: tfNetwork},
	"gridscale_paas":        {fieldNetworkUUID: tfNetwork},
	"gridscale_postgresql":  {fieldNetworkUUID: tfNetwork},
	"gridscale_mysql":       {fieldNetworkUUID: tfNetwork},
	tfMySQL8:                {fieldNetworkUUID: tfNetwork},
	"gridscale_mariadb":     {fieldNetworkUUID: tfNetwork},
	"gridscale_memcached":   {fieldNetworkUUID: tfNetwork},
	"gridscale_redis_cache": {fieldNetworkUUID: tfNetwork},
	"gridscale_redis_store": {fieldNetworkUUID: tfNetwork},
	"gridscale_sqlserver":   {fieldNetworkUUID: tfNetwork},
	"gridscale_filesystem":  {fieldNetworkUUID: tfNetwork},
}

// assertReferences checks every expected edge on the given (configured)
// provider: the reference must exist and carry the expected non-empty
// TerraformName (the deprecated Type form is not used in this provider).
func assertReferences(t *testing.T, pc *ujconfig.Provider) {
	t.Helper()
	for tfName, edges := range wantReferences {
		r, ok := pc.Resources[tfName]
		if !ok {
			t.Errorf("resource %q not present in configured provider", tfName)
			continue
		}
		for path, wantTarget := range edges {
			ref, ok := r.References[path]
			if !ok {
				t.Errorf("%s: reference for field path %q not configured", tfName, path)
				continue
			}
			if ref.TerraformName != wantTarget {
				t.Errorf("%s: reference %q TerraformName = %q, want %q",
					tfName, path, ref.TerraformName, wantTarget)
			}
		}
	}
}

// assertNoDeprecatedSecurityZoneRefs ensures the deprecated PaaS
// security_zone_uuid field is NOT wired as a reference on any resource:
// upstream deprecates it in favor of network_uuid, and wiring it would steer
// users onto an API gridscale is removing (E7-S01 scope; research Q3/Q5).
func assertNoDeprecatedSecurityZoneRefs(t *testing.T, pc *ujconfig.Provider) {
	t.Helper()
	for tfName, r := range pc.Resources {
		for path := range r.References {
			if path == "security_zone_uuid" {
				t.Errorf("%s: deprecated field %q must not carry a reference (use %s)",
					tfName, path, fieldNetworkUUID)
			}
		}
	}
}

// TestReferences asserts the cross-resource reference wiring on the
// cluster-scoped provider configuration.
func TestReferences(t *testing.T) {
	pc := GetProvider()
	if pc == nil {
		t.Fatal("GetProvider() returned nil")
	}
	assertReferences(t, pc)
	assertNoDeprecatedSecurityZoneRefs(t, pc)
}

// TestReferencesNamespaced asserts the same wiring on the namespaced provider
// configuration, since configureReferences is applied by both constructors.
func TestReferencesNamespaced(t *testing.T) {
	pc := GetProviderNamespaced()
	if pc == nil {
		t.Fatal("GetProviderNamespaced() returned nil")
	}
	assertReferences(t, pc)
	assertNoDeprecatedSecurityZoneRefs(t, pc)
}

// TestReferenceExtractors asserts that the one non-UUID edge —
// loadbalancer backend_server.host, whose value is the referenced IPv4
// object's IP address, not its UUID — carries a custom extractor. The default
// extractor returns the external name (UUID), which would resolve to the
// wrong value for a host field.
func TestReferenceExtractors(t *testing.T) {
	for name, pc := range map[string]*ujconfig.Provider{
		"cluster":    GetProvider(),
		"namespaced": GetProviderNamespaced(),
	} {
		t.Run(name, func(t *testing.T) {
			r, ok := pc.Resources["gridscale_loadbalancer"]
			if !ok {
				t.Fatal("gridscale_loadbalancer not present in configured provider")
			}
			ref, ok := r.References["backend_server.host"]
			if !ok {
				t.Fatal("reference for backend_server.host not configured")
			}
			if ref.Extractor == "" {
				t.Error("backend_server.host reference must set a custom Extractor " +
					"(default external-name extractor would yield the UUID, not the IP)")
			}
		})
	}
}

package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Terraform resource names referenced by more than one edge, extracted to
// constants to keep the reference wiring DRY.
const (
	tfIPv4         = "gridscale_ipv4"
	tfIPv6         = "gridscale_ipv6"
	tfNetwork      = "gridscale_network"
	tfStorage      = "gridscale_storage"
	tfSecurityZone = "gridscale_paas_securityzone"
)

// configureReferences wires cross-resource references so that fields which
// hold another gridscale object's UUID resolve from a Crossplane reference or
// selector instead of a raw ID. Every referenced gridscale resource uses the
// provider-assigned UUID as its external name (IdentifierFromProvider), so the
// default extractor (external name) yields the correct ".id" value for all of
// the edges below.
//
// Only UUID-valued edges are wired here. Edges that interpolate a non-UUID
// attribute (e.g. an IPv4 object's ".ip" address, as seen in the loadbalancer
// backend_server "host" field) are intentionally left out: the default
// extractor returns the external name (UUID), not the IP, so a reference there
// would resolve to the wrong value.
func configureReferences(p *ujconfig.Provider) {
	// gridscale_server: attachment fields pointing at other objects.
	p.AddResourceConfigurator("gridscale_server", func(r *ujconfig.Resource) {
		r.References["ipv4"] = ujconfig.Reference{
			TerraformName: tfIPv4,
		}
		r.References["ipv6"] = ujconfig.Reference{
			TerraformName: tfIPv6,
		}
		r.References["isoimage"] = ujconfig.Reference{
			TerraformName: "gridscale_isoimage",
		}
		// nested network[] block -> gridscale_network
		r.References["network.object_uuid"] = ujconfig.Reference{
			TerraformName: tfNetwork,
		}
		// nested storage[] block -> gridscale_storage
		r.References["storage.object_uuid"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})

	// gridscale_loadbalancer: listener IPs and per-rule SSL certificate.
	p.AddResourceConfigurator("gridscale_loadbalancer", func(r *ujconfig.Resource) {
		r.References["listen_ipv4_uuid"] = ujconfig.Reference{
			TerraformName: tfIPv4,
		}
		r.References["listen_ipv6_uuid"] = ujconfig.Reference{
			TerraformName: tfIPv6,
		}
		// nested forwarding_rule[] block -> gridscale_ssl_certificate
		r.References["forwarding_rule.certificate_uuid"] = ujconfig.Reference{
			TerraformName: "gridscale_ssl_certificate",
		}
	})

	// gridscale_template: base snapshot.
	p.AddResourceConfigurator("gridscale_template", func(r *ujconfig.Resource) {
		r.References["snapshot_uuid"] = ujconfig.Reference{
			TerraformName: "gridscale_snapshot",
		}
	})

	// gridscale_storage_clone: source storage.
	p.AddResourceConfigurator("gridscale_storage_clone", func(r *ujconfig.Resource) {
		r.References["source_storage_id"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})

	// Snapshot / schedule resources anchored to a storage.
	p.AddResourceConfigurator("gridscale_snapshot", func(r *ujconfig.Resource) {
		r.References["storage_uuid"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})
	p.AddResourceConfigurator("gridscale_snapshotschedule", func(r *ujconfig.Resource) {
		r.References["storage_uuid"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})
	p.AddResourceConfigurator("gridscale_backupschedule", func(r *ujconfig.Resource) {
		r.References["storage_uuid"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})

	// PaaS / k8s network + security-zone edges. These are inferred from the
	// schema field names (not example-backed like the edges above) but map
	// cleanly onto the corresponding gridscale resources.
	p.AddResourceConfigurator("gridscale_paas", func(r *ujconfig.Resource) {
		r.References["network_uuid"] = ujconfig.Reference{
			TerraformName: tfNetwork,
		}
		r.References["security_zone_uuid"] = ujconfig.Reference{
			TerraformName: tfSecurityZone,
		}
	})
	p.AddResourceConfigurator("gridscale_k8s", func(r *ujconfig.Resource) {
		r.References["security_zone_uuid"] = ujconfig.Reference{
			TerraformName: tfSecurityZone,
		}
	})
}

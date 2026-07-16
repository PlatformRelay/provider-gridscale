package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// Terraform resource names referenced by more than one edge, extracted to
// constants to keep the reference wiring DRY.
const (
	tfIPv4    = "gridscale_ipv4"
	tfIPv6    = "gridscale_ipv6"
	tfMySQL8  = "gridscale_mysql8_0"
	tfNetwork = "gridscale_network"
	tfStorage = "gridscale_storage"
)

// ipExtractor resolves the referenced IPv4 object's observed ".ip" address
// (status.atProvider.ip) instead of its external name. Used for the one edge
// whose Terraform value is an IP address rather than a UUID.
const ipExtractor = `github.com/crossplane/upjet/v2/pkg/resource.ExtractParamPath("ip",true)`

// paasFamily lists every PaaS-style resource sharing the network_uuid
// attachment field (upstream docs: k8s.html.md, paas.html.md, ...).
var paasFamily = []string{
	"gridscale_k8s",
	"gridscale_paas",
	"gridscale_postgresql",
	"gridscale_mysql",
	tfMySQL8,
	"gridscale_mariadb",
	"gridscale_memcached",
	"gridscale_redis_cache",
	"gridscale_redis_store",
	"gridscale_sqlserver",
	"gridscale_filesystem",
}

// configureReferences wires cross-resource references so that fields which
// hold another gridscale object's UUID resolve from a Crossplane reference or
// selector instead of a raw ID (ADR-0005). Every referenced gridscale
// resource uses the provider-assigned UUID as its external name
// (IdentifierFromProvider), so the default extractor (external name) yields
// the correct ".id" value for all UUID-keyed edges below.
//
// Two deliberate exceptions:
//
//   - loadbalancer backend_server.host holds the referenced IPv4 object's IP
//     address, not its UUID, so it carries a custom ExtractParamPath("ip")
//     extractor reading the observed address from status.atProvider.
//   - the deprecated PaaS security_zone_uuid field is intentionally NOT wired
//     (upstream deprecates it in favor of network_uuid; a reference would
//     steer users onto an API gridscale is removing).
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
		// nested network[] block -> gridscale_network / gridscale_firewall
		r.References["network.object_uuid"] = ujconfig.Reference{
			TerraformName: tfNetwork,
		}
		r.References["network.firewall_template_uuid"] = ujconfig.Reference{
			TerraformName: "gridscale_firewall",
		}
		// nested storage[] block -> gridscale_storage
		r.References["storage.object_uuid"] = ujconfig.Reference{
			TerraformName: tfStorage,
		}
	})

	// gridscale_loadbalancer: listener IPs, backend hosts and per-rule SSL
	// certificate.
	p.AddResourceConfigurator("gridscale_loadbalancer", func(r *ujconfig.Resource) {
		r.References["listen_ipv4_uuid"] = ujconfig.Reference{
			TerraformName: tfIPv4,
		}
		r.References["listen_ipv6_uuid"] = ujconfig.Reference{
			TerraformName: tfIPv6,
		}
		// nested backend_server[] block: host is the referenced IPv4's IP
		// address (non-UUID edge -> custom extractor).
		r.References["backend_server.host"] = ujconfig.Reference{
			TerraformName: tfIPv4,
			Extractor:     ipExtractor,
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
	for _, name := range []string{
		"gridscale_snapshot",
		"gridscale_snapshotschedule",
		"gridscale_backupschedule",
	} {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["storage_uuid"] = ujconfig.Reference{
				TerraformName: tfStorage,
			}
		})
	}

	// PaaS family: network attachment. Only the non-deprecated network_uuid is
	// wired; see the function comment for why security_zone_uuid is skipped.
	for _, name := range paasFamily {
		p.AddResourceConfigurator(name, func(r *ujconfig.Resource) {
			r.References["network_uuid"] = ujconfig.Reference{
				TerraformName: tfNetwork,
			}
		})
	}
}

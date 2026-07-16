package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// configureLoadbalancer corrects the gridscale_loadbalancer.status field
// (finding LB-1). Upstream (v2.3.0) declares it Optional with Default "active",
// but status is server-derived — it is only ever set from the API response, never
// read from user config in Create/Update. As declared it causes a perpetual plan
// diff whenever the load balancer reports a status other than "active", and it is
// the only resource in the provider whose "status" is Optional rather than
// Computed (every other resource marks it Computed).
//
// Forcing Computed=true (and clearing Optional/Default) moves it to
// status.atProvider as an observed-only field, matching the other 29 resources and
// eliminating the drift. Applied as a config-time override rather than a
// schema.json edit so `make generate` (which re-dumps schema.json from the pinned
// provider binary) reproduces it and check-diff stays clean. Remove once fixed
// upstream and re-vendored.
func configureLoadbalancer(p *ujconfig.Provider) {
	p.AddResourceConfigurator("gridscale_loadbalancer", func(r *ujconfig.Resource) {
		if s, ok := r.TerraformResource.Schema["status"]; ok {
			s.Computed = true
			s.Optional = false
			s.Default = nil
		}
	})
}

package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// sensitiveFieldOverrides lists top-level credential fields that the upstream
// gridscale/terraform-provider-gridscale schema (v2.3.0) fails to flag as
// Sensitive (finding U-1). Each holds a real, round-tripped object-storage/console
// credential — the sibling password/kubeconfig fields on the same resources are
// correctly marked Sensitive, so this is an upstream oversight, not by design.
var sensitiveFieldOverrides = map[string][]string{
	// S3 log-delivery credentials for the managed k8s cluster.
	"gridscale_k8s": {"log_delivery_access_key", "log_delivery_secret_key"},
	// S3 credentials for pgaudit log delivery.
	"gridscale_postgresql": {"pgaudit_log_access_key", "pgaudit_log_secret_key"},
	// Token granting VNC console access to the server.
	"gridscale_server": {"console_token"},
}

// configureSensitiveFields forces Sensitive=true on the fields in
// sensitiveFieldOverrides. Without it, upjet renders these credentials as plain
// strings in spec.forProvider (and status.atProvider) of the generated CRDs, so
// the secret values are persisted in plaintext in the Kubernetes resource; with
// it, each is routed through a SecretKeySelector / connection secret instead.
//
// This is applied as a config-time override rather than a config/schema.json edit
// on purpose: `make generate` re-dumps schema.json from the pinned provider binary
// (which carries the bug), so a raw schema edit would be reverted and fail
// check-diff. This configurator re-applies on every generate, keeping the fix
// durable and CI-reproducible. Remove once the fix lands upstream and is re-vendored.
func configureSensitiveFields(p *ujconfig.Provider) {
	for resource, fields := range sensitiveFieldOverrides {
		fields := fields
		p.AddResourceConfigurator(resource, func(r *ujconfig.Resource) {
			for _, f := range fields {
				if s, ok := r.TerraformResource.Schema[f]; ok {
					s.Sensitive = true
				}
			}
		})
	}
}

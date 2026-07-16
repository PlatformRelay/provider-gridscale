package config

import (
	"encoding/json"
	"regexp"
	"testing"

	"sigs.k8s.io/yaml"
)

// TestMetadataKeys asserts the embedded provider-metadata.yaml stays aligned with
// the Terraform provider schema: every metadata resource key is a valid
// gridscale_* terraform name (no spaces / uppercase), appears in schema.json
// resource_schemas, and every schema resource has a metadata entry.
//
// The last check catches scraper gaps (e.g. gridscale_filesystem) that would
// otherwise leave a CRD without scraped/stub metadata after generate.
func TestMetadataKeys(t *testing.T) {
	metaKeys := metadataResourceKeys(t)
	schemaKeys := schemaResourceKeys(t)

	keyRe := regexp.MustCompile(`^gridscale_[a-z0-9_]+$`)
	for key := range metaKeys {
		if !keyRe.MatchString(key) {
			t.Errorf("metadata key %q: want gridscale_* with only lowercase letters, digits, underscores (no spaces/uppercase)", key)
		}
		if _, ok := schemaKeys[key]; !ok {
			t.Errorf("metadata key %q is not present in schema.json resource_schemas", key)
		}
	}

	for key := range schemaKeys {
		if _, ok := metaKeys[key]; !ok {
			t.Errorf("schema resource %q has no entry in provider-metadata.yaml", key)
		}
	}
}

func metadataResourceKeys(t *testing.T) map[string]struct{} {
	t.Helper()
	var meta struct {
		Resources map[string]json.RawMessage `json:"resources"`
	}
	if err := yaml.Unmarshal([]byte(providerMetadata), &meta); err != nil {
		t.Fatalf("parse provider-metadata.yaml: %v", err)
	}
	if len(meta.Resources) == 0 {
		t.Fatal("provider-metadata.yaml has no resources")
	}
	out := make(map[string]struct{}, len(meta.Resources))
	for k := range meta.Resources {
		out[k] = struct{}{}
	}
	return out
}

func schemaResourceKeys(t *testing.T) map[string]struct{} {
	t.Helper()
	var schema struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]json.RawMessage `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal([]byte(providerSchema), &schema); err != nil {
		t.Fatalf("parse schema.json: %v", err)
	}
	out := make(map[string]struct{})
	for _, ps := range schema.ProviderSchemas {
		for k := range ps.ResourceSchemas {
			out[k] = struct{}{}
		}
	}
	if len(out) == 0 {
		t.Fatal("schema.json has no resource_schemas")
	}
	return out
}

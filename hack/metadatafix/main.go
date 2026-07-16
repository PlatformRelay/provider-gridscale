// Command metadatafix applies deterministic post-scrape corrections to
// config/provider-metadata.yaml.
//
// The upjet scraper (see apis/generate.go) keys each resource's metadata entry
// by the upstream doc page's page_title. Eight gridscale doc pages use a
// human-readable title (e.g. "ISO Image") instead of the terraform resource
// name (e.g. "gridscale_isoimage"). Because the key no longer matches a
// gridscale_* resource, upjet cannot attach the scraped argument descriptions
// to those CRDs. This tool re-keys those eight entries to their terraform
// resource names so the descriptions land on the generated CRDs.
//
// It also repairs a corrupted upstream `capacity` argument description on the
// ISO Image page ("storage/ISO Image/ISO Image/snapshot", a doubled segment),
// restoring it to the canonical "storage/ISO Image/template/snapshot" wording
// used by the sibling template/storage pages.
//
// Finally, it injects a minimal stub entry for any schema.json resource that
// the scraper omitted (currently gridscale_filesystem - no upstream doc page).
// Stubs are string-inserted in alphabetical key order so the scraper's YAML
// formatting is preserved (no round-trip churn). The step is idempotent.
//
// The corrections are exact line/string replacements so the scraper's output
// formatting is preserved verbatim (no YAML round-trip churn) and the step is
// idempotent: re-running on already-corrected content is a no-op.
//
// This runs as part of `make generate`, immediately after the scraper writes
// the file, so the fix is durable and survives regeneration (it does not rely
// on a hand edit that check-diff would revert).
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// keyRewrites maps a scraped resource-entry key line to its corrected form.
// Keys are matched as whole lines (4-space YAML indent + title + colon) so we
// never touch the same title where it legitimately appears as a name:/title:
// value or inside argumentDocs.
var keyRewrites = map[string]string{
	"    ISO Image:":                        "    gridscale_isoimage:",
	"    imported marketplace application:": "    gridscale_marketplace_application_import:",
	"    marketplace application:":          "    gridscale_marketplace_application:",
	"    object storage access key:":        "    gridscale_object_storage_accesskey:",
	"    object storage bucket:":            "    gridscale_object_storage_bucket:",
	"    storage backup schedule:":          "    gridscale_backupschedule:",
	"    storage snapshot:":                 "    gridscale_snapshot:",
	"    storage snapshot schedule:":        "    gridscale_snapshotschedule:",
}

// textRewrites maps a corrupted scraped string to its corrected form. These are
// exact substring replacements applied to the whole file.
var textRewrites = map[string]string{
	"The capacity of a storage/ISO Image/ISO Image/snapshot in GB.": "The capacity of a storage/ISO Image/template/snapshot in GB.",
}

// fixMetadata applies the deterministic corrections to the scraped metadata
// content: line-scoped key re-keying followed by file-scoped text corrections.
// It is a pure function of its input and is idempotent — re-running it on
// already-corrected content returns byte-identical output.
func fixMetadata(content string) string {
	// Line-scoped key re-keying. Whole-line matching (not substring) is
	// deliberate: it prevents "marketplace application" from clobbering the
	// embedded substring in "imported marketplace application", and prevents
	// touching a title where it legitimately appears as a name:/title: value.
	lines := strings.Split(content, "\n")
	for i, line := range lines {
		if repl, ok := keyRewrites[line]; ok {
			lines[i] = repl
		}
	}
	content = strings.Join(lines, "\n")

	// File-scoped text corrections.
	for from, to := range textRewrites {
		content = strings.ReplaceAll(content, from, to)
	}

	return content
}

// resourceKeyLine returns the whole-line YAML key form used for top-level
// entries under resources: ("    <name>:").
func resourceKeyLine(name string) string {
	return "    " + name + ":"
}

// stubResourceEntry returns a minimal metadata block for a schema resource the
// scraper did not document. Shape matches neighbouring scraped entries so upjet
// accepts the key; argumentDocs stays empty until real docs appear upstream.
func stubResourceEntry(name string) string {
	return resourceKeyLine(name) + "\n" +
		"        subCategory: \"\"\n" +
		"        description: Manages a " + name + " resource. Stub entry - upstream provider docs were not scraped for this resource.\n" +
		"        name: " + name + "\n" +
		"        title: " + name + "\n" +
		"        argumentDocs: {}\n" +
		"        importStatements: []\n"
}

// existingResourceKeys returns the set of top-level resource keys present in
// the scraped YAML (4-space indent, bare "name:" line).
func existingResourceKeys(content string) map[string]struct{} {
	out := make(map[string]struct{})
	for _, line := range strings.Split(content, "\n") {
		if !strings.HasPrefix(line, "    ") || !strings.HasSuffix(line, ":") {
			continue
		}
		// Resource keys are exactly 4 spaces; nested keys use 8+.
		if strings.HasPrefix(line, "        ") {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(line, "    "), ":")
		if name == "" || strings.Contains(name, " ") {
			// Skip non-resource lines; title-keyed scraper entries may still
			// contain spaces before re-keying - those are handled by fixMetadata.
			continue
		}
		out[name] = struct{}{}
	}
	return out
}

// injectMissingResourceStubs inserts a minimal stub for each schema resource
// that has no metadata entry yet. Stubs are placed immediately before the next
// existing resource key in alphabetical order (or appended under resources: if
// none sorts after). Idempotent when entries already exist.
func injectMissingResourceStubs(content string, schemaKeys []string) string {
	if len(schemaKeys) == 0 {
		return content
	}
	present := existingResourceKeys(content)
	missing := make([]string, 0)
	for _, key := range schemaKeys {
		if _, ok := present[key]; ok {
			continue
		}
		missing = append(missing, key)
	}
	if len(missing) == 0 {
		return content
	}
	sort.Strings(missing)

	// Known keys sorted so we can find the insertion neighbour for each stub.
	known := make([]string, 0, len(present)+len(missing))
	for k := range present {
		known = append(known, k)
	}
	known = append(known, missing...)
	sort.Strings(known)

	for _, key := range missing {
		stub := stubResourceEntry(key)
		insertBefore := ""
		for _, k := range known {
			if k > key {
				if _, ok := present[k]; ok {
					insertBefore = k
					break
				}
			}
		}
		if insertBefore != "" {
			needle := resourceKeyLine(insertBefore) + "\n"
			idx := strings.Index(content, needle)
			if idx < 0 {
				// Fall through to append if the neighbour line was reformatted.
				content += stub
			} else {
				content = content[:idx] + stub + content[idx:]
			}
		} else {
			// No later neighbour: append after the last line, preserving a
			// trailing newline if the file already had one.
			if content != "" && !strings.HasSuffix(content, "\n") {
				content += "\n"
			}
			content += stub
		}
		present[key] = struct{}{}
	}
	return content
}

// loadSchemaResourceKeys reads terraform provider schema.json and returns the
// sorted list of resource type names under resource_schemas.
func loadSchemaResourceKeys(schemaPath string) ([]string, error) {
	raw, err := os.ReadFile(schemaPath) // #nosec G304 G703 — build-time sibling of metadata
	if err != nil {
		return nil, err
	}
	var schema struct {
		ProviderSchemas map[string]struct {
			ResourceSchemas map[string]json.RawMessage `json:"resource_schemas"`
		} `json:"provider_schemas"`
	}
	if err := json.Unmarshal(raw, &schema); err != nil {
		return nil, err
	}
	keys := make([]string, 0)
	for _, ps := range schema.ProviderSchemas {
		for k := range ps.ResourceSchemas {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	return keys, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: metadatafix <path-to-provider-metadata.yaml>")
		os.Exit(2)
	}
	path := os.Args[1]

	// path is a build-time codegen argument supplied by apis/generate.go
	// (always ../config/provider-metadata.yaml), not untrusted input.
	raw, err := os.ReadFile(path) // #nosec G304 G703
	if err != nil {
		fmt.Fprintf(os.Stderr, "metadatafix: read %s: %v\n", path, err)
		os.Exit(1)
	}

	content := fixMetadata(string(raw))

	schemaPath := filepath.Join(filepath.Dir(path), "schema.json")
	schemaKeys, err := loadSchemaResourceKeys(schemaPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "metadatafix: read schema %s: %v\n", schemaPath, err)
		os.Exit(1)
	}
	content = injectMissingResourceStubs(content, schemaKeys)

	if content == string(raw) {
		// Idempotent no-op (e.g. upstream already fixed the docs).
		return
	}

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil { // #nosec G703
		fmt.Fprintf(os.Stderr, "metadatafix: write %s: %v\n", path, err)
		os.Exit(1)
	}
}

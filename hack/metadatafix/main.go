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
// The corrections are exact line/string replacements so the scraper's output
// formatting is preserved verbatim (no YAML round-trip churn) and the step is
// idempotent: re-running on already-corrected content is a no-op.
//
// This runs as part of `make generate`, immediately after the scraper writes
// the file, so the fix is durable and survives regeneration (it does not rely
// on a hand edit that check-diff would revert).
package main

import (
	"fmt"
	"os"
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
	content := string(raw)

	// Line-scoped key re-keying.
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

	if content == string(raw) {
		// Idempotent no-op (e.g. upstream already fixed the docs).
		return
	}

	if err := os.WriteFile(path, []byte(content), 0o600); err != nil { // #nosec G703
		fmt.Fprintf(os.Stderr, "metadatafix: write %s: %v\n", path, err)
		os.Exit(1)
	}
}

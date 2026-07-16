package config

import (
	ujconfig "github.com/crossplane/upjet/v2/pkg/config"
)

// configureGridscale applies resource-specific overrides where upjet's automatic
// group/kind derivation produces invalid output.
func configureGridscale(p *ujconfig.Provider) {
	// gridscale_mysql8_0: the trailing "_0" version segment makes upjet derive an
	// invalid Go kind ("0"). Give it an explicit, valid Kind.
	p.AddResourceConfigurator(tfMySQL8, func(r *ujconfig.Resource) {
		r.Kind = "MySQL8"
	})
	// gridscale_storage_import: kind "Import" -> controller package "import", a Go
	// reserved word. Use a non-keyword Kind (package becomes "storageimport").
	p.AddResourceConfigurator("gridscale_storage_import", func(r *ujconfig.Resource) {
		r.Kind = "StorageImport"
	})
}

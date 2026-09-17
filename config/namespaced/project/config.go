package project

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_project resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_project", func(r *config.Resource) {
		r.Kind = "Project"
		r.ShortGroup = "project"
	})
}

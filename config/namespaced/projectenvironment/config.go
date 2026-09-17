package projectenvironment

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_project_environment resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_project_environment", func(r *config.Resource) {
		r.Kind = "ProjectEnvironment"
		r.ShortGroup = "project"
		// project_id is the owning Project's real API ID - let a ProjectEnvironment
		// XR reference a Project XR by name instead of hardcoding its ID.
		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
	})
}

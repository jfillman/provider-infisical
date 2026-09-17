package projectidentity

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_project_identity resource.
//
// This is the resource that directly replaces the current hand-rolled
// operator's ensure_project_membership() call - a real, previously-confirmed
// non-idempotent bug (a live 400 on retry, the first time that code path ever
// ran against real data). The upstream resource's own adopt_existing option
// ("if the identity is already a member of the project... the existing
// membership is adopted and its roles are updated to match the desired state
// instead of returning an error") is exactly the idempotent upsert behavior
// that bug needed - available for free here, not something this config has
// to implement.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_project_identity", func(r *config.Resource) {
		r.Kind = "ProjectIdentity"
		r.ShortGroup = "project"
		r.References["project_id"] = config.Reference{
			TerraformName: "infisical_project",
		}
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
	})
}

package identity

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_identity resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity", func(r *config.Resource) {
		r.Kind = "Identity"
		r.ShortGroup = "identity"
		// org_id is a plain string the XR author supplies directly (same
		// INFISICAL_ORG_ID-shaped value the current hand-rolled operator
		// already requires) - it doesn't reference another XR in this
		// catalog, so no config.Reference here.
	})
}

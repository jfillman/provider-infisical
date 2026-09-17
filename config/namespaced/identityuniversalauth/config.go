package identityuniversalauth

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_identity_universal_auth resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity_universal_auth", func(r *config.Resource) {
		r.Kind = "IdentityUniversalAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
	})
}

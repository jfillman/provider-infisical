package identitykubernetesauth

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_identity_kubernetes_auth resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity_kubernetes_auth", func(r *config.Resource) {
		r.Kind = "IdentityKubernetesAuth"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
	})
}

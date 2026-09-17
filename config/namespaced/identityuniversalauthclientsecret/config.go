package identityuniversalauthclientsecret

import "github.com/crossplane/upjet/v2/pkg/config"

// Configure infisical_identity_universal_auth_client_secret resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity_universal_auth_client_secret", func(r *config.Resource) {
		r.Kind = "IdentityUniversalAuthClientSecret"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}
		// client_id/client_secret are Read-Only+Sensitive in the upstream
		// schema - upjet publishes sensitive computed fields via the managed
		// resource's native connection secret (writeConnectionSecretToRef)
		// automatically, replacing the hand-rolled write_credentials_secret
		// step the current operator does itself.
	})
}

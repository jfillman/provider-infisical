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
		// token_reviewer_jwt is a long-lived ServiceAccount token Infisical uses to call
		// TokenReview on the cluster. The upstream Terraform schema does not mark it
		// Sensitive, so upjet generated a plain spec.forProvider.tokenReviewerJwt string
		// and every IdentityKubernetesAuth carried the token in etcd, in every
		// composition round trip and in function payloads (airframe review C2). Marking it
		// sensitive makes upjet generate spec.forProvider.tokenReviewerJwtSecretRef (a
		// LocalSecretKeySelector: a Secret in the resource's own namespace) instead, and
		// the provider reads the value only when it talks to Terraform.
		r.TerraformResource.Schema["token_reviewer_jwt"].Sensitive = true
	})
}

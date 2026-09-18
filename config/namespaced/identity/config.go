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

		// Publishes the provider-assigned identity id via
		// writeConnectionSecretToRef, under key "identityId" - lets any
		// consumer (a plain Helm-rendered manifest, not just a Crossplane
		// Composition with access to .observed.resources) get this identity's
		// id into a real k8s Secret natively, without a bespoke
		// external-name-annotation read-back step. Default upjet behavior
		// publishes an EMPTY connection secret for this resource (confirmed
		// live) since "id" isn't itself a sensitive Terraform attribute.
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			id, ok := attr["id"].(string)
			if !ok || id == "" {
				return nil, nil
			}
			return map[string][]byte{"identityId": []byte(id)}, nil
		}
	})
}

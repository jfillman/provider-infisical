package identityuniversalauthclientsecret

import (
	"context"

	"github.com/crossplane/upjet/v2/pkg/config"
)

// nilUUID seeds the Terraform state's `id` before the client secret exists.
// A well-formed UUID that can never match a real client secret.
const nilUUID = "00000000-0000-0000-0000-000000000000"

// Configure infisical_identity_universal_auth_client_secret resource.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("infisical_identity_universal_auth_client_secret", func(r *config.Resource) {
		r.Kind = "IdentityUniversalAuthClientSecret"
		r.ShortGroup = "identity"
		r.References["identity_id"] = config.Reference{
			TerraformName: "infisical_identity",
		}

		// Before Create the external name is empty (IdentifierFromProvider), and
		// upjet still seeds the Terraform state with id = external name and runs
		// a refresh. This resource's Read then calls
		//   GET /api/v1/auth/universal-auth/identities/{id}/client-secrets/{clientSecretId}
		// with an EMPTY clientSecretId, which is really the LIST endpoint - it
		// returns an array, and the Terraform provider fails unmarshalling it into
		// a single object ("cannot unmarshal array into Go struct field
		// GetIdentityUniversalAuthClientSecretResponse.clientSecretData"). Observe
		// therefore errors before Create can ever run.
		//
		// Seeding a well-formed id that cannot exist makes the same Read hit the
		// by-id endpoint, get a 404, and report "not found" - the normal path for
		// a resource that has not been created yet. Only affects the seeded state;
		// once Create returns, the real id replaces it as the external name.
		r.ExternalName.GetIDFn = func(_ context.Context, externalName string, _ map[string]any, _ map[string]any) (string, error) {
			if externalName == "" {
				return nilUUID, nil
			}
			return externalName, nil
		}

		// client_id/client_secret are Read-Only+Sensitive in the upstream
		// schema - upjet publishes sensitive computed fields via the managed
		// resource's native connection secret (writeConnectionSecretToRef)
		// automatically, replacing the hand-rolled write_credentials_secret
		// step the current operator does itself.
	})
}

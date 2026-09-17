package config

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
//
// All seven MVP resources are confirmed (via each resource's own "## Import"
// docs section, or its Read-Only `id` schema attribute where no Import
// section is published) to use a provider-assigned ID, never a user-supplied
// or composite one - so every entry uses the same identifierFromProvider()
// helper, matching the identical shape the template's own null_resource
// example already used.
var ExternalNameConfigs = map[string]config.ExternalName{
	// Imports by the project's own ID (docs/resources/project.md).
	"infisical_project": identifierFromProvider(),
	// Imports by the project-environment's own ID (docs/resources/project_environment.md).
	"infisical_project_environment": identifierFromProvider(),
	// Imports by the identity's own ID (docs/resources/identity.md).
	"infisical_identity": identifierFromProvider(),
	// No published Import section, but `id` is a Read-Only, provider-assigned
	// attribute (docs/resources/identity_universal_auth.md) - one per identity.
	"infisical_identity_universal_auth": identifierFromProvider(),
	// Same as above - `id` (and `client_id`) are Read-Only/provider-assigned
	// (docs/resources/identity_universal_auth_client_secret.md).
	"infisical_identity_universal_auth_client_secret": identifierFromProvider(),
	// Imports by the identity's own ID (docs/resources/identity_kubernetes_auth.md).
	"infisical_identity_kubernetes_auth": identifierFromProvider(),
	// `id` is a Read-Only "ID of the project identity role" attribute
	// (docs/resources/project_identity.md) - no separate Import section
	// published, treated the same conservative way as the two universal-auth
	// resources above.
	"infisical_project_identity": identifierFromProvider(),
}

// identifierFromProvider is the shared strategy for every MVP resource: the
// real external name is exactly the API's own `id` field, nothing derived or
// templated. Named distinctly from the template's original idWithStub() only
// to read clearly now that every resource in this table uses it, not just one
// sample.
func identifierFromProvider() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		en, _ := config.IDAsExternalName(tfstate)
		return en, nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}

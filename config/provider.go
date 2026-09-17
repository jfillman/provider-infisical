package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	"github.com/jfillman/provider-infisical/config/namespaced/identity"
	"github.com/jfillman/provider-infisical/config/namespaced/identitykubernetesauth"
	"github.com/jfillman/provider-infisical/config/namespaced/identityuniversalauth"
	"github.com/jfillman/provider-infisical/config/namespaced/identityuniversalauthclientsecret"
	"github.com/jfillman/provider-infisical/config/namespaced/project"
	"github.com/jfillman/provider-infisical/config/namespaced/projectenvironment"
	"github.com/jfillman/provider-infisical/config/namespaced/projectidentity"
)

const (
	resourcePrefix = "infisical"
	modulePath     = "github.com/jfillman/provider-infisical"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("infisical.hangar.io"),
		// Deliberately empty, not ExternalNameConfigured() - an unset/non-empty
		// IncludeList defaults to including every resource in the Terraform
		// provider's schema (upjet's own doc: "Defaults to []string{".+"}
		// which would include all resources"). Every MVP resource is
		// namespaced-only (matching how every other XRD in this catalog is
		// already namespaced-only), so the cluster-scoped provider half
		// should generate nothing rather than silently doubling the CRD
		// surface with un-configured, reference-less cluster variants.
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	// No cluster-scoped resources in the MVP set - every Infisical resource
	// we configure (project/identity/etc.) is namespaced, see
	// GetProviderNamespaced below.

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("infisical.m.hangar.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// MVP resource set - mirrors what the hand-rolled
		// infisical-secretstore-operator does today, nothing more.
		project.Configure,
		projectenvironment.Configure,
		identity.Configure,
		identityuniversalauth.Configure,
		identityuniversalauthclientsecret.Configure,
		identitykubernetesauth.Configure,
		projectidentity.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	identity "github.com/jfillman/provider-infisical/internal/controller/namespaced/identity/identity"
	identitykubernetesauth "github.com/jfillman/provider-infisical/internal/controller/namespaced/identity/identitykubernetesauth"
	identityuniversalauth "github.com/jfillman/provider-infisical/internal/controller/namespaced/identity/identityuniversalauth"
	identityuniversalauthclientsecret "github.com/jfillman/provider-infisical/internal/controller/namespaced/identity/identityuniversalauthclientsecret"
	project "github.com/jfillman/provider-infisical/internal/controller/namespaced/project/project"
	projectenvironment "github.com/jfillman/provider-infisical/internal/controller/namespaced/project/projectenvironment"
	projectidentity "github.com/jfillman/provider-infisical/internal/controller/namespaced/project/projectidentity"
	providerconfig "github.com/jfillman/provider-infisical/internal/controller/namespaced/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identity.Setup,
		identitykubernetesauth.Setup,
		identityuniversalauth.Setup,
		identityuniversalauthclientsecret.Setup,
		project.Setup,
		projectenvironment.Setup,
		projectidentity.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		identity.SetupGated,
		identitykubernetesauth.SetupGated,
		identityuniversalauth.SetupGated,
		identityuniversalauthclientsecret.SetupGated,
		project.SetupGated,
		projectenvironment.SetupGated,
		projectidentity.SetupGated,
		providerconfig.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupWebhookWithManager registers conversion webhooks for all resource kinds in the group.
func SetupWebhookWithManager(mgr ctrl.Manager) error {
	for _, setup := range []func(ctrl.Manager) error{
		identity.SetupWebhookWithManager,
		identitykubernetesauth.SetupWebhookWithManager,
		identityuniversalauth.SetupWebhookWithManager,
		identityuniversalauthclientsecret.SetupWebhookWithManager,
		project.SetupWebhookWithManager,
		projectenvironment.SetupWebhookWithManager,
		projectidentity.SetupWebhookWithManager,
		providerconfig.SetupWebhookWithManager,
	} {
		if err := setup(mgr); err != nil {
			return err
		}
	}
	return nil
}

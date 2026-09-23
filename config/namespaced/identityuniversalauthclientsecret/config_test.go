package identityuniversalauthclientsecret_test

import (
	"context"
	"testing"

	"github.com/crossplane/upjet/v2/pkg/config"

	ic "github.com/jfillman/provider-infisical/config"
)

func TestSeededIDNeverEmpty(t *testing.T) {
	p := ic.GetProviderNamespaced()
	var r *config.Resource
	for name, res := range p.Resources {
		if name == "infisical_identity_universal_auth_client_secret" {
			r = res
		}
	}
	if r == nil {
		t.Fatal("resource not registered")
	}
	id, _ := r.ExternalName.GetIDFn(context.Background(), "", nil, nil)
	if id == "" {
		t.Fatal("seeded id is empty: Read would hit the LIST endpoint")
	}
	id, _ = r.ExternalName.GetIDFn(context.Background(), "real-id", nil, nil)
	if id != "real-id" {
		t.Fatalf("a real external name must pass through, got %q", id)
	}
}

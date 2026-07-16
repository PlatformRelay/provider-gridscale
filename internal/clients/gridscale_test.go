package clients

import (
	"context"
	"reflect"
	"strings"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	"github.com/crossplane/upjet/v2/pkg/terraform"

	namespacedv1beta1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/v1beta1"
)

func TestBuildProviderConfiguration(t *testing.T) {
	const (
		testUUID   = "user-uuid"
		testToken  = "api-token"
		testAPIURL = "https://api.gridscale.io"
	)

	cases := map[string]struct {
		creds map[string]string
		want  map[string]any
	}{
		"UUIDAndTokenMapped": {
			creds: map[string]string{
				keyUUID:  testUUID,
				keyToken: testToken,
			},
			want: map[string]any{
				keyUUID:  testUUID,
				keyToken: testToken,
			},
		},
		"APIURLPassedThroughWhenPresent": {
			creds: map[string]string{
				keyUUID:   testUUID,
				keyToken:  testToken,
				keyAPIURL: testAPIURL,
			},
			want: map[string]any{
				keyUUID:   testUUID,
				keyToken:  testToken,
				keyAPIURL: testAPIURL,
			},
		},
		"APIURLOmittedWhenAbsent": {
			creds: map[string]string{
				keyUUID:  testUUID,
				keyToken: testToken,
			},
			want: map[string]any{
				keyUUID:  testUUID,
				keyToken: testToken,
			},
		},
		"APIURLOmittedWhenPresentButEmpty": {
			creds: map[string]string{
				keyUUID:   testUUID,
				keyToken:  testToken,
				keyAPIURL: "",
			},
			want: map[string]any{
				keyUUID:  testUUID,
				keyToken: testToken,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got, err := buildProviderConfiguration(tc.creds)
			if err != nil {
				t.Fatalf("buildProviderConfiguration(%v): unexpected error: %v", tc.creds, err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("buildProviderConfiguration(%v) = %v, want %v", tc.creds, got, tc.want)
			}
		})
	}
}

// --- E2-S11: non-live credential-wiring regression tests ---------------------
//
// These guard against the CRED-1/CRED-2 audit-finding shape: credentials never
// reaching the Terraform provider, or reaching it under the wrong keys (e.g.
// username/password instead of uuid/token). They run the full
// TerraformSetupBuilder against a controller-runtime fake client — no network,
// no live API.

// setupWithSecretBlob seeds a namespaced ProviderConfig + credential Secret
// holding the given JSON blob and runs TerraformSetupBuilder for a modern MR.
func setupWithSecretBlob(t *testing.T, blob string) (terraform.Setup, error) {
	t.Helper()
	s := buildScheme(t)
	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: testNamespace},
		Spec:       namespacedv1beta1.ProviderConfigSpec{Credentials: secretCredentials(testNamespace)},
	}
	secret := credentialSecret(testNamespace, blob)
	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc, secret).Build()
	return TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newModernMR())
}

func TestTerraformSetupConfigurationExactShape(t *testing.T) {
	cases := map[string]struct {
		blob string
		want map[string]any
	}{
		"UUIDTokenAndAPIURL": {
			blob: `{"uuid":"` + credUUID + `","token":"` + credToken + `","api_url":"` + credAPIURL + `"}`,
			want: map[string]any{
				keyUUID:   credUUID,
				keyToken:  credToken,
				keyAPIURL: credAPIURL,
			},
		},
		"UUIDAndTokenOnly": {
			blob: `{"uuid":"` + credUUID + `","token":"` + credToken + `"}`,
			want: map[string]any{
				keyUUID:  credUUID,
				keyToken: credToken,
			},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			setup, err := setupWithSecretBlob(t, tc.blob)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			// Exact-shape assertion: exactly the gridscale provider keys, no
			// extras, no renames.
			if !reflect.DeepEqual(map[string]any(setup.Configuration), tc.want) {
				t.Errorf("Configuration = %v, want exactly %v", setup.Configuration, tc.want)
			}
		})
	}
}

// TestTerraformSetupConfigurationNotCRED1CRED2Shape is the explicit regression
// guard for audit findings CRED-1 (credentials never passed to Terraform, i.e.
// empty Configuration) and CRED-2 (credentials passed under wrong auth keys,
// e.g. username/password).
func TestTerraformSetupConfigurationNotCRED1CRED2Shape(t *testing.T) {
	setup, err := setupWithSecretBlob(t,
		`{"uuid":"`+credUUID+`","token":"`+credToken+`","api_url":"`+credAPIURL+`"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// CRED-1 shape: nothing wired through at all.
	if len(setup.Configuration) == 0 {
		t.Fatal("Configuration is empty: credentials were not passed to the Terraform provider (CRED-1 regression)")
	}
	// CRED-2 shape: wrong auth keys for the gridscale provider schema.
	for _, wrongKey := range []string{"username", "password"} {
		if _, ok := setup.Configuration[wrongKey]; ok {
			t.Errorf("Configuration contains %q: wrong auth key for the gridscale provider (CRED-2 regression)", wrongKey)
		}
	}
	for _, requiredKey := range []string{keyUUID, keyToken} {
		if v, ok := setup.Configuration[requiredKey]; !ok || v == "" {
			t.Errorf("Configuration missing required gridscale auth key %q (got %v)", requiredKey, v)
		}
	}
}

// TestTerraformSetupBuilderMissingRequiredKeys asserts that a secret which
// parses as JSON but lacks the required gridscale keys fails fast with a
// wrapped, descriptive error instead of handing empty credentials to
// Terraform (where auth would fail later with a cryptic provider error).
func TestTerraformSetupBuilderMissingRequiredKeys(t *testing.T) {
	cases := map[string]struct {
		blob        string
		wantMissing []string
	}{
		"MissingUUID": {
			blob:        `{"token":"` + credToken + `"}`,
			wantMissing: []string{keyUUID},
		},
		"MissingToken": {
			blob:        `{"uuid":"` + credUUID + `"}`,
			wantMissing: []string{keyToken},
		},
		"EmptyValues": {
			blob:        `{"uuid":"","token":""}`,
			wantMissing: []string{keyUUID, keyToken},
		},
		"EmptyObject": {
			blob:        `{}`,
			wantMissing: []string{keyUUID, keyToken},
		},
		"WrongAuthKeysOnly": {
			// The CRED-2 shape on the input side: a secret authored with the
			// wrong key names must be rejected, not silently mapped to empty
			// uuid/token.
			blob:        `{"username":"` + credUUID + `","password":"` + credToken + `"}`,
			wantMissing: []string{keyUUID, keyToken},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err := setupWithSecretBlob(t, tc.blob)
			if err == nil {
				t.Fatalf("expected a descriptive error for secret %s, got nil", tc.blob)
			}
			if !strings.Contains(err.Error(), "missing required") {
				t.Errorf("error = %q, want a descriptive %q error", err.Error(), "missing required")
			}
			for _, key := range tc.wantMissing {
				if !strings.Contains(err.Error(), key) {
					t.Errorf("error = %q, want it to name the missing key %q", err.Error(), key)
				}
			}
		})
	}
}

package clients

import (
	"context"
	"strings"
	"testing"

	xpv1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/v2/pkg/resource"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	clusterv1beta1 "github.com/PlatformRelay/provider-gridscale/apis/cluster/v1beta1"
	namespacedv1beta1 "github.com/PlatformRelay/provider-gridscale/apis/namespaced/v1beta1"
)

// These tests exercise the credential-resolution surface of gridscale.go end to
// end against a controller-runtime fake client: TerraformSetupBuilder,
// resolveProviderConfig, resolveModern and resolveLegacy. They are hermetic (no
// network, no cluster) and assert both the happy path (a Setup whose
// Configuration carries the mapped uuid/token/api_url) and the audit-critical
// error paths (missing providerConfigRef, missing secret, malformed JSON).

const (
	credUUID   = "11111111-2222-3333-4444-555555555555"
	credToken  = "super-secret-token"
	credAPIURL = "https://api.gridscale.io"

	testNamespace = "team-a"
	pcName        = "gridscale-pc"
	secretName    = "gridscale-creds"
	secretKey     = "credentials"
)

// buildScheme registers everything the resolution paths touch: the namespaced +
// cluster ProviderConfig/PCU types (resolveModern calls Scheme().New and both
// trackers Apply a PCU) and corev1 (CommonCredentialExtractor Gets a Secret).
func buildScheme(t *testing.T) *runtime.Scheme {
	t.Helper()
	s := runtime.NewScheme()
	if err := clusterv1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		t.Fatalf("add cluster scheme: %v", err)
	}
	if err := namespacedv1beta1.SchemeBuilder.AddToScheme(s); err != nil {
		t.Fatalf("add namespaced scheme: %v", err)
	}
	if err := corev1.AddToScheme(s); err != nil {
		t.Fatalf("add corev1 scheme: %v", err)
	}
	return s
}

// credentialSecret builds a Secret whose single key holds the JSON credential
// blob the gridscale provider expects (uuid/token/api_url).
func credentialSecret(namespace, blob string) *corev1.Secret {
	return &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: secretName, Namespace: namespace},
		Data:       map[string][]byte{secretKey: []byte(blob)},
	}
}

func secretCredentials(namespace string) namespacedv1beta1.ProviderCredentials {
	return namespacedv1beta1.ProviderCredentials{
		Source: xpv1.CredentialsSourceSecret,
		CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
			SecretRef: &xpv1.SecretKeySelector{
				SecretReference: xpv1.SecretReference{Name: secretName, Namespace: namespace},
				Key:             secretKey,
			},
		},
	}
}

// --- Fake managed resources ------------------------------------------------
//
// Minimal hand-rolled fakes that satisfy exactly the crossplane-runtime
// interfaces resolveProviderConfig dispatches on. They are never round-tripped
// through the client (only their getters are read), so they need not be
// registered in the scheme.

// fakeModernMR implements resource.ModernManaged.
type fakeModernMR struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	pcRef *xpv1.ProviderConfigReference
}

func (m *fakeModernMR) GetObjectKind() schema.ObjectKind { return &m.TypeMeta }
func (m *fakeModernMR) DeepCopyObject() runtime.Object   { return m }

func (m *fakeModernMR) GetProviderConfigReference() *xpv1.ProviderConfigReference  { return m.pcRef }
func (m *fakeModernMR) SetProviderConfigReference(p *xpv1.ProviderConfigReference) { m.pcRef = p }

func (m *fakeModernMR) GetManagementPolicies() xpv1.ManagementPolicies   { return nil }
func (m *fakeModernMR) SetManagementPolicies(_ xpv1.ManagementPolicies)  {}
func (m *fakeModernMR) SetConditions(_ ...xpv1.Condition)                {}
func (m *fakeModernMR) GetCondition(_ xpv1.ConditionType) xpv1.Condition { return xpv1.Condition{} }
func (m *fakeModernMR) GetWriteConnectionSecretToReference() *xpv1.LocalSecretReference {
	return nil
}
func (m *fakeModernMR) SetWriteConnectionSecretToReference(_ *xpv1.LocalSecretReference) {}

// fakeLegacyMR implements resource.LegacyManaged (cluster-scoped, untyped ref).
type fakeLegacyMR struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	pcRef *xpv1.Reference
}

func (m *fakeLegacyMR) GetObjectKind() schema.ObjectKind { return &m.TypeMeta }
func (m *fakeLegacyMR) DeepCopyObject() runtime.Object   { return m }

func (m *fakeLegacyMR) GetProviderConfigReference() *xpv1.Reference  { return m.pcRef }
func (m *fakeLegacyMR) SetProviderConfigReference(p *xpv1.Reference) { m.pcRef = p }

func (m *fakeLegacyMR) GetManagementPolicies() xpv1.ManagementPolicies   { return nil }
func (m *fakeLegacyMR) SetManagementPolicies(_ xpv1.ManagementPolicies)  {}
func (m *fakeLegacyMR) SetConditions(_ ...xpv1.Condition)                {}
func (m *fakeLegacyMR) GetCondition(_ xpv1.ConditionType) xpv1.Condition { return xpv1.Condition{} }
func (m *fakeLegacyMR) GetDeletionPolicy() xpv1.DeletionPolicy           { return "" }
func (m *fakeLegacyMR) SetDeletionPolicy(_ xpv1.DeletionPolicy)          {}
func (m *fakeLegacyMR) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return nil
}
func (m *fakeLegacyMR) SetWriteConnectionSecretToReference(_ *xpv1.SecretReference) {}

// compile-time assertions that our fakes satisfy the dispatched interfaces.
var (
	_ resource.ModernManaged = &fakeModernMR{}
	_ resource.LegacyManaged = &fakeLegacyMR{} //nolint:staticcheck // exercising the legacy path on purpose
)

func newModernMR() *fakeModernMR {
	return &fakeModernMR{
		ObjectMeta: metav1.ObjectMeta{Name: "modern-mr", Namespace: testNamespace, UID: "modern-uid"},
		pcRef:      &xpv1.ProviderConfigReference{Name: pcName, Kind: namespacedv1beta1.ProviderConfigKind},
	}
}

func newLegacyMR() *fakeLegacyMR {
	return &fakeLegacyMR{
		ObjectMeta: metav1.ObjectMeta{Name: "legacy-mr", UID: "legacy-uid"},
		pcRef:      &xpv1.Reference{Name: pcName},
	}
}

// --- TerraformSetupBuilder (modern happy path) -----------------------------

func TestTerraformSetupBuilderModernMapsCredentials(t *testing.T) {
	s := buildScheme(t)

	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: testNamespace},
		Spec:       namespacedv1beta1.ProviderConfigSpec{Credentials: secretCredentials(testNamespace)},
	}
	secret := credentialSecret(testNamespace,
		`{"uuid":"`+credUUID+`","token":"`+credToken+`","api_url":"`+credAPIURL+`"}`)

	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc, secret).Build()

	setup, err := TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newModernMR())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := setup.Configuration[keyUUID]; got != credUUID {
		t.Errorf("uuid = %v, want %v", got, credUUID)
	}
	if got := setup.Configuration[keyToken]; got != credToken {
		t.Errorf("token = %v, want %v", got, credToken)
	}
	if got := setup.Configuration[keyAPIURL]; got != credAPIURL {
		t.Errorf("api_url = %v, want %v", got, credAPIURL)
	}
	if setup.Version != "1.5.7" || setup.Requirement.Source != "gridscale/gridscale" || setup.Requirement.Version != "2.3.0" {
		t.Errorf("provider requirement not propagated: %+v", setup)
	}
}

func TestTerraformSetupBuilderOmitsEmptyAPIURL(t *testing.T) {
	s := buildScheme(t)

	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: testNamespace},
		Spec:       namespacedv1beta1.ProviderConfigSpec{Credentials: secretCredentials(testNamespace)},
	}
	// No api_url key at all: buildProviderConfiguration must not inject one.
	secret := credentialSecret(testNamespace, `{"uuid":"`+credUUID+`","token":"`+credToken+`"}`)

	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc, secret).Build()

	setup, err := TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newModernMR())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := setup.Configuration[keyAPIURL]; ok {
		t.Errorf("api_url should be absent, got %v", setup.Configuration[keyAPIURL])
	}
}

// --- TerraformSetupBuilder (legacy happy path) -----------------------------

func TestTerraformSetupBuilderLegacyMapsCredentials(t *testing.T) {
	s := buildScheme(t)

	// Legacy PC is cluster-scoped and resolveLegacy does NOT rewrite the secret
	// namespace, so the SecretRef namespace and the Secret must be explicit and
	// aligned.
	pc := &clusterv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName},
		Spec:       clusterv1beta1.ProviderConfigSpec{Credentials: legacyCredentials()},
	}
	secret := credentialSecret(testNamespace,
		`{"uuid":"`+credUUID+`","token":"`+credToken+`"}`)

	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc, secret).Build()

	setup, err := TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newLegacyMR())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got := setup.Configuration[keyUUID]; got != credUUID {
		t.Errorf("uuid = %v, want %v", got, credUUID)
	}
	if got := setup.Configuration[keyToken]; got != credToken {
		t.Errorf("token = %v, want %v", got, credToken)
	}
}

func legacyCredentials() clusterv1beta1.ProviderCredentials {
	return clusterv1beta1.ProviderCredentials{
		Source: xpv1.CredentialsSourceSecret,
		CommonCredentialSelectors: xpv1.CommonCredentialSelectors{
			SecretRef: &xpv1.SecretKeySelector{
				SecretReference: xpv1.SecretReference{Name: secretName, Namespace: testNamespace},
				Key:             secretKey,
			},
		},
	}
}

// --- Error paths -----------------------------------------------------------

func TestTerraformSetupBuilderBadJSON(t *testing.T) {
	s := buildScheme(t)

	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: testNamespace},
		Spec:       namespacedv1beta1.ProviderConfigSpec{Credentials: secretCredentials(testNamespace)},
	}
	secret := credentialSecret(testNamespace, `this-is-not-json`)

	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc, secret).Build()

	_, err := TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newModernMR())
	if err == nil {
		t.Fatal("expected error for malformed credential JSON, got nil")
	}
	if !strings.Contains(err.Error(), errUnmarshalCredentials) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errUnmarshalCredentials)
	}
}

func TestTerraformSetupBuilderMissingSecret(t *testing.T) {
	s := buildScheme(t)

	pc := &namespacedv1beta1.ProviderConfig{
		ObjectMeta: metav1.ObjectMeta{Name: pcName, Namespace: testNamespace},
		Spec:       namespacedv1beta1.ProviderConfigSpec{Credentials: secretCredentials(testNamespace)},
	}
	// PC present, Secret absent.
	cl := fake.NewClientBuilder().WithScheme(s).WithObjects(pc).Build()

	_, err := TerraformSetupBuilder("1.5.7", "gridscale/gridscale", "2.3.0")(
		context.Background(), cl, newModernMR())
	if err == nil {
		t.Fatal("expected error when secret is missing, got nil")
	}
	if !strings.Contains(err.Error(), errExtractCredentials) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errExtractCredentials)
	}
}

func TestResolveModernNoProviderConfigRef(t *testing.T) {
	s := buildScheme(t)
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	mr := newModernMR()
	mr.pcRef = nil

	_, err := resolveProviderConfig(context.Background(), cl, mr)
	if err == nil {
		t.Fatal("expected error when providerConfigRef is nil, got nil")
	}
	if !strings.Contains(err.Error(), errNoProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errNoProviderConfig)
	}
}

func TestResolveLegacyNoProviderConfigRef(t *testing.T) {
	s := buildScheme(t)
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	mr := newLegacyMR()
	mr.pcRef = nil

	_, err := resolveProviderConfig(context.Background(), cl, mr)
	if err == nil {
		t.Fatal("expected error when providerConfigRef is nil, got nil")
	}
	if !strings.Contains(err.Error(), errNoProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errNoProviderConfig)
	}
}

func TestResolveModernMissingProviderConfig(t *testing.T) {
	s := buildScheme(t)
	// No PC seeded at all.
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	_, err := resolveProviderConfig(context.Background(), cl, newModernMR())
	if err == nil {
		t.Fatal("expected error when ProviderConfig is missing, got nil")
	}
	if !strings.Contains(err.Error(), errGetProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errGetProviderConfig)
	}
}

func TestResolveLegacyMissingProviderConfig(t *testing.T) {
	s := buildScheme(t)
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	_, err := resolveProviderConfig(context.Background(), cl, newLegacyMR())
	if err == nil {
		t.Fatal("expected error when ProviderConfig is missing, got nil")
	}
	if !strings.Contains(err.Error(), errGetProviderConfig) {
		t.Errorf("error = %q, want it to contain %q", err.Error(), errGetProviderConfig)
	}
}

func TestResolveProviderConfigNotManaged(t *testing.T) {
	s := buildScheme(t)
	cl := fake.NewClientBuilder().WithScheme(s).Build()

	// A Managed resource that is neither Legacy nor Modern hits the default branch.
	_, err := resolveProviderConfig(context.Background(), cl, &notManaged{})
	if err == nil {
		t.Fatal("expected error for a non-managed resource, got nil")
	}
}

// notManaged is a resource.Managed-shaped stand-in that is NOT recognised as a
// LegacyManaged or ModernManaged, exercising the default branch of
// resolveProviderConfig. It only needs to satisfy resource.Managed.
type notManaged struct {
	metav1.TypeMeta
	metav1.ObjectMeta
}

func (m *notManaged) GetObjectKind() schema.ObjectKind               { return &m.TypeMeta }
func (m *notManaged) DeepCopyObject() runtime.Object                 { return m }
func (*notManaged) GetManagementPolicies() xpv1.ManagementPolicies   { return nil }
func (*notManaged) SetManagementPolicies(_ xpv1.ManagementPolicies)  {}
func (*notManaged) SetConditions(_ ...xpv1.Condition)                {}
func (*notManaged) GetCondition(_ xpv1.ConditionType) xpv1.Condition { return xpv1.Condition{} }

var _ resource.Managed = &notManaged{}

// --- toSharedPCSpec --------------------------------------------------------

func TestToSharedPCSpecNil(t *testing.T) {
	got, err := toSharedPCSpec(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != nil {
		t.Errorf("expected nil spec for nil PC, got %+v", got)
	}
}

func TestToSharedPCSpecRoundTrips(t *testing.T) {
	pc := &clusterv1beta1.ProviderConfig{
		Spec: clusterv1beta1.ProviderConfigSpec{Credentials: legacyCredentials()},
	}
	got, err := toSharedPCSpec(pc)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("expected non-nil shared spec")
	}
	if got.Credentials.Source != xpv1.CredentialsSourceSecret {
		t.Errorf("source = %v, want %v", got.Credentials.Source, xpv1.CredentialsSourceSecret)
	}
	if got.Credentials.SecretRef == nil || got.Credentials.SecretRef.Name != secretName {
		t.Errorf("secretRef not carried across: %+v", got.Credentials.SecretRef)
	}
}

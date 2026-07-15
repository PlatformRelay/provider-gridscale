package clients

import (
	"reflect"
	"testing"
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
			got := buildProviderConfiguration(tc.creds)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("buildProviderConfiguration(%v) = %v, want %v", tc.creds, got, tc.want)
			}
		})
	}
}

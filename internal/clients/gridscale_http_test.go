// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package clients_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PlatformRelay/provider-gridscale/internal/clients"
)

func TestGridscaleClientGet_Success(t *testing.T) {
	type payload struct {
		Foo string `json:"foo"`
	}
	want := payload{Foo: "bar"}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Auth-UserID") == "" {
			t.Error("missing X-Auth-UserID header")
		}
		if r.Header.Get("X-Auth-Token") == "" {
			t.Error("missing X-Auth-Token header")
		}
		if r.URL.Path != "/objects/test" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(want) //nolint:errcheck
	}))
	defer srv.Close()

	c := clients.NewGridscaleClient("test-uuid", "test-token", srv.URL)
	var got payload
	if err := c.Get(context.Background(), "/objects/test", &got); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Foo != want.Foo {
		t.Errorf("got %q, want %q", got.Foo, want.Foo)
	}
}

func TestGridscaleClientGet_HTTPError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "forbidden", http.StatusForbidden)
	}))
	defer srv.Close()

	c := clients.NewGridscaleClient("u", "t", srv.URL)
	var out any
	err := c.Get(context.Background(), "/x", &out)
	if err == nil {
		t.Fatal("expected non-nil error for 403 response")
	}
}

func TestGridscaleClientGet_DefaultBaseURL(t *testing.T) {
	// NewGridscaleClient with empty baseURL should use the default gridscale URL
	// (we cannot call it without a real server, just verify no panic and error wraps)
	c := clients.NewGridscaleClient("u", "t", "")
	_ = c // just verify construction doesn't panic
}

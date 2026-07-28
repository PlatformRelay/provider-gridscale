// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package backuplist

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1alpha1 "github.com/PlatformRelay/provider-gridscale/apis/cluster/storage/v1alpha1"
	"github.com/PlatformRelay/provider-gridscale/internal/clients"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newTestClient(t *testing.T, handler http.HandlerFunc) (*clients.GridscaleClient, func()) {
	t.Helper()
	srv := httptest.NewServer(handler)
	c := clients.NewGridscaleClient("uuid", "token", srv.URL)
	return c, srv.Close
}

func TestObserve_PopulatesBackups(t *testing.T) {
	respBody := backupListResponse{
		StorageBackups: []struct {
			ObjectUUID string  `json:"object_uuid"`
			Name       string  `json:"name"`
			Capacity   float64 `json:"capacity"`
			CreateTime string  `json:"create_time"`
		}{
			{ObjectUUID: "abc", Name: "backup-1", Capacity: 10, CreateTime: "2024-01-01"},
		},
	}
	c, cleanup := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(respBody) //nolint:errcheck
	})
	defer cleanup()

	ext := &external{client: c}
	bl := &v1alpha1.BackupList{
		ObjectMeta: metav1.ObjectMeta{Name: "test"},
		Spec: v1alpha1.BackupListSpec{
			ForProvider: v1alpha1.BackupListParameters{StorageUUID: "storage-uuid"},
		},
	}

	obs, err := ext.Observe(context.Background(), bl)
	if err != nil {
		t.Fatalf("Observe error: %v", err)
	}
	if !obs.ResourceExists {
		t.Error("ResourceExists should be true")
	}
	if len(bl.Status.AtProvider.StorageBackups) != 1 {
		t.Fatalf("expected 1 backup, got %d", len(bl.Status.AtProvider.StorageBackups))
	}
	if bl.Status.AtProvider.StorageBackups[0].ObjectUUID != "abc" {
		t.Errorf("unexpected uuid: %s", bl.Status.AtProvider.StorageBackups[0].ObjectUUID)
	}
}

func TestObserve_EmptyList(t *testing.T) {
	respBody := backupListResponse{}
	c, cleanup := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(respBody) //nolint:errcheck
	})
	defer cleanup()

	ext := &external{client: c}
	bl := &v1alpha1.BackupList{
		ObjectMeta: metav1.ObjectMeta{Name: "test"},
		Spec: v1alpha1.BackupListSpec{
			ForProvider: v1alpha1.BackupListParameters{StorageUUID: "storage-uuid"},
		},
	}

	obs, err := ext.Observe(context.Background(), bl)
	if err != nil {
		t.Fatalf("Observe error: %v", err)
	}
	if !obs.ResourceExists {
		t.Error("ResourceExists should be true even with empty backup list")
	}
	if len(bl.Status.AtProvider.StorageBackups) != 0 {
		t.Errorf("expected 0 backups, got %d", len(bl.Status.AtProvider.StorageBackups))
	}
}

func TestCreate_ReturnsObserveOnlyError(t *testing.T) {
	ext := &external{client: clients.NewGridscaleClient("u", "t", "http://localhost:1")}
	_, err := ext.Create(context.Background(), &v1alpha1.BackupList{})
	if err == nil {
		t.Fatal("Create should return observe-only error")
	}
}

func TestUpdate_ReturnsObserveOnlyError(t *testing.T) {
	ext := &external{client: clients.NewGridscaleClient("u", "t", "http://localhost:1")}
	_, err := ext.Update(context.Background(), &v1alpha1.BackupList{})
	if err == nil {
		t.Fatal("Update should return observe-only error")
	}
}

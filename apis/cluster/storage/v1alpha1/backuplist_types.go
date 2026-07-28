// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// StorageBackupEntry describes a single backup entry returned by the gridscale
// storage backups API.
type StorageBackupEntry struct {
	ObjectUUID string  `json:"objectUuid,omitempty"`
	Name       string  `json:"name,omitempty"`
	Capacity   float64 `json:"capacity,omitempty"`
	CreateTime string  `json:"createTime,omitempty"`
}

// BackupListParameters are the configurable fields of a BackupList.
type BackupListParameters struct {
	// StorageUUID is the UUID of the storage whose backups to list.
	// +kubebuilder:validation:Required
	StorageUUID string `json:"storageUUID"`
}

// BackupListObservation are the observable fields of a BackupList.
type BackupListObservation struct {
	// StorageBackups is the list of backups for the storage.
	StorageBackups []StorageBackupEntry `json:"storageBackups,omitempty"`
}

// BackupListSpec defines the desired state of BackupList.
type BackupListSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BackupListParameters `json:"forProvider"`
}

// BackupListStatus defines the observed state of BackupList.
type BackupListStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BackupListObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gridscale}

// BackupList is the Schema for the BackupLists API. Lists storage backups for a given storage UUID.
type BackupList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BackupListSpec   `json:"spec"`
	Status            BackupListStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupListList contains a list of BackupList.
type BackupListList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupList `json:"items"`
}

// Repository type metadata.
var (
	BackupList_Kind             = "BackupList"
	BackupList_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BackupList_Kind}.String()
	BackupList_KindAPIVersion   = BackupList_Kind + "." + CRDGroupVersion.String()
	BackupList_GroupVersionKind = CRDGroupVersion.WithKind(BackupList_Kind)
)

func init() {
	SchemeBuilder.Register(&BackupList{}, &BackupListList{})
}

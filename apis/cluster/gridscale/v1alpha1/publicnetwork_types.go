// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/v2/apis/common/v1"
)

// PublicNetworkParameters are the configurable fields of a PublicNetwork.
// No inputs are required — the controller discovers the public network automatically.
type PublicNetworkParameters struct{}

// PublicNetworkObservation are the observable fields of a PublicNetwork.
type PublicNetworkObservation struct {
	Name            *string  `json:"name,omitempty"`
	Status          *string  `json:"status,omitempty"`
	NetworkType     *string  `json:"networkType,omitempty"`
	LocationUUID    *string  `json:"locationUUID,omitempty"`
	LocationName    *string  `json:"locationName,omitempty"`
	LocationCountry *string  `json:"locationCountry,omitempty"`
	LocationIATA    *string  `json:"locationIATA,omitempty"`
	L2Security      *bool    `json:"l2Security,omitempty"`
	DeleteBlock     *bool    `json:"deleteBlock,omitempty"`
	Labels          []string `json:"labels,omitempty"`
	CreateTime      *string  `json:"createTime,omitempty"`
	ChangeTime      *string  `json:"changeTime,omitempty"`
}

// PublicNetworkSpec defines the desired state of PublicNetwork.
type PublicNetworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PublicNetworkParameters `json:"forProvider"`
}

// PublicNetworkStatus defines the observed state of PublicNetwork.
type PublicNetworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PublicNetworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gridscale}

// PublicNetwork is the Schema for the PublicNetworks API. Returns the account's public network details.
type PublicNetwork struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicNetworkSpec   `json:"spec"`
	Status            PublicNetworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PublicNetworkList contains a list of PublicNetwork.
type PublicNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PublicNetwork `json:"items"`
}

// Repository type metadata.
var (
	PublicNetwork_Kind             = "PublicNetwork"
	PublicNetwork_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PublicNetwork_Kind}.String()
	PublicNetwork_KindAPIVersion   = PublicNetwork_Kind + "." + CRDGroupVersion.String()
	PublicNetwork_GroupVersionKind = CRDGroupVersion.WithKind(PublicNetwork_Kind)
)

func init() {
	SchemeBuilder.Register(&PublicNetwork{}, &PublicNetworkList{})
}

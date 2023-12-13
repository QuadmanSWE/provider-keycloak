/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ProtocolMapperObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ProtocolMapperParameters struct {

	// The mapper's associated client. Cannot be used at the same time as client_scope_id.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// Reference to a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDRef *v1.Reference `json:"clientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate clientId.
	// +kubebuilder:validation:Optional
	ClientIDSelector *v1.Selector `json:"clientIdSelector,omitempty" tf:"-"`

	// The mapper's associated client scope. Cannot be used at the same time as client_id.
	// +kubebuilder:validation:Optional
	ClientScopeID *string `json:"clientScopeId,omitempty" tf:"client_scope_id,omitempty"`

	// +kubebuilder:validation:Required
	Config map[string]*string `json:"config" tf:"config,omitempty"`

	// A human-friendly name that will appear in the Keycloak console.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The protocol of the client (openid-connect / saml).
	// +kubebuilder:validation:Required
	Protocol *string `json:"protocol" tf:"protocol,omitempty"`

	// The type of the protocol mapper.
	// +kubebuilder:validation:Required
	ProtocolMapper *string `json:"protocolMapper" tf:"protocol_mapper,omitempty"`

	// The realm id where the associated client or client scope exists.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`
}

// ProtocolMapperSpec defines the desired state of ProtocolMapper
type ProtocolMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProtocolMapperParameters `json:"forProvider"`
}

// ProtocolMapperStatus defines the observed state of ProtocolMapper.
type ProtocolMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProtocolMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolMapper is the Schema for the ProtocolMappers API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ProtocolMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProtocolMapperSpec   `json:"spec"`
	Status            ProtocolMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProtocolMapperList contains a list of ProtocolMappers
type ProtocolMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProtocolMapper `json:"items"`
}

// Repository type metadata.
var (
	ProtocolMapper_Kind             = "ProtocolMapper"
	ProtocolMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProtocolMapper_Kind}.String()
	ProtocolMapper_KindAPIVersion   = ProtocolMapper_Kind + "." + CRDGroupVersion.String()
	ProtocolMapper_GroupVersionKind = CRDGroupVersion.WithKind(ProtocolMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&ProtocolMapper{}, &ProtocolMapperList{})
}
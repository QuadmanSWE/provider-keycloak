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

type ClientServiceAccountRealmRoleInitParameters struct {

	// The realm that the client and role belong to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The name of the role that is assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Client in openidclient to populate serviceAccountUserId.
	// +kubebuilder:validation:Optional
	ServiceAccountUserClientIDRef *v1.Reference `json:"serviceAccountUserClientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate serviceAccountUserId.
	// +kubebuilder:validation:Optional
	ServiceAccountUserClientIDSelector *v1.Selector `json:"serviceAccountUserClientIdSelector,omitempty" tf:"-"`

	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.ServiceAccountRoleIDExtractor()
	// +crossplane:generate:reference:refFieldName=ServiceAccountUserClientIDRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountUserClientIDSelector
	ServiceAccountUserID *string `json:"serviceAccountUserId,omitempty" tf:"service_account_user_id,omitempty"`
}

type ClientServiceAccountRealmRoleObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The realm that the client and role belong to.
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// The name of the role that is assigned.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	ServiceAccountUserID *string `json:"serviceAccountUserId,omitempty" tf:"service_account_user_id,omitempty"`
}

type ClientServiceAccountRealmRoleParameters struct {

	// The realm that the client and role belong to.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	RealmID *string `json:"realmId,omitempty" tf:"realm_id,omitempty"`

	// Reference to a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDRef *v1.Reference `json:"realmIdRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realmId.
	// +kubebuilder:validation:Optional
	RealmIDSelector *v1.Selector `json:"realmIdSelector,omitempty" tf:"-"`

	// The name of the role that is assigned.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// Reference to a Client in openidclient to populate serviceAccountUserId.
	// +kubebuilder:validation:Optional
	ServiceAccountUserClientIDRef *v1.Reference `json:"serviceAccountUserClientIdRef,omitempty" tf:"-"`

	// Selector for a Client in openidclient to populate serviceAccountUserId.
	// +kubebuilder:validation:Optional
	ServiceAccountUserClientIDSelector *v1.Selector `json:"serviceAccountUserClientIdSelector,omitempty" tf:"-"`

	// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/openidclient/v1alpha1.Client
	// +crossplane:generate:reference:extractor=github.com/crossplane-contrib/provider-keycloak/config/common.ServiceAccountRoleIDExtractor()
	// +crossplane:generate:reference:refFieldName=ServiceAccountUserClientIDRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountUserClientIDSelector
	// +kubebuilder:validation:Optional
	ServiceAccountUserID *string `json:"serviceAccountUserId,omitempty" tf:"service_account_user_id,omitempty"`
}

// ClientServiceAccountRealmRoleSpec defines the desired state of ClientServiceAccountRealmRole
type ClientServiceAccountRealmRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClientServiceAccountRealmRoleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClientServiceAccountRealmRoleInitParameters `json:"initProvider,omitempty"`
}

// ClientServiceAccountRealmRoleStatus defines the observed state of ClientServiceAccountRealmRole.
type ClientServiceAccountRealmRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClientServiceAccountRealmRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ClientServiceAccountRealmRole is the Schema for the ClientServiceAccountRealmRoles API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type ClientServiceAccountRealmRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   ClientServiceAccountRealmRoleSpec   `json:"spec"`
	Status ClientServiceAccountRealmRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClientServiceAccountRealmRoleList contains a list of ClientServiceAccountRealmRoles
type ClientServiceAccountRealmRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClientServiceAccountRealmRole `json:"items"`
}

// Repository type metadata.
var (
	ClientServiceAccountRealmRole_Kind             = "ClientServiceAccountRealmRole"
	ClientServiceAccountRealmRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClientServiceAccountRealmRole_Kind}.String()
	ClientServiceAccountRealmRole_KindAPIVersion   = ClientServiceAccountRealmRole_Kind + "." + CRDGroupVersion.String()
	ClientServiceAccountRealmRole_GroupVersionKind = CRDGroupVersion.WithKind(ClientServiceAccountRealmRole_Kind)
)

func init() {
	SchemeBuilder.Register(&ClientServiceAccountRealmRole{}, &ClientServiceAccountRealmRoleList{})
}

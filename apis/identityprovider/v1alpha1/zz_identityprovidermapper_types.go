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

type IdentityProviderMapperInitParameters struct {

	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	// The alias of the associated identity provider.
	// IDP Alias
	IdentityProviderAlias *string `json:"identityProviderAlias,omitempty" tf:"identity_provider_alias,omitempty"`

	// The type of the identity provider mapper. This can be a format string that includes a %s - this will be replaced by the provider id.
	// IDP Mapper Type
	IdentityProviderMapper *string `json:"identityProviderMapper,omitempty" tf:"identity_provider_mapper,omitempty"`

	// The name of the mapper.
	// IDP Mapper Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the realm.
	// Realm Name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Reference to a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmRef *v1.Reference `json:"realmRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmSelector *v1.Selector `json:"realmSelector,omitempty" tf:"-"`
}

type IdentityProviderMapperObservation struct {

	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The alias of the associated identity provider.
	// IDP Alias
	IdentityProviderAlias *string `json:"identityProviderAlias,omitempty" tf:"identity_provider_alias,omitempty"`

	// The type of the identity provider mapper. This can be a format string that includes a %s - this will be replaced by the provider id.
	// IDP Mapper Type
	IdentityProviderMapper *string `json:"identityProviderMapper,omitempty" tf:"identity_provider_mapper,omitempty"`

	// The name of the mapper.
	// IDP Mapper Name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the realm.
	// Realm Name
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`
}

type IdentityProviderMapperParameters struct {

	// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ExtraConfig map[string]*string `json:"extraConfig,omitempty" tf:"extra_config,omitempty"`

	// The alias of the associated identity provider.
	// IDP Alias
	// +kubebuilder:validation:Optional
	IdentityProviderAlias *string `json:"identityProviderAlias,omitempty" tf:"identity_provider_alias,omitempty"`

	// The type of the identity provider mapper. This can be a format string that includes a %s - this will be replaced by the provider id.
	// IDP Mapper Type
	// +kubebuilder:validation:Optional
	IdentityProviderMapper *string `json:"identityProviderMapper,omitempty" tf:"identity_provider_mapper,omitempty"`

	// The name of the mapper.
	// IDP Mapper Name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of the realm.
	// Realm Name
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-keycloak/apis/realm/v1alpha1.Realm
	// +kubebuilder:validation:Optional
	Realm *string `json:"realm,omitempty" tf:"realm,omitempty"`

	// Reference to a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmRef *v1.Reference `json:"realmRef,omitempty" tf:"-"`

	// Selector for a Realm in realm to populate realm.
	// +kubebuilder:validation:Optional
	RealmSelector *v1.Selector `json:"realmSelector,omitempty" tf:"-"`
}

// IdentityProviderMapperSpec defines the desired state of IdentityProviderMapper
type IdentityProviderMapperSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     IdentityProviderMapperParameters `json:"forProvider"`
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
	InitProvider IdentityProviderMapperInitParameters `json:"initProvider,omitempty"`
}

// IdentityProviderMapperStatus defines the observed state of IdentityProviderMapper.
type IdentityProviderMapperStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        IdentityProviderMapperObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// IdentityProviderMapper is the Schema for the IdentityProviderMappers API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,keycloak}
type IdentityProviderMapper struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityProviderAlias) || (has(self.initProvider) && has(self.initProvider.identityProviderAlias))",message="spec.forProvider.identityProviderAlias is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.identityProviderMapper) || (has(self.initProvider) && has(self.initProvider.identityProviderMapper))",message="spec.forProvider.identityProviderMapper is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   IdentityProviderMapperSpec   `json:"spec"`
	Status IdentityProviderMapperStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IdentityProviderMapperList contains a list of IdentityProviderMappers
type IdentityProviderMapperList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IdentityProviderMapper `json:"items"`
}

// Repository type metadata.
var (
	IdentityProviderMapper_Kind             = "IdentityProviderMapper"
	IdentityProviderMapper_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: IdentityProviderMapper_Kind}.String()
	IdentityProviderMapper_KindAPIVersion   = IdentityProviderMapper_Kind + "." + CRDGroupVersion.String()
	IdentityProviderMapper_GroupVersionKind = CRDGroupVersion.WithKind(IdentityProviderMapper_Kind)
)

func init() {
	SchemeBuilder.Register(&IdentityProviderMapper{}, &IdentityProviderMapperList{})
}

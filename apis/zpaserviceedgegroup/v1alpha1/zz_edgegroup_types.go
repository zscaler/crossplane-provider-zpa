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

type EdgeGroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of the version profile.
	VersionProfileVisibilityScope *string `json:"versionProfileVisibilityScope,omitempty" tf:"version_profile_visibility_scope,omitempty"`
}

type EdgeGroupParameters struct {

	// This field controls dynamic discovery of the servers.
	// +kubebuilder:validation:Optional
	CityCountry *string `json:"cityCountry,omitempty" tf:"city_country,omitempty"`

	// This field is an array of app-connector-id only.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// Description of the Service Edge Group.
	// Description of the Service Edge Group.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this Service Edge Group is enabled or not. Default value: true Supported values: true, false
	// Whether this Service Edge Group is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Enable or disable public access for the Service Edge Group. Default value: FALSE Supported values: DEFAULT, TRUE, FALSE
	// Enable or disable public access for the Service Edge Group.
	// +kubebuilder:validation:Optional
	IsPublic *bool `json:"isPublic,omitempty" tf:"is_public,omitempty"`

	// Latitude for the Service Edge Group. Integer or decimal with values in the range of -90 to 90
	// Latitude for the Service Edge Group.
	// +kubebuilder:validation:Required
	Latitude *string `json:"latitude" tf:"latitude,omitempty"`

	// Location for the Service Edge Group.
	// Location for the Service Edge Group.
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Longitude for the Service Edge Group. Integer or decimal with values in the range of -180 to 180
	// Longitude for the Service Edge Group.
	// +kubebuilder:validation:Required
	Longitude *string `json:"longitude" tf:"longitude,omitempty"`

	// Name of the Service Edge Group.
	// Name of the Service Edge Group.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether the default version profile of the App Connector Group is applied or overridden. Default: true. Supported values: true, false
	// Whether the default version profile of the App Connector Group is applied or overridden.
	// +kubebuilder:validation:Optional
	OverrideVersionProfile *bool `json:"overrideVersionProfile,omitempty" tf:"override_version_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ServiceEdges []ServiceEdgesParameters `json:"serviceEdges,omitempty" tf:"service_edges,omitempty"`

	// Trusted networks for this Service Edge Group. List of trusted network objects
	// +kubebuilder:validation:Optional
	TrustedNetworks []TrustedNetworksParameters `json:"trustedNetworks,omitempty" tf:"trusted_networks,omitempty"`

	// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: SUNDAY List of valid days (i.e., Sunday, Monday)
	// Service Edges in this group will attempt to update to a newer version of the software during this specified day.
	// +kubebuilder:validation:Optional
	UpgradeDay *string `json:"upgradeDay,omitempty" tf:"upgrade_day,omitempty"`

	// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: 66600 Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than 86400, in 15 minute intervals
	// Service Edges in this group will attempt to update to a newer version of the software during this specified time.
	// +kubebuilder:validation:Optional
	UpgradeTimeInSecs *string `json:"upgradeTimeInSecs,omitempty" tf:"upgrade_time_in_secs,omitempty"`

	// ID of the version profile. To learn more, see Version Profile Use Cases. This value becomes required if the value for override_version_profile is set to true.
	// ID of the version profile.
	// +kubebuilder:validation:Optional
	VersionProfileID *string `json:"versionProfileId,omitempty" tf:"version_profile_id,omitempty"`

	// Name of the Service Edge Group.
	// ID of the version profile.
	// +kubebuilder:validation:Optional
	VersionProfileName *string `json:"versionProfileName,omitempty" tf:"version_profile_name,omitempty"`
}

type ServiceEdgesObservation struct {
}

type ServiceEdgesParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

type TrustedNetworksObservation struct {
}

type TrustedNetworksParameters struct {

	// +kubebuilder:validation:Optional
	ID []*string `json:"id,omitempty" tf:"id,omitempty"`
}

// EdgeGroupSpec defines the desired state of EdgeGroup
type EdgeGroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EdgeGroupParameters `json:"forProvider"`
}

// EdgeGroupStatus defines the observed state of EdgeGroup.
type EdgeGroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EdgeGroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGroup is the Schema for the EdgeGroups API. Creates and manages ZPA Service Edge Group details.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type EdgeGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EdgeGroupSpec   `json:"spec"`
	Status            EdgeGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EdgeGroupList contains a list of EdgeGroups
type EdgeGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EdgeGroup `json:"items"`
}

// Repository type metadata.
var (
	EdgeGroup_Kind             = "EdgeGroup"
	EdgeGroup_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EdgeGroup_Kind}.String()
	EdgeGroup_KindAPIVersion   = EdgeGroup_Kind + "." + CRDGroupVersion.String()
	EdgeGroup_GroupVersionKind = CRDGroupVersion.WithKind(EdgeGroup_Kind)
)

func init() {
	SchemeBuilder.Register(&EdgeGroup{}, &EdgeGroupList{})
}

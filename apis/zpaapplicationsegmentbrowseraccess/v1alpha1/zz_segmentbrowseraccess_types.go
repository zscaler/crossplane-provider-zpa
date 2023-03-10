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

type ClientlessAppsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClientlessAppsParameters struct {

	// - If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: true and false
	// If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.
	// +kubebuilder:validation:Optional
	AllowOptions *bool `json:"allowOptions,omitempty" tf:"allow_options,omitempty"`

	// - Port for the BA app.
	// Port for the BA app.
	// +kubebuilder:validation:Required
	ApplicationPort *string `json:"applicationPort" tf:"application_port,omitempty"`

	// - Protocol for the BA app. Supported values: HTTP and HTTPS
	// Protocol for the BA app.
	// +kubebuilder:validation:Required
	ApplicationProtocol *string `json:"applicationProtocol" tf:"application_protocol,omitempty"`

	// - ID of the BA certificate. Refer to the data source documentation for zpa_ba_certificate
	// ID of the BA certificate.
	// +kubebuilder:validation:Required
	CertificateID *string `json:"certificateId" tf:"certificate_id,omitempty"`

	// +kubebuilder:validation:Optional
	Cname *string `json:"cname,omitempty" tf:"cname,omitempty"`

	// Description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// - Domain name or IP address of the BA app.
	// Domain name or IP address of the BA app.
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Whether this app is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Hidden *bool `json:"hidden,omitempty" tf:"hidden,omitempty"`

	// +kubebuilder:validation:Optional
	LocalDomain *string `json:"localDomain,omitempty" tf:"local_domain,omitempty"`

	// Name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Path *string `json:"path,omitempty" tf:"path,omitempty"`

	// Indicates whether Use Untrusted Certificates is enabled or disabled for a BA app.
	// +kubebuilder:validation:Optional
	TrustUntrustedCert *bool `json:"trustUntrustedCert,omitempty" tf:"trust_untrusted_cert,omitempty"`
}

type SegmentBrowserAccessObservation struct {

	// +kubebuilder:validation:Required
	ClientlessApps []ClientlessAppsObservation `json:"clientlessApps,omitempty" tf:"clientless_apps,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type SegmentBrowserAccessParameters struct {

	// Indicates whether users can bypass ZPA to access applications. Default value is: NEVER and supported values are: ALWAYS, NEVER and ON_NET. The value NEVER indicates the use of the client forwarding policy.
	// Indicates whether users can bypass ZPA to access applications. Default: NEVER. Supported values: ALWAYS, NEVER, ON_NET. The value NEVER indicates the use of the client forwarding policy.
	// +kubebuilder:validation:Optional
	BypassType *string `json:"bypassType,omitempty" tf:"bypass_type,omitempty"`

	// +kubebuilder:validation:Required
	ClientlessApps []ClientlessAppsParameters `json:"clientlessApps" tf:"clientless_apps,omitempty"`

	// Default: DEFAULT. Supported values: DEFAULT, SIEM
	// +kubebuilder:validation:Optional
	ConfigSpace *string `json:"configSpace,omitempty" tf:"config_space,omitempty"`

	// Description of the application.
	// Description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// List of domains and IPs.
	// List of domains and IPs.
	// +kubebuilder:validation:Required
	DomainNames []*string `json:"domainNames" tf:"domain_names,omitempty"`

	// Whether Double Encryption is enabled or disabled for the app.
	// Whether Double Encryption is enabled or disabled for the app.
	// +kubebuilder:validation:Optional
	DoubleEncrypt *bool `json:"doubleEncrypt,omitempty" tf:"double_encrypt,omitempty"`

	// Whether this app is enabled or not.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Default: DEFAULT. Supported values: DEFAULT, NONE
	// +kubebuilder:validation:Optional
	HealthCheckType *string `json:"healthCheckType,omitempty" tf:"health_check_type,omitempty"`

	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	// Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
	// +kubebuilder:validation:Optional
	HealthReporting *string `json:"healthReporting,omitempty" tf:"health_reporting,omitempty"`

	// If Source IP Anchoring for use with ZIA, is enabled or disabled for the app. Supported values are true and false
	// +kubebuilder:validation:Optional
	IPAnchored *bool `json:"ipAnchored,omitempty" tf:"ip_anchored,omitempty"`

	// +kubebuilder:validation:Optional
	IcmpAccessType *string `json:"icmpAccessType,omitempty" tf:"icmp_access_type,omitempty"`

	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	// Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
	// +kubebuilder:validation:Optional
	IsCnameEnabled *bool `json:"isCnameEnabled,omitempty" tf:"is_cname_enabled,omitempty"`

	// Name of the application.
	// Name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Whether this app is enabled or not.
	// +kubebuilder:validation:Optional
	PassiveHealthEnabled *bool `json:"passiveHealthEnabled,omitempty" tf:"passive_health_enabled,omitempty"`

	// List of Segment Group IDs
	// +kubebuilder:validation:Required
	SegmentGroupID *string `json:"segmentGroupId" tf:"segment_group_id,omitempty"`

	// Name of the application.
	// +kubebuilder:validation:Optional
	SegmentGroupName *string `json:"segmentGroupName,omitempty" tf:"segment_group_name,omitempty"`

	// List of Server Group IDs
	// List of the server group IDs.
	// +kubebuilder:validation:Required
	ServerGroups []ServerGroupsParameters `json:"serverGroups" tf:"server_groups,omitempty"`

	// tcp port range
	// +kubebuilder:validation:Optional
	TCPPortRange []TCPPortRangeParameters `json:"tcpPortRange,omitempty" tf:"tcp_port_range,omitempty"`

	// TCP port ranges used to access the app.
	// TCP port ranges used to access the app.
	// +kubebuilder:validation:Optional
	TCPPortRanges []*string `json:"tcpPortRanges,omitempty" tf:"tcp_port_ranges,omitempty"`

	// udp port range
	// +kubebuilder:validation:Optional
	UDPPortRange []UDPPortRangeParameters `json:"udpPortRange,omitempty" tf:"udp_port_range,omitempty"`

	// UDP port ranges used to access the app.
	// UDP port ranges used to access the app.
	// +kubebuilder:validation:Optional
	UDPPortRanges []*string `json:"udpPortRanges,omitempty" tf:"udp_port_ranges,omitempty"`
}

type ServerGroupsObservation struct {
}

type ServerGroupsParameters struct {

	// +kubebuilder:validation:Required
	ID []*string `json:"id" tf:"id,omitempty"`
}

type TCPPortRangeObservation struct {
}

type TCPPortRangeParameters struct {

	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from"`

	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to"`
}

type UDPPortRangeObservation struct {
}

type UDPPortRangeParameters struct {

	// +kubebuilder:validation:Optional
	From *string `json:"from,omitempty" tf:"from"`

	// +kubebuilder:validation:Optional
	To *string `json:"to,omitempty" tf:"to"`
}

// SegmentBrowserAccessSpec defines the desired state of SegmentBrowserAccess
type SegmentBrowserAccessSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SegmentBrowserAccessParameters `json:"forProvider"`
}

// SegmentBrowserAccessStatus defines the observed state of SegmentBrowserAccess.
type SegmentBrowserAccessStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SegmentBrowserAccessObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SegmentBrowserAccess is the Schema for the SegmentBrowserAccesss API. Creates and manages ZPA Browser Access Application Segment.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,zpa}
type SegmentBrowserAccess struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SegmentBrowserAccessSpec   `json:"spec"`
	Status            SegmentBrowserAccessStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SegmentBrowserAccessList contains a list of SegmentBrowserAccesss
type SegmentBrowserAccessList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SegmentBrowserAccess `json:"items"`
}

// Repository type metadata.
var (
	SegmentBrowserAccess_Kind             = "SegmentBrowserAccess"
	SegmentBrowserAccess_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SegmentBrowserAccess_Kind}.String()
	SegmentBrowserAccess_KindAPIVersion   = SegmentBrowserAccess_Kind + "." + CRDGroupVersion.String()
	SegmentBrowserAccess_GroupVersionKind = CRDGroupVersion.WithKind(SegmentBrowserAccess_Kind)
)

func init() {
	SchemeBuilder.Register(&SegmentBrowserAccess{}, &SegmentBrowserAccessList{})
}

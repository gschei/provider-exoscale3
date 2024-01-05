// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClusterInitParameters struct {

	// (Set of String, Deprecated)
	Addons []*string `json:"addons,omitempty" tf:"addons,omitempty"`

	// (Boolean) Enable automatic upgrading of the control plane version.
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// (String) The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni *string `json:"cni,omitempty" tf:"cni,omitempty"`

	// form text describing the cluster.
	// A free-form text describing the cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Deploy the Exoscale Cloud Controller Manager in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `json:"exoscaleCcm,omitempty" tf:"exoscale_ccm,omitempty"`

	// (Map of String) A map of key/value labels.
	// A map of key/value labels.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (Boolean) Deploy the Kubernetes Metrics Server in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool `json:"metricsServer,omitempty" tf:"metrics_server,omitempty"`

	// (String) The SKS cluster name.
	// The SKS cluster name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below. (see below for nested schema)
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc []OidcInitParameters `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// (String) The service level of the control plane (pro or starter; default: pro; may only be set at creation time).
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// (String) The version of the control plane (default: latest version available from the API; see exo compute sks versions for reference; may only be set at creation time).
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// (String) ❗ The Exoscale Zone name.
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ClusterObservation struct {

	// (Set of String, Deprecated)
	Addons []*string `json:"addons,omitempty" tf:"addons,omitempty"`

	// server).
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCA *string `json:"aggregationCa,omitempty" tf:"aggregation_ca,omitempty"`

	// (Boolean) Enable automatic upgrading of the control plane version.
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// (String) The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	Cni *string `json:"cni,omitempty" tf:"cni,omitempty"`

	// (String) The CA certificate (in PEM format) for TLS communications between control plane components.
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCA *string `json:"controlPlaneCa,omitempty" tf:"control_plane_ca,omitempty"`

	// (String) The cluster creation date.
	// The cluster creation date.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// form text describing the cluster.
	// A free-form text describing the cluster.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The cluster API endpoint.
	// The cluster API endpoint.
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// (Boolean) Deploy the Exoscale Cloud Controller Manager in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `json:"exoscaleCcm,omitempty" tf:"exoscale_ccm,omitempty"`

	// (String) The SKS cluster ID.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCA *string `json:"kubeletCa,omitempty" tf:"kubelet_ca,omitempty"`

	// (Map of String) A map of key/value labels.
	// A map of key/value labels.
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (Boolean) Deploy the Kubernetes Metrics Server in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool `json:"metricsServer,omitempty" tf:"metrics_server,omitempty"`

	// (String) The SKS cluster name.
	// The SKS cluster name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Set of String) The list of exoscale_sks_nodepool (IDs) attached to the cluster.
	// The list of [exoscale_sks_nodepool](./sks_nodepool.md) (IDs) attached to the cluster.
	Nodepools []*string `json:"nodepools,omitempty" tf:"nodepools,omitempty"`

	// (Block List, Max: 1) An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below. (see below for nested schema)
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc []OidcObservation `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// (String) The service level of the control plane (pro or starter; default: pro; may only be set at creation time).
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// (String) The cluster state.
	// The cluster state.
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// (String) The version of the control plane (default: latest version available from the API; see exo compute sks versions for reference; may only be set at creation time).
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// (String) ❗ The Exoscale Zone name.
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type ClusterParameters struct {

	// (Set of String, Deprecated)
	// +kubebuilder:validation:Optional
	Addons []*string `json:"addons,omitempty" tf:"addons,omitempty"`

	// (Boolean) Enable automatic upgrading of the control plane version.
	// Enable automatic upgrading of the control plane version.
	// +kubebuilder:validation:Optional
	AutoUpgrade *bool `json:"autoUpgrade,omitempty" tf:"auto_upgrade,omitempty"`

	// (String) The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	// The CNI plugin that is to be used. Available options are "calico" or "cilium". Defaults to "calico". Setting empty string will result in a cluster with no CNI.
	// +kubebuilder:validation:Optional
	Cni *string `json:"cni,omitempty" tf:"cni,omitempty"`

	// form text describing the cluster.
	// A free-form text describing the cluster.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Boolean) Deploy the Exoscale Cloud Controller Manager in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	// +kubebuilder:validation:Optional
	ExoscaleCcm *bool `json:"exoscaleCcm,omitempty" tf:"exoscale_ccm,omitempty"`

	// (Map of String) A map of key/value labels.
	// A map of key/value labels.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (Boolean) Deploy the Kubernetes Metrics Server in the control plane (boolean; default: true; may only be set at creation time).
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	// +kubebuilder:validation:Optional
	MetricsServer *bool `json:"metricsServer,omitempty" tf:"metrics_server,omitempty"`

	// (String) The SKS cluster name.
	// The SKS cluster name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (Block List, Max: 1) An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below. (see below for nested schema)
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	// +kubebuilder:validation:Optional
	Oidc []OidcParameters `json:"oidc,omitempty" tf:"oidc,omitempty"`

	// (String) The service level of the control plane (pro or starter; default: pro; may only be set at creation time).
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	// +kubebuilder:validation:Optional
	ServiceLevel *string `json:"serviceLevel,omitempty" tf:"service_level,omitempty"`

	// (String) The version of the control plane (default: latest version available from the API; see exo compute sks versions for reference; may only be set at creation time).
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`

	// (String) ❗ The Exoscale Zone name.
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

type OidcInitParameters struct {

	// (String) The OpenID client ID.
	// The OpenID client ID.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String) An OpenID JWT claim to use as the user's group.
	// An OpenID JWT claim to use as the user's group.
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// (String) An OpenID prefix prepended to group claims.
	// An OpenID prefix prepended to group claims.
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// (String) The OpenID provider URL.
	// The OpenID provider URL.
	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	// (Map of String) A map of key/value pairs that describes a required claim in the OpenID Token.
	// A map of key/value pairs that describes a required claim in the OpenID Token.
	RequiredClaim map[string]*string `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// (String) An OpenID JWT claim to use as the user name.
	// An OpenID JWT claim to use as the user name.
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// (String) An OpenID prefix prepended to username claims.
	// An OpenID prefix prepended to username claims.
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

type OidcObservation struct {

	// (String) The OpenID client ID.
	// The OpenID client ID.
	ClientID *string `json:"clientId,omitempty" tf:"client_id,omitempty"`

	// (String) An OpenID JWT claim to use as the user's group.
	// An OpenID JWT claim to use as the user's group.
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// (String) An OpenID prefix prepended to group claims.
	// An OpenID prefix prepended to group claims.
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// (String) The OpenID provider URL.
	// The OpenID provider URL.
	IssuerURL *string `json:"issuerUrl,omitempty" tf:"issuer_url,omitempty"`

	// (Map of String) A map of key/value pairs that describes a required claim in the OpenID Token.
	// A map of key/value pairs that describes a required claim in the OpenID Token.
	RequiredClaim map[string]*string `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// (String) An OpenID JWT claim to use as the user name.
	// An OpenID JWT claim to use as the user name.
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// (String) An OpenID prefix prepended to username claims.
	// An OpenID prefix prepended to username claims.
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

type OidcParameters struct {

	// (String) The OpenID client ID.
	// The OpenID client ID.
	// +kubebuilder:validation:Optional
	ClientID *string `json:"clientId" tf:"client_id,omitempty"`

	// (String) An OpenID JWT claim to use as the user's group.
	// An OpenID JWT claim to use as the user's group.
	// +kubebuilder:validation:Optional
	GroupsClaim *string `json:"groupsClaim,omitempty" tf:"groups_claim,omitempty"`

	// (String) An OpenID prefix prepended to group claims.
	// An OpenID prefix prepended to group claims.
	// +kubebuilder:validation:Optional
	GroupsPrefix *string `json:"groupsPrefix,omitempty" tf:"groups_prefix,omitempty"`

	// (String) The OpenID provider URL.
	// The OpenID provider URL.
	// +kubebuilder:validation:Optional
	IssuerURL *string `json:"issuerUrl" tf:"issuer_url,omitempty"`

	// (Map of String) A map of key/value pairs that describes a required claim in the OpenID Token.
	// A map of key/value pairs that describes a required claim in the OpenID Token.
	// +kubebuilder:validation:Optional
	RequiredClaim map[string]*string `json:"requiredClaim,omitempty" tf:"required_claim,omitempty"`

	// (String) An OpenID JWT claim to use as the user name.
	// An OpenID JWT claim to use as the user name.
	// +kubebuilder:validation:Optional
	UsernameClaim *string `json:"usernameClaim,omitempty" tf:"username_claim,omitempty"`

	// (String) An OpenID prefix prepended to username claims.
	// An OpenID prefix prepended to username claims.
	// +kubebuilder:validation:Optional
	UsernamePrefix *string `json:"usernamePrefix,omitempty" tf:"username_prefix,omitempty"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterParameters `json:"forProvider"`
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
	InitProvider ClusterInitParameters `json:"initProvider,omitempty"`
}

// ClusterStatus defines the observed state of Cluster.
type ClusterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Cluster is the Schema for the Clusters API. Manage Exoscale Scalable Kubernetes Service (SKS) https://community.exoscale.com/documentation/sks/ Clusters.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,exoscale3}
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zone) || (has(self.initProvider) && has(self.initProvider.zone))",message="spec.forProvider.zone is a required parameter"
	Spec   ClusterSpec   `json:"spec"`
	Status ClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cluster `json:"items"`
}

// Repository type metadata.
var (
	Cluster_Kind             = "Cluster"
	Cluster_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Cluster_Kind}.String()
	Cluster_KindAPIVersion   = Cluster_Kind + "." + CRDGroupVersion.String()
	Cluster_GroupVersionKind = CRDGroupVersion.WithKind(Cluster_Kind)
)

func init() {
	SchemeBuilder.Register(&Cluster{}, &ClusterList{})
}

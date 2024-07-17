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

type DomainInitParameters struct {
	CnameCoalescingEnabled *bool `json:"cnameCoalescingEnabled,omitempty" tf:"cname_coalescing_enabled,omitempty"`

	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Contract *string `json:"contract,omitempty" tf:"contract,omitempty"`

	DefaultErrorPenalty *float64 `json:"defaultErrorPenalty,omitempty" tf:"default_error_penalty,omitempty"`

	DefaultSSLClientCertificate *string `json:"defaultSslClientCertificate,omitempty" tf:"default_ssl_client_certificate,omitempty"`

	DefaultSSLClientPrivateKey *string `json:"defaultSslClientPrivateKey,omitempty" tf:"default_ssl_client_private_key,omitempty"`

	DefaultTimeoutPenalty *float64 `json:"defaultTimeoutPenalty,omitempty" tf:"default_timeout_penalty,omitempty"`

	EmailNotificationList []*string `json:"emailNotificationList,omitempty" tf:"email_notification_list,omitempty"`

	EndUserMappingEnabled *bool `json:"endUserMappingEnabled,omitempty" tf:"end_user_mapping_enabled,omitempty"`

	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	LoadFeedback *bool `json:"loadFeedback,omitempty" tf:"load_feedback,omitempty"`

	LoadImbalancePercentage *float64 `json:"loadImbalancePercentage,omitempty" tf:"load_imbalance_percentage,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If set (true) we will sign the domain's resource records so that they can be validated by a validating resolver.
	SignAndServe *bool `json:"signAndServe,omitempty" tf:"sign_and_serve,omitempty"`

	// The signing algorithm to use for signAndServe. One of the following values: RSA_SHA1, RSA_SHA256, RSA_SHA512, ECDSA_P256_SHA256, ECDSA_P384_SHA384, ED25519, ED448.
	SignAndServeAlgorithm *string `json:"signAndServeAlgorithm,omitempty" tf:"sign_and_serve_algorithm,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	WaitOnComplete *bool `json:"waitOnComplete,omitempty" tf:"wait_on_complete,omitempty"`
}

type DomainObservation struct {
	CnameCoalescingEnabled *bool `json:"cnameCoalescingEnabled,omitempty" tf:"cname_coalescing_enabled,omitempty"`

	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	Contract *string `json:"contract,omitempty" tf:"contract,omitempty"`

	DefaultErrorPenalty *float64 `json:"defaultErrorPenalty,omitempty" tf:"default_error_penalty,omitempty"`

	DefaultHealthMax *float64 `json:"defaultHealthMax,omitempty" tf:"default_health_max,omitempty"`

	DefaultHealthMultiplier *float64 `json:"defaultHealthMultiplier,omitempty" tf:"default_health_multiplier,omitempty"`

	DefaultHealthThreshold *float64 `json:"defaultHealthThreshold,omitempty" tf:"default_health_threshold,omitempty"`

	DefaultMaxUnreachablePenalty *float64 `json:"defaultMaxUnreachablePenalty,omitempty" tf:"default_max_unreachable_penalty,omitempty"`

	DefaultSSLClientCertificate *string `json:"defaultSslClientCertificate,omitempty" tf:"default_ssl_client_certificate,omitempty"`

	DefaultSSLClientPrivateKey *string `json:"defaultSslClientPrivateKey,omitempty" tf:"default_ssl_client_private_key,omitempty"`

	DefaultTimeoutPenalty *float64 `json:"defaultTimeoutPenalty,omitempty" tf:"default_timeout_penalty,omitempty"`

	DefaultUnreachableThreshold *float64 `json:"defaultUnreachableThreshold,omitempty" tf:"default_unreachable_threshold,omitempty"`

	EmailNotificationList []*string `json:"emailNotificationList,omitempty" tf:"email_notification_list,omitempty"`

	EndUserMappingEnabled *bool `json:"endUserMappingEnabled,omitempty" tf:"end_user_mapping_enabled,omitempty"`

	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LoadFeedback *bool `json:"loadFeedback,omitempty" tf:"load_feedback,omitempty"`

	LoadImbalancePercentage *float64 `json:"loadImbalancePercentage,omitempty" tf:"load_imbalance_percentage,omitempty"`

	MapUpdateInterval *float64 `json:"mapUpdateInterval,omitempty" tf:"map_update_interval,omitempty"`

	MaxProperties *float64 `json:"maxProperties,omitempty" tf:"max_properties,omitempty"`

	MaxResources *float64 `json:"maxResources,omitempty" tf:"max_resources,omitempty"`

	MaxTTL *float64 `json:"maxTtl,omitempty" tf:"max_ttl,omitempty"`

	MaxTestTimeout *float64 `json:"maxTestTimeout,omitempty" tf:"max_test_timeout,omitempty"`

	MinPingableRegionFraction *float64 `json:"minPingableRegionFraction,omitempty" tf:"min_pingable_region_fraction,omitempty"`

	MinTTL *float64 `json:"minTtl,omitempty" tf:"min_ttl,omitempty"`

	MinTestInterval *float64 `json:"minTestInterval,omitempty" tf:"min_test_interval,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	PingInterval *float64 `json:"pingInterval,omitempty" tf:"ping_interval,omitempty"`

	PingPacketSize *float64 `json:"pingPacketSize,omitempty" tf:"ping_packet_size,omitempty"`

	RoundRobinPrefix *string `json:"roundRobinPrefix,omitempty" tf:"round_robin_prefix,omitempty"`

	ServermonitorLivenessCount *float64 `json:"servermonitorLivenessCount,omitempty" tf:"servermonitor_liveness_count,omitempty"`

	ServermonitorLoadCount *float64 `json:"servermonitorLoadCount,omitempty" tf:"servermonitor_load_count,omitempty"`

	ServermonitorPool *string `json:"servermonitorPool,omitempty" tf:"servermonitor_pool,omitempty"`

	// If set (true) we will sign the domain's resource records so that they can be validated by a validating resolver.
	SignAndServe *bool `json:"signAndServe,omitempty" tf:"sign_and_serve,omitempty"`

	// The signing algorithm to use for signAndServe. One of the following values: RSA_SHA1, RSA_SHA256, RSA_SHA512, ECDSA_P256_SHA256, ECDSA_P384_SHA384, ED25519, ED448.
	SignAndServeAlgorithm *string `json:"signAndServeAlgorithm,omitempty" tf:"sign_and_serve_algorithm,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	WaitOnComplete *bool `json:"waitOnComplete,omitempty" tf:"wait_on_complete,omitempty"`
}

type DomainParameters struct {

	// +kubebuilder:validation:Optional
	CnameCoalescingEnabled *bool `json:"cnameCoalescingEnabled,omitempty" tf:"cname_coalescing_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Comment *string `json:"comment,omitempty" tf:"comment,omitempty"`

	// +kubebuilder:validation:Optional
	Contract *string `json:"contract,omitempty" tf:"contract,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultErrorPenalty *float64 `json:"defaultErrorPenalty,omitempty" tf:"default_error_penalty,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultSSLClientCertificate *string `json:"defaultSslClientCertificate,omitempty" tf:"default_ssl_client_certificate,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultSSLClientPrivateKey *string `json:"defaultSslClientPrivateKey,omitempty" tf:"default_ssl_client_private_key,omitempty"`

	// +kubebuilder:validation:Optional
	DefaultTimeoutPenalty *float64 `json:"defaultTimeoutPenalty,omitempty" tf:"default_timeout_penalty,omitempty"`

	// +kubebuilder:validation:Optional
	EmailNotificationList []*string `json:"emailNotificationList,omitempty" tf:"email_notification_list,omitempty"`

	// +kubebuilder:validation:Optional
	EndUserMappingEnabled *bool `json:"endUserMappingEnabled,omitempty" tf:"end_user_mapping_enabled,omitempty"`

	// +kubebuilder:validation:Optional
	Group *string `json:"group,omitempty" tf:"group,omitempty"`

	// +kubebuilder:validation:Optional
	LoadFeedback *bool `json:"loadFeedback,omitempty" tf:"load_feedback,omitempty"`

	// +kubebuilder:validation:Optional
	LoadImbalancePercentage *float64 `json:"loadImbalancePercentage,omitempty" tf:"load_imbalance_percentage,omitempty"`

	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// If set (true) we will sign the domain's resource records so that they can be validated by a validating resolver.
	// +kubebuilder:validation:Optional
	SignAndServe *bool `json:"signAndServe,omitempty" tf:"sign_and_serve,omitempty"`

	// The signing algorithm to use for signAndServe. One of the following values: RSA_SHA1, RSA_SHA256, RSA_SHA512, ECDSA_P256_SHA256, ECDSA_P384_SHA384, ED25519, ED448.
	// +kubebuilder:validation:Optional
	SignAndServeAlgorithm *string `json:"signAndServeAlgorithm,omitempty" tf:"sign_and_serve_algorithm,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	WaitOnComplete *bool `json:"waitOnComplete,omitempty" tf:"wait_on_complete,omitempty"`
}

// DomainSpec defines the desired state of Domain
type DomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainParameters `json:"forProvider"`
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
	InitProvider DomainInitParameters `json:"initProvider,omitempty"`
}

// DomainStatus defines the observed state of Domain.
type DomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Domain is the Schema for the Domains API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,akamai}
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	Spec   DomainSpec   `json:"spec"`
	Status DomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainList contains a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

// Repository type metadata.
var (
	Domain_Kind             = "Domain"
	Domain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Domain_Kind}.String()
	Domain_KindAPIVersion   = Domain_Kind + "." + CRDGroupVersion.String()
	Domain_GroupVersionKind = CRDGroupVersion.WithKind(Domain_Kind)
)

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
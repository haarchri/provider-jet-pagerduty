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

type IntegrationObservation struct {

	// The ID of the Event Orchestration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationParameters struct {
}

type OrchestrationObservation struct {

	// The ID of the Event Orchestration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An integration for the Event Orchestration.
	// +kubebuilder:validation:Optional
	Integration []IntegrationObservation `json:"integration,omitempty" tf:"integration,omitempty"`

	Routes *float64 `json:"routes,omitempty" tf:"routes,omitempty"`
}

type OrchestrationParameters struct {

	// A human-friendly description of the Event Orchestration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An integration for the Event Orchestration.
	// +kubebuilder:validation:Optional
	Integration []IntegrationParameters `json:"integration,omitempty" tf:"integration,omitempty"`

	// Name of the Event Orchestration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// ID of the team that owns the Event Orchestration. If none is specified, only admins have access.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	// +kubebuilder:validation:Optional
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// Reference to a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamRef *v1.Reference `json:"teamRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

type ParametersObservation struct {

	// Routing key that routes to this Orchestration.
	RoutingKey *string `json:"routingKey,omitempty" tf:"routing_key,omitempty"`

	// Type of the routing key. global is the default type.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ParametersParameters struct {
}

// OrchestrationSpec defines the desired state of Orchestration
type OrchestrationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrchestrationParameters `json:"forProvider"`
}

// OrchestrationStatus defines the observed state of Orchestration.
type OrchestrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Orchestration is the Schema for the Orchestrations API. Creates and manages a Global Event Orchestration in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Orchestration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrchestrationSpec   `json:"spec"`
	Status            OrchestrationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationList contains a list of Orchestrations
type OrchestrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Orchestration `json:"items"`
}

// Repository type metadata.
var (
	Orchestration_Kind             = "Orchestration"
	Orchestration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Orchestration_Kind}.String()
	Orchestration_KindAPIVersion   = Orchestration_Kind + "." + CRDGroupVersion.String()
	Orchestration_GroupVersionKind = CRDGroupVersion.WithKind(Orchestration_Kind)
)

func init() {
	SchemeBuilder.Register(&Orchestration{}, &OrchestrationList{})
}

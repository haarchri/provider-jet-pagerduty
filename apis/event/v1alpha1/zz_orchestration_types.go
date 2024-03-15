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

type IntegrationInitParameters struct {
}

type IntegrationObservation struct {

	// The ID of the Event Orchestration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Label *string `json:"label,omitempty" tf:"label,omitempty"`

	Parameters []ParametersObservation `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationParameters struct {
}

type OrchestrationInitParameters struct {

	// A human-friendly description of the Event Orchestration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An integration for the Event Orchestration.
	Integration []IntegrationInitParameters `json:"integration,omitempty" tf:"integration,omitempty"`

	// Name of the Event Orchestration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// ID of the team that owns the Event Orchestration. If none is specified, only admins have access.
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-pagerduty/apis/team/v1alpha1.Team
	Team *string `json:"team,omitempty" tf:"team,omitempty"`

	// Reference to a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamRef *v1.Reference `json:"teamRef,omitempty" tf:"-"`

	// Selector for a Team in team to populate team.
	// +kubebuilder:validation:Optional
	TeamSelector *v1.Selector `json:"teamSelector,omitempty" tf:"-"`
}

type OrchestrationObservation struct {

	// A human-friendly description of the Event Orchestration.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Event Orchestration.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An integration for the Event Orchestration.
	Integration []IntegrationObservation `json:"integration,omitempty" tf:"integration,omitempty"`

	// Name of the Event Orchestration.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Routes *float64 `json:"routes,omitempty" tf:"routes,omitempty"`

	// ID of the team that owns the Event Orchestration. If none is specified, only admins have access.
	Team *string `json:"team,omitempty" tf:"team,omitempty"`
}

type OrchestrationParameters struct {

	// A human-friendly description of the Event Orchestration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An integration for the Event Orchestration.
	// +kubebuilder:validation:Optional
	Integration []IntegrationParameters `json:"integration,omitempty" tf:"integration,omitempty"`

	// Name of the Event Orchestration.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

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

type ParametersInitParameters struct {
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
	InitProvider OrchestrationInitParameters `json:"initProvider,omitempty"`
}

// OrchestrationStatus defines the observed state of Orchestration.
type OrchestrationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Orchestration is the Schema for the Orchestrations API. Creates and manages an Event Orchestration in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type Orchestration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   OrchestrationSpec   `json:"spec"`
	Status OrchestrationStatus `json:"status,omitempty"`
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

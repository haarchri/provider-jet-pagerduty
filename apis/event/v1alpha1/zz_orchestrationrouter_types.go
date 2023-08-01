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

type ActionsInitParameters struct {

	// The ID of the target Service for the resulting alert.
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type ActionsObservation struct {

	// The ID of the target Service for the resulting alert.
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type ActionsParameters struct {

	// The ID of the target Service for the resulting alert.
	// +kubebuilder:validation:Optional
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type CatchAllInitParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	Actions []ActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`
}

type CatchAllObservation struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	Actions []ActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`
}

type CatchAllParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	// +kubebuilder:validation:Optional
	Actions []ActionsParameters `json:"actions,omitempty" tf:"actions,omitempty"`
}

type ConditionInitParameters struct {

	// A PCL condition string.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type ConditionObservation struct {

	// A PCL condition string.
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type ConditionParameters struct {

	// A PCL condition string.
	// +kubebuilder:validation:Optional
	Expression *string `json:"expression,omitempty" tf:"expression,omitempty"`
}

type OrchestrationRouterInitParameters struct {

	// When none of the rules match an event, the event will be routed according to the catch_all settings.
	CatchAll []CatchAllInitParameters `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// The Router contains a single set of rules  (the "start" set).
	Set []SetInitParameters `json:"set,omitempty" tf:"set,omitempty"`
}

type OrchestrationRouterObservation struct {

	// When none of the rules match an event, the event will be routed according to the catch_all settings.
	CatchAll []CatchAllObservation `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// ID of the Event Orchestration to which the Router belongs.
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// ID of the start set. Router supports only one set and it's id has to be start
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Router contains a single set of rules  (the "start" set).
	Set []SetObservation `json:"set,omitempty" tf:"set,omitempty"`
}

type OrchestrationRouterParameters struct {

	// When none of the rules match an event, the event will be routed according to the catch_all settings.
	// +kubebuilder:validation:Optional
	CatchAll []CatchAllParameters `json:"catchAll,omitempty" tf:"catch_all,omitempty"`

	// ID of the Event Orchestration to which the Router belongs.
	// +crossplane:generate:reference:type=Orchestration
	// +kubebuilder:validation:Optional
	EventOrchestration *string `json:"eventOrchestration,omitempty" tf:"event_orchestration,omitempty"`

	// Reference to a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationRef *v1.Reference `json:"eventOrchestrationRef,omitempty" tf:"-"`

	// Selector for a Orchestration to populate eventOrchestration.
	// +kubebuilder:validation:Optional
	EventOrchestrationSelector *v1.Selector `json:"eventOrchestrationSelector,omitempty" tf:"-"`

	// The Router contains a single set of rules  (the "start" set).
	// +kubebuilder:validation:Optional
	Set []SetParameters `json:"set,omitempty" tf:"set,omitempty"`
}

type RuleActionsInitParameters struct {

	// The ID of the target Service for the resulting alert.
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type RuleActionsObservation struct {

	// The ID of the target Service for the resulting alert.
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type RuleActionsParameters struct {

	// The ID of the target Service for the resulting alert.
	// +kubebuilder:validation:Optional
	RouteTo *string `json:"routeTo,omitempty" tf:"route_to,omitempty"`
}

type RuleInitParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	Actions []RuleActionsInitParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// Each of these conditions is evaluated to check if an event matches this rule. The rule is considered a match if any of these conditions match. If none are provided, the event will always match against the rule.
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A description of this rule's purpose.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type RuleObservation struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	Actions []RuleActionsObservation `json:"actions,omitempty" tf:"actions,omitempty"`

	// Each of these conditions is evaluated to check if an event matches this rule. The rule is considered a match if any of these conditions match. If none are provided, the event will always match against the rule.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// ID of the start set. Router supports only one set and it's id has to be start
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A description of this rule's purpose.
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type RuleParameters struct {

	// Actions that will be taken to change the resulting alert and incident, when an event matches this rule.
	// +kubebuilder:validation:Optional
	Actions []RuleActionsParameters `json:"actions,omitempty" tf:"actions,omitempty"`

	// Each of these conditions is evaluated to check if an event matches this rule. The rule is considered a match if any of these conditions match. If none are provided, the event will always match against the rule.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Indicates whether the rule is disabled and would therefore not be evaluated.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// A description of this rule's purpose.
	// +kubebuilder:validation:Optional
	Label *string `json:"label,omitempty" tf:"label,omitempty"`
}

type SetInitParameters struct {

	// ID of the start set. Router supports only one set and it's id has to be start
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Router evaluates Events against these Rules, one at a time, and routes each Event to a specific Service based on the first rule that matches.
	Rule []RuleInitParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

type SetObservation struct {

	// ID of the start set. Router supports only one set and it's id has to be start
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Router evaluates Events against these Rules, one at a time, and routes each Event to a specific Service based on the first rule that matches.
	Rule []RuleObservation `json:"rule,omitempty" tf:"rule,omitempty"`
}

type SetParameters struct {

	// ID of the start set. Router supports only one set and it's id has to be start
	// +kubebuilder:validation:Optional
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The Router evaluates Events against these Rules, one at a time, and routes each Event to a specific Service based on the first rule that matches.
	// +kubebuilder:validation:Optional
	Rule []RuleParameters `json:"rule,omitempty" tf:"rule,omitempty"`
}

// OrchestrationRouterSpec defines the desired state of OrchestrationRouter
type OrchestrationRouterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     OrchestrationRouterParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider OrchestrationRouterInitParameters `json:"initProvider,omitempty"`
}

// OrchestrationRouterStatus defines the observed state of OrchestrationRouter.
type OrchestrationRouterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        OrchestrationRouterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationRouter is the Schema for the OrchestrationRouters API. Creates and manages a Router for Global Event Orchestration in PagerDuty.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,pagerduty}
type OrchestrationRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.catchAll) || has(self.initProvider.catchAll)",message="catchAll is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.set) || has(self.initProvider.set)",message="set is a required parameter"
	Spec   OrchestrationRouterSpec   `json:"spec"`
	Status OrchestrationRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OrchestrationRouterList contains a list of OrchestrationRouters
type OrchestrationRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrchestrationRouter `json:"items"`
}

// Repository type metadata.
var (
	OrchestrationRouter_Kind             = "OrchestrationRouter"
	OrchestrationRouter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: OrchestrationRouter_Kind}.String()
	OrchestrationRouter_KindAPIVersion   = OrchestrationRouter_Kind + "." + CRDGroupVersion.String()
	OrchestrationRouter_GroupVersionKind = CRDGroupVersion.WithKind(OrchestrationRouter_Kind)
)

func init() {
	SchemeBuilder.Register(&OrchestrationRouter{}, &OrchestrationRouterList{})
}

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/webhook/resourcesemantics"
)

// Check that SampleSource can be validated and can be defaulted.
var _ runtime.Object = (*SampleSource)(nil)

var _ resourcesemantics.GenericCRD = (*SampleSource)(nil)

// Check that the type conforms to the duck Knative Resource shape.
var _ duckv1.KRShaped = (*SampleSource)(nil)

// SampleSourceSpec defines the desired state of SampleSource
// +kubebuilder:categories=all,knative,eventing,sources
type SampleSourceSpec struct {
	// ServiceAccountName holds the name of the Kubernetes service account
	// as which the underlying K8s resources should be run. If unspecified
	// this will default to the "default" service account for the namespace
	// in which the SampleSource exists.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`

	// Interval is the time between event emissions.
	// +optional
	Interval *metav1.Duration `json:"interval,omitempty"`

	// inherits duck/v1 SourceSpec, which currently provides:
	// * Sink - a reference to an object that will resolve to a domain name or
	//   a URI directly to use as the sink.
	// * CloudEventOverrides - defines overrides to control the output format
	//   and modifications of the event sent to the sink.
	duckv1.SourceSpec `json:",inline"`
}

// SecretValueFromSource represents the source of a secret value
type SecretValueFromSource struct {
	// The Secret key to select from.
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}

const (
	// SampleSourceConditionReady has status True when the
	// SampleSource is ready to send events.
	SampleSourceConditionReady = apis.ConditionReady

	// SampleSourceConditionSecretsProvided has status True when the
	// SampleSource has valid secret references
	SampleSourceConditionSecretsProvided apis.ConditionType = "SecretsProvided"

	// SampleSourceConditionSinkProvided has status True when the
	// SampleSource has been configured with a sink target.
	SampleSourceConditionSinkProvided apis.ConditionType = "SinkProvided"

	// SampleSourceConditionWebhookConfigured has a status True when the
	// SampleSource has been configured with a webhook.
	SampleSourceConditionWebhookConfigured apis.ConditionType = "WebhookConfigured"
)

var SampleSourceCondSet = apis.NewLivingConditionSet(
	SampleSourceConditionSecretsProvided,
	SampleSourceConditionSinkProvided,
	SampleSourceConditionWebhookConfigured)

// SampleSourceStatus defines the observed state of SampleSource
type SampleSourceStatus struct {
	// inherits duck/v1 SourceStatus, which currently provides:
	// * ObservedGeneration - the 'Generation' of the Service that was last
	//   processed by the controller.
	// * Conditions - the latest available observations of a resource's current
	//   state.
	// * SinkURI - the current active sink URI that has been configured for the
	//   Source.
	duckv1.SourceStatus `json:",inline"`
}

func (*SampleSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("SampleSource")
}

// GetConditionSet retrieves the condition set for this resource. Implements the KRShaped interface.
func (*SampleSource) GetConditionSet() apis.ConditionSet {
	return SampleSourceCondSet
}

// GetStatus retrieves the duck status for this resource. Implements the KRShaped interface.
func (g *SampleSource) GetStatus() *duckv1.Status {
	return &g.Status.Status
}

// GetCondition returns the condition currently associated with the given type, or nil.
func (s *SampleSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return SampleSourceCondSet.Manage(s).GetCondition(t)
}

// IsReady returns true if the resource is ready overall.
func (s *SampleSourceStatus) IsReady() bool {
	return SampleSourceCondSet.Manage(s).IsHappy()
}

// InitializeConditions sets relevant unset conditions to Unknown state.
func (s *SampleSourceStatus) InitializeConditions() {
	SampleSourceCondSet.Manage(s).InitializeConditions()
}

// MarkSecrets sets the condition that the source has a valid spec
func (s *SampleSourceStatus) MarkSecrets() {
	SampleSourceCondSet.Manage(s).MarkTrue(SampleSourceConditionSecretsProvided)
}

// MarkNoSecrets sets the condition that the source does not have a valid spec
func (s *SampleSourceStatus) MarkNoSecrets(reason, messageFormat string, messageA ...interface{}) {
	SampleSourceCondSet.Manage(s).MarkFalse(SampleSourceConditionSecretsProvided, reason, messageFormat, messageA...)
}

// MarkSink sets the condition that the source has a sink configured.
func (s *SampleSourceStatus) MarkSink(uri *apis.URL) {
	s.SinkURI = uri
	if uri != nil {
		SampleSourceCondSet.Manage(s).MarkTrue(SampleSourceConditionSinkProvided)
	} else {
		SampleSourceCondSet.Manage(s).MarkUnknown(SampleSourceConditionSinkProvided,
			"SinkEmpty", "Sink has resolved to empty.")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (s *SampleSourceStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	SampleSourceCondSet.Manage(s).MarkFalse(SampleSourceConditionSinkProvided, reason, messageFormat, messageA...)
}

// MarkWebhookConfigured sets the condition that the source has set its webhook configured.
func (s *SampleSourceStatus) MarkWebhookConfigured() {
	SampleSourceCondSet.Manage(s).MarkTrue(SampleSourceConditionWebhookConfigured)
}

// MarkWebhookNotConfigured sets the condition that the source does not have its webhook configured.
func (s *SampleSourceStatus) MarkWebhookNotConfigured(reason, messageFormat string, messageA ...interface{}) {
	SampleSourceCondSet.Manage(s).MarkFalse(SampleSourceConditionWebhookConfigured, reason, messageFormat, messageA...)
}

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleSource is the Schema for the SampleSources API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:categories=all,knative,eventing,sources
type SampleSource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleSourceSpec   `json:"spec,omitempty"`
	Status SampleSourceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleSourceList contains a list of SampleSource
type SampleSourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SampleSource `json:"items"`
}

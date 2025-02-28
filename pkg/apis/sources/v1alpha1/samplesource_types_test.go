package v1alpha1

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	corev1 "k8s.io/api/core/v1"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

var _ = duck.VerifyType(&SampleSource{}, &duckv1.Conditions{})

func TestSampleSourceGetConditionSet(t *testing.T) {
	r := &SampleSource{}

	if got, want := r.GetConditionSet().GetTopLevelConditionType(), apis.ConditionReady; got != want {
		t.Errorf("GetTopLevelCondition=%v, want=%v", got, want)
	}
}

func TestSampleSourceGetStatus(t *testing.T) {
	status := &duckv1.Status{}
	config := SampleSource{
		Status: SampleSourceStatus{
			SourceStatus: duckv1.SourceStatus{Status: *status},
		},
	}

	if !cmp.Equal(config.GetStatus(), status) {
		t.Errorf("GetStatus did not retrieve status. Got=%v Want=%v", config.GetStatus(), status)
	}
}

func TestSampleSourceStatusIsReady(t *testing.T) {
	tests := []struct {
		name string
		s    *SampleSourceStatus
		want bool
	}{{
		name: "uninitialized",
		s:    &SampleSourceStatus{},
		want: false,
	}, {
		name: "initialized",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			return s
		}(),
		want: false,
	}, {
		name: "mark sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			return s
		}(),
		want: false,
	}, {
		name: "mark secrets",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSecrets()
			return s
		}(),
		want: false,
	}, {
		name: "mark webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSecrets()
			return s
		}(),
		want: false,
	}, {
		name: "mark sink, secrets, webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			return s
		}(),
		want: true,
	}, {
		name: "mark sink, secrets, webhook, then no sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkNoSink("Testing", "")
			return s
		}(),
		want: false,
	}, {
		name: "mark sink, secrets, webhook, then no secrets",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkNoSecrets("Testing", "")
			return s
		}(),
		want: false,
	}, {
		name: "mark sink, secrets, webhook, then no webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkWebhookNotConfigured("Testing", "")
			return s
		}(),
		want: false,
	}, {
		name: "mark sink nil, secrets, webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(nil)
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			return s
		}(),
		want: false,
	}, {
		name: "mark sink nil, secrets, webhook, then sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(nil)
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkSink(apis.HTTP("example"))
			return s
		}(),
		want: true,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.IsReady()
			if diff := cmp.Diff(test.want, got); diff != "" {
				t.Errorf("%s: unexpected condition (-want, +got) = %v", test.name, diff)
			}
		})
	}
}

func TestSampleSourceStatusGetCondition(t *testing.T) {
	tests := []struct {
		name      string
		s         *SampleSourceStatus
		condQuery apis.ConditionType
		want      *apis.Condition
	}{{
		name:      "uninitialized",
		s:         &SampleSourceStatus{},
		condQuery: SampleSourceConditionReady,
		want:      nil,
	}, {
		name: "initialized",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark secrets",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSecrets()
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkWebhookConfigured()
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionUnknown,
		},
	}, {
		name: "mark sink, secrets, webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionTrue,
		},
	}, {
		name: "mark sink, secrets, webhook, then no sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkNoSink("Testing", "hi%s", "")
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:    SampleSourceConditionReady,
			Status:  corev1.ConditionFalse,
			Reason:  "Testing",
			Message: "hi",
		},
	}, {
		name: "mark sink, secrets, webhook, then no secrets",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkNoSecrets("Testing", "hi%s", "")
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:    SampleSourceConditionReady,
			Status:  corev1.ConditionFalse,
			Reason:  "Testing",
			Message: "hi",
		},
	}, {
		name: "mark sink, secrets, webhook, then no webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(apis.HTTP("example"))
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkWebhookNotConfigured("Testing", "hi%s", "")
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:    SampleSourceConditionReady,
			Status:  corev1.ConditionFalse,
			Reason:  "Testing",
			Message: "hi",
		},
	}, {
		name: "mark sink nil, secrets, webhook",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(nil)
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:    SampleSourceConditionReady,
			Status:  corev1.ConditionUnknown,
			Reason:  "SinkEmpty",
			Message: "Sink has resolved to empty.",
		},
	}, {
		name: "mark sink nil, secrets, webhook, then sink",
		s: func() *SampleSourceStatus {
			s := &SampleSourceStatus{}
			s.InitializeConditions()
			s.MarkSink(nil)
			s.MarkSecrets()
			s.MarkWebhookConfigured()
			s.MarkSink(apis.HTTP("example"))
			return s
		}(),
		condQuery: SampleSourceConditionReady,
		want: &apis.Condition{
			Type:   SampleSourceConditionReady,
			Status: corev1.ConditionTrue,
		},
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.s.GetCondition(test.condQuery)
			ignoreTime := cmpopts.IgnoreFields(apis.Condition{},
				"LastTransitionTime", "Severity")
			if diff := cmp.Diff(test.want, got, ignoreTime); diff != "" {
				t.Errorf("unexpected condition (-want, +got) = %v", diff)
			}
		})
	}
}

func TestSampleSource_GetGroupVersionKind(t *testing.T) {
	src := SampleSource{}
	gvk := src.GetGroupVersionKind()

	if gvk.Kind != "SampleSource" {
		t.Errorf("Should be 'SampleSource'.")
	}
}

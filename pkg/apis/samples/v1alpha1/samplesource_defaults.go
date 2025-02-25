package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

// SetDefaults mutates SampleSource.
func (s *SampleSource) SetDefaults(ctx context.Context) {
	// Add code for Mutating admission webhook.

	// example: If ServiceAccountName is unspecified, default to the "default" service account.
	if s != nil && s.Spec.ServiceAccountName == "" {
		s.Spec.ServiceAccountName = "default"
	}

	// example: If Interval is unspecified, default to "10s".
	if s != nil && s.Spec.Interval == "" {
		s.Spec.Interval = "10s"
	}

	// call SetDefaults against duckv1.Destination with a context of ObjectMeta of SampleSource.
	withNS := apis.WithinParent(ctx, s.ObjectMeta)
	s.Spec.Sink.SetDefaults(withNS)
}

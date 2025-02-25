package v1alpha1

import (
	"context"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	duckv1 "knative.dev/pkg/apis/duck/v1"

	"github.com/google/go-cmp/cmp"
)

func TestSampleSourceDefaults(t *testing.T) {
	testCases := map[string]struct {
		initial  SampleSource
		expected SampleSource
	}{
		"nil spec": {
			initial: SampleSource{},
			expected: SampleSource{
				Spec: SampleSourceSpec{
					ServiceAccountName: "default",
					Interval:           "10s",
				},
			},
		},
		"no namespace in sink reference": {
			initial: SampleSource{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "parent",
				},
				Spec: SampleSourceSpec{
					ServiceAccountName: "default",
					Interval:           "10s",
					SourceSpec: duckv1.SourceSpec{
						Sink: duckv1.Destination{
							Ref: &duckv1.KReference{},
						},
					},
				},
			},
			expected: SampleSource{
				ObjectMeta: v1.ObjectMeta{
					Namespace: "parent",
				},
				Spec: SampleSourceSpec{
					ServiceAccountName: "default",
					Interval:           "10s",
					SourceSpec: duckv1.SourceSpec{
						Sink: duckv1.Destination{
							Ref: &duckv1.KReference{
								Namespace: "parent",
							},
						},
					},
				},
			},
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			tc.initial.SetDefaults(context.TODO())
			if diff := cmp.Diff(tc.expected, tc.initial); diff != "" {
				t.Fatalf("Unexpected defaults (-want, +got): %s", diff)
			}
		})
	}
}

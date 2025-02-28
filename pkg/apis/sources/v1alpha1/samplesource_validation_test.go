package v1alpha1

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"knative.dev/pkg/webhook/resourcesemantics"

	"knative.dev/pkg/apis"
)

func TestSampleSourceValidation(t *testing.T) {
	testCases := map[string]struct {
		cr   resourcesemantics.GenericCRD
		want *apis.FieldError
	}{
		"empty sink": {
			cr: &SampleSource{
				Spec: SampleSourceSpec{},
			},
			want: func() *apis.FieldError {
				var errs *apis.FieldError
				fe := apis.ErrGeneric("expected at least one, got none", "ref", "uri").ViaField("spec.sink")
				errs = errs.Also(fe)
				return errs
			}(),
		},
	}

	for n, test := range testCases {
		t.Run(n, func(t *testing.T) {
			got := test.cr.Validate(context.Background())
			if diff := cmp.Diff(test.want.Error(), got.Error()); diff != "" {
				t.Errorf("%s: validate (-want, +got) = %v", n, diff)
			}
		})
	}
}

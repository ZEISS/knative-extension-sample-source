package v1alpha1

import (
	"context"

	"knative.dev/pkg/apis"
)

func (g *SampleSource) Validate(ctx context.Context) *apis.FieldError {
	return g.Spec.Validate(ctx).ViaField("spec")
}

func (gs *SampleSourceSpec) Validate(ctx context.Context) *apis.FieldError {
	var errs *apis.FieldError

	// TODO: there are more requirements for SampleSource. Add them here.

	// Validate sink
	errs = errs.Also(gs.Sink.Validate(ctx).ViaField("sink"))

	return errs
}

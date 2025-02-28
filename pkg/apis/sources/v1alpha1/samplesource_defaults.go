package v1alpha1

import (
	"context"
)

func (g *SampleSource) SetDefaults(ctx context.Context) {
	g.Spec.SetDefaults(ctx)
}

func (gs *SampleSourceSpec) SetDefaults(ctx context.Context) {
	// Nothing yet.
}

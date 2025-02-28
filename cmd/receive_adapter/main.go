package main

import (
	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/signals"

	sampleadapter "github.com/zeiss/zeiss/knative-extension-sample-source/pkg/adapter"
)

func main() {
	ctx := signals.NewContext()
	ctx = adapter.WithInjectorEnabled(ctx)

	adapter.MainWithContext(ctx, "samplesource", sampleadapter.NewEnvConfig, sampleadapter.NewAdapter)
}

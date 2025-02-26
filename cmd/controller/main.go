package main

import (
	// The set of controllers this controller process runs.
	"github.com/zeiss/knative-extension-sample-source/pkg/reconciler/sample"

	// This defines the shared main for injected controllers.
	"knative.dev/pkg/injection/sharedmain"
)

func main() {
	sharedmain.Main("sample-source-controller", sample.NewController)
}

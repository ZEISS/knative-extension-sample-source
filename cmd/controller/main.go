package main

import (
	"github.com/zeiss/knative-extension-sample-source/pkg/reconciler/sample"
	"knative.dev/pkg/injection/sharedmain"
)

const (
	component = "sample-controller"
)

func main() {
	sharedmain.Main(component, sample.NewController)
}

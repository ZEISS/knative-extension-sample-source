//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/ko"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/vektra/mockery/v2"
	_ "github.com/zeiss/pkg/cmd/runproc"
	_ "gotest.tools/gotestsum"
	_ "k8s.io/code-generator"
	_ "knative.dev/hack"
	_ "knative.dev/pkg/codegen/cmd/injection-gen"
	_ "mvdan.cc/gofumpt"

	// codegen: hack/generate-knative.sh
	_ "knative.dev/pkg/hack"

	_ "k8s.io/code-generator/cmd/client-gen"
	_ "k8s.io/code-generator/cmd/deepcopy-gen"
	_ "k8s.io/code-generator/cmd/defaulter-gen"
	_ "k8s.io/code-generator/cmd/informer-gen"
	_ "k8s.io/code-generator/cmd/lister-gen"
	_ "k8s.io/kube-openapi/cmd/openapi-gen"
	_ "knative.dev/pkg/codegen/cmd/injection-gen"
)

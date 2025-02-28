//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/ko"
	_ "github.com/goreleaser/goreleaser/v2"
	_ "gotest.tools/gotestsum"
	_ "knative.dev/hack"
	_ "knative.dev/pkg/hack"
	_ "mvdan.cc/gofumpt"
)

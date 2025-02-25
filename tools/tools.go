//go:build tools
// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/goreleaser/goreleaser"
	_ "github.com/vektra/mockery/v2"
	_ "github.com/zeiss/pkg/cmd/runproc"
	_ "gotest.tools/gotestsum"
	_ "knative.dev/hack"
	_ "knative.dev/pkg/hack"
	_ "mvdan.cc/gofumpt"
)

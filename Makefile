.DEFAULT_GOAL := build

-include .env
export

# Go variables
GO 					?= go
GO_RUN_TOOLS 		?= $(GO) run 
GO_TEST 			?= $(GO_RUN_TOOLS) gotest.tools/gotestsum --format pkgname
GO_RELEASER 		?= $(GO_RUN_TOOLS) github.com/goreleaser/goreleaser
GO_MOD				?= $(shell ${GO} list -m)

KO 					?= $(GO_RUN_TOOLS) github.com/google/ko

KNATIVE_SERVING		?= 1.17.0
KNATIVE_EVENTING	?= 1.17.2

.PHONY: release
release: ## Release the project.
	$(GO_RELEASER) release --clean

.PHONY: build
build: ## Build the binary file.
	$(GO_RELEASER) build --snapshot --clean

.PHONY: generate
generate: ## Generate code.
	$(GO) generate ./...

.PHONY: mocks
mocks: ## Generate mocks.
	$(GO_RUN_TOOLS) github.com/vektra/mockery/v2

.PHONY: deploy
deploy: ## Deploy the project.
	$(KO) apply -j 1 -BRf config/

.PHONY: setup
setup: ## Setup the project.
	kind create cluster --config cluster.yaml


.PHONY: fmt
fmt: ## Run go fmt against code.
	$(GO_RUN_TOOLS) mvdan.cc/gofumpt -w .

.PHONY: start
start: ## Run air live reload. Create a .air.toml file to configure.

.PHONY: vet
vet: ## Run go vet against code.
	$(GO) vet ./...

.PHONY: test
test: fmt vet ## Run tests.
	mkdir -p .test/reports
	$(GO_TEST) --junitfile .test/reports/unit-test.xml -- -race ./... -count=1 -short -cover -coverprofile .test/reports/unit-test-coverage.out

.PHONY: lint
lint: ## Run lint.
	$(GO_RUN_TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint run --timeout 5m -c .golangci.yml

.PHONY: clean
clean: ## Remove previous build.
	@rm -rf .test .dist
	@find . -type f -name '*.gen.go' -exec rm {} +
	@git checkout go.mod

.PHONY: help
help: ## Display this help screen.
	@grep -E '^[a-z.A-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
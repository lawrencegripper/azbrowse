DEV_CONTAINER_TAG:=ghcr.io/lawrencegripper/azbrowse/devcontainer:latest
DEV_CONTAINER_SNAPBASE_TAG:=ghcr.io/lawrencegripper/azbrowse/snapbase:latest
-include .env

# Used to override with richgo for colorized test output
GO_BINARY?=go

.PHONY: checks test build fail

## ----------Targets------------
## all:
##		Default action - check, test and build
all: ci

## help: 
## 		Show this help
help : Makefile
	@sed -n 's/^##//p' $<

## test-go:
## 		Run short go unit tests
test-go:
	$(GO_BINARY) test -mod=vendor ./...

## test-python
##		Run python tests, mostly used for swagger gen
test-python: swagger-update-requirements
	$(HOME)/.local/bin/pytest ./scripts/swagger_update/test_swagger_update.py

## test:
## 		Run all tests
test: test-go test-python

## checks:
##		Check/lint code
checks:
	golangci-lint run

## build: 
##		Build azbrowse binary
build:
	$(GO_BINARY) build -mod=vendor ./cmd/azbrowse

## ci: 
##		Build lint and check
ci: swagger-codegen checks test build

## debug:
##		Starts azbrowse using Delve ready for debugging from VSCode.
debug:
	dlv debug ./cmd/azbrowse --headless --listen localhost:2345 --api-version 2 -- ${ARGS}

## debug-fuzzer:
##		Starts azbrowse using Delve ready for debugging from VSCode and running the fuzzer.
debug-fuzzer:
	dlv debug ./cmd/azbrowse --headless --listen localhost:2345 --api-version 2 -- -fuzzer 5

## dcterminal 
##		Starts an interactive terminal running inside the devcontainer
dcterminal:
	docker exec -it -w /workspaces/azbrowse azbdev /bin/bash 

## fmt: 
##		Format code
fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

## run:
##		Quick command to lint and then launch azbrowse
run: install
	azbrowse --debug

## fuzz-from:
##		Runs azbrowse fuzzer which browses resource attempting to find problems
fuzz: checks install
	azbrowse --fuzzer 5

## fuzz:
##		Runs azbrowse fuzzer which browses resource attempting to find problems
## 		Accepts `node_id=/some/path/here` to start at a certain location in subscriptions
fuzz-from: checks install
	azbrowse --fuzzer 5 --navigate ${node_id}

## install: 
##		Build and install azbrowse on this machine
install:
	$(GO_BINARY) install -mod=vendor ./cmd/azbrowse

## ----------Advanced Targets------------
## docs-update:
##		Generate the docs for the command line
docs-update: install
	AZB_GEN_COMMAND_MARKDOWN=TRUE azbrowse

## swagger-update:
##		Download the latest swagger definitions for Azure services and filter to the latest versions
swagger-update: swagger-update-requirements
	python3 ./scripts/swagger_update/app.py
	
swagger-update-requirements:
	pip3 install -q -r scripts/swagger_update/requirements.txt
## swagger-codegen:
##		Generate the code needed for browse services from the swagger definitions
##		set VERBOSE=true to see full output
swagger-codegen:
	export GO111MODULE=on
	$(GO_BINARY) run -mod=vendor ./cmd/swagger-codegen/ 
	# Format the generated code
	gofmt -s -w internal/pkg/expanders/swagger-armspecs.generated.go
	gofmt -s -w internal/pkg/expanders/search.generated.go
	gofmt -s -w internal/pkg/expanders/databricks.generated.go
	# Build the generated go files to check for any go build issues
	$(GO_BINARY) build -mod=vendor internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go internal/pkg/expanders/test_utils.go
	# Test the generated code initalizes
	$(GO_BINARY) test -mod=vendor -v internal/pkg/expanders/swagger-armspecs_test.go internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go

## autocomplete-install:
## 		Add autocompletion for azbrowse to your bash prompt
autocomplete-install: install
	echo 'source <(azbrowse completion bash)' >> ~/.bashrc
	echo "Create a new shell and autocomplete will be working"

## autocomplete-test:
##		Invoke autocompletion for subscriptions and time the result
autocomplete-test: install
	@/bin/bash -c "time azbrowse __complete --subscription LG"
	@/bin/bash -c "time azbrowse __complete --navigate /"

## autocomplete-clear:
##		Invoke autocompletion for subscriptions and time the result
autocomplete-clear: 
	rm ~/.azbrowse/*

## selfupdate-test:
##		Launches AzBrowse with a low version number to allow testing of the self-update feature
selfupdate-test:
	$(GO_BINARY) install -mod=vendor -i -ldflags "-X main.version=0.0.1-testupdate" ./cmd/azbrowse 
	AZBROWSE_FORCE_UPDATE=true azbrowse

## local-ci
##		Run CI from inside a devcontainer or codespace, the build won't be published but the same process will be followed as PR builds
local-ci:
	./scripts/local-ci.sh

## devcontainer-local-ci
##     Run CI from outside a devcontainer on a box with docker and the devcontainer/cli installed https://github.com/devcontainers/cli/
devcontainer-local-ci:
	# Run 'npm install -g @devcontainers/cli' to install the CLI
	devcontainer up --workspace-folder ${PWD}
	devcontainer run-user-commands --workspace-folder ${PWD}
	devcontainer exec --workspace-folder ${PWD} ./scripts/local-ci.sh
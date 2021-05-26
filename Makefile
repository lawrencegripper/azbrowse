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
	$(GO_BINARY) test -p 1 -short ./...

## test:
## 		Run all tests
test: swagger-update-requirements terraform-hack-init
	echo ${go}
	pytest ./scripts/swagger_update/test_swagger_update.py
	$(GO_BINARY) test ./...

## checks:
##		Check/lint code
checks:
	golangci-lint run

## build: 
##		Build azbrowse binary
build:
	$(GO_BINARY) build ./cmd/azbrowse

## ci: 
##		Build lint and check
ci: swagger-codegen checks test
	$(GO_BINARY) build ./cmd/azbrowse

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
	azbrowse --fuzzer 5 --debug

## fuzz:
##		Runs azbrowse fuzzer which browses resource attempting to find problems
## 		Accepts `node_id=/some/path/here` to start at a certain location in subscriptions
fuzz-from: checks install
	azbrowse --fuzzer 5 --navigate ${node_id}

## install: 
##		Build and install azbrowse on this machine
install:
	$(GO_BINARY) install ./cmd/azbrowse

## terraform-hack-init:
##		Install terraform providers for tests
terraform-hack-init:
	./hack/init.sh

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
	$(GO_BINARY) run ./cmd/swagger-codegen/ 
	# Format the generated code
	gofmt -s -w internal/pkg/expanders/swagger-armspecs.generated.go
	gofmt -s -w internal/pkg/expanders/search.generated.go
	gofmt -s -w internal/pkg/expanders/databricks.generated.go
	# Build the generated go files to check for any go build issues
	$(GO_BINARY) build internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go internal/pkg/expanders/test_utils.go
	# Test the generated code initalizes
	$(GO_BINARY) test -v internal/pkg/expanders/swagger-armspecs_test.go internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go

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
	$(GO_BINARY) install -i -ldflags "-X main.version=0.0.1-testupdate" ./cmd/azbrowse 
	AZBROWSE_FORCE_UPDATE=true azbrowse

## devcontainer:
##		Builds the devcontainer used for VSCode and CI
devcontainer:
	@echo "Building devcontainer using tag: $(DEV_CONTAINER_TAG)"
	# Get cached layers by pulling previous version (leading dash means it's optional, will continue on failure)
	-docker pull $(DEV_CONTAINER_TAG)
	-docker pull $(DEV_CONTAINER_SNAPBASE_TAG)
	# Build the devcontainer: Hide output if it builds to keep things clean
	docker build -f ./.devcontainer/snapbase.Dockerfile ./.devcontainer --cache-from $(DEV_CONTAINER_SNAPBASE_TAG) -t $(DEV_CONTAINER_SNAPBASE_TAG) --build-arg BUILDKIT_INLINE_CACHE=1
	docker build -f ./.devcontainer/Dockerfile ./.devcontainer --cache-from $(DEV_CONTAINER_TAG) -t $(DEV_CONTAINER_TAG) --build-arg BUILDKIT_INLINE_CACHE=1

## devcontainer-push:
##		Pushes the devcontainer image for caching to speed up builds
devcontainer-push:
	./scripts/docker_login.sh
	docker push $(DEV_CONTAINER_TAG)
	docker push $(DEV_CONTAINER_SNAPBASE_TAG)

## devcontainer-integration:
##		Used for locally running integration tests
devcontainer-integration:
ifdef DEVCONTAINER
	$(error This target can only be run outside of the devcontainer as it mounts files and this fails within a devcontainer. Don't worry all it needs is docker)
endif
	@echo "Using tag: $(DEV_CONTAINER_TAG)"
	@docker run -v ${PWD}:${PWD} \
		--entrypoint /bin/bash \
		--workdir ${PWD} \
		-t $(DEV_CONTAINER_TAG) \
		-f ${PWD}/scripts/ci_integration_tests.sh

## devcontainer-release:
##		Used by the build to create, test and publish
devcontainer-release:
ifdef DEVCONTAINER
	$(error This target can only be run outside of the devcontainer as it mounts files and this fails within a devcontainer. Don't worry all it needs is docker)
endif
	@echo "Using tag: $(DEV_CONTAINER_TAG)"
	# Note command mirrors required envs from host into container. Using '@' to avoid showing values in CI output.
	# Docs
	# - Build/CI/PR etc for controlling build vs build and publish release
	# - "-v var/rundocker.socket" to allow docker builds via mounted docker socket
	# - "privileged" and "--device" to allow fuse testing
	@docker run -v ${PWD}:${PWD} \
		-e "TERM=xterm-256color" \
		-e BUILD_NUMBER="${BUILD_NUMBER}" \
		-e IS_CI="${IS_CI}" \
		-e IS_PR="${IS_PR}" \
		-e BRANCH="${BRANCH}" \
		-e GITHUB_TOKEN="${GITHUB_TOKEN}" \
		-e DOCKER_USERNAME="${DOCKER_USERNAME}" \
		-e DOCKER_PASSWORD="${DOCKER_PASSWORD}" \
		-e DEV_CONTAINER_TAG="$(DEV_CONTAINER_TAG)" \
		-e SNAPCRAFT_LOGIN="$(SNAPCRAFT_LOGIN)" \
		-v /var/run/docker.sock:/var/run/docker.sock \
		--privileged \
		--device /dev/fuse \
		--workdir "${PWD}" \
		$(DEV_CONTAINER_TAG) \
		ruby ${PWD}/scripts/release.rb
		

asfs-build:
	$(GO_BINARY) build ./cmd/azfs

azfs-run:
	-@fusermount -u /mnt/azfs
	mkdir -p /mnt/azfs
	$(GO_BINARY) run ./cmd/azfs --mount /mnt/azfs --enableEdit

# azfs-test:
# 		Tests the azfs filesystem with unit tests and mocked api
#
# A '.env' file like the following is required 
#
# > TESTSUB=yoursubhere
# > TESTRESOURCE=/rghere/resourcehere
#
# The resource specified should have a value of 'replaceme' as a tag
azfs-test:
	$(GO_BINARY) test -count=1 -timeout 30s -short ./internal/pkg/filesystem

# azfs-integration:
# 		Tests the azfs filesystem with integration tests
#
# A '.env' file like the following is required 
#
# > TESTSUB=yoursubhere
# > TESTRESOURCE=/rghere/resourcehere
#
# The resource specified should have a value of 'replaceme' as a tag
azfs-integration:
	TESTSUB=${TESTSUB} TESTRESOURCE=${TESTRESOURCE} $(GO_BINARY) test -v -count=1 -timeout 30s ./internal/pkg/filesystem

# add-sample-queries:
# 		Copy some sample queries to the correct location. Used for testing the feature.
add-sample-queries:
	mkdir -p ~/.azbrowse/
	cp ./samplequeries/* ~/.azbrowse/
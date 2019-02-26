.PHONY: dep checks test build
all: dep checks test build

setup:
	./scripts/install_dev_tools.sh

dep:
	dep ensure -v --vendor-only

test:
	go test -v -short ./...

integration:
	go test ./...

build: test dep checks
	go build ./cmd/azbrowse

install:
	go install ./cmd/azbrowse

checks:
	# Moving to a more strict linting set from commit: 0b4b09bfbf67e533d3b831d29e14c8250e2b53ca
	golangci-lint run --new-from-rev=0b4b09bfbf67e533d3b831d29e14c8250e2b53ca

ci-docker:
	docker run -it -e BUILD_NUMBER=${TRAVIS_BUILD_NUMBER} -v $(CURDIR):/go/src/github.com/lawrencegripper/azbrowse golang:1.11 bash -f /go/src/github.com/lawrencegripper/azbrowse/scripts/release.sh

swagger-codegen:
	go run cmd/swagger-codegen/main.go --output-file ./internal/pkg/handlers/swagger.generated.go 
	# Format the generated code
	gofmt -s -w internal/pkg/handlers/swagger.generated.go
	# Build the generated go files to check for any go build issues
	go build internal/pkg/handlers/swagger.generated.go internal/pkg/handlers/swagger.go internal/pkg/handlers/types.go 
	# Test the generated code initalizes
	go test -v internal/pkg/handlers/swagger_test.go internal/pkg/handlers/swagger.generated.go internal/pkg/handlers/swagger.go internal/pkg/handlers/types.go
	
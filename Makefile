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

build: dep swagger-codegen test checks 
	go build ./cmd/azbrowse

install:
	go install ./cmd/azbrowse

checks:
	golangci-lint run

ci-docker:
	docker run -it -e BUILD_NUMBER=999-localci -v /var/run/docker.sock:/var/run/docker.sock -v $(CURDIR):/go/src/github.com/lawrencegripper/azbrowse golang:1.11.5 bash -f /go/src/github.com/lawrencegripper/azbrowse/scripts/ci.sh

swagger-codegen:
	go run cmd/swagger-codegen/main.go --output-file ./internal/pkg/handlers/swagger.generated.go 
	# Format the generated code
	gofmt -s -w internal/pkg/handlers/swagger.generated.go
	# Build the generated go files to check for any go build issues
	go build internal/pkg/handlers/swagger.generated.go internal/pkg/handlers/swagger.go internal/pkg/handlers/types.go 
	# Test the generated code initalizes
	go test -v internal/pkg/handlers/swagger_test.go internal/pkg/handlers/swagger.generated.go internal/pkg/handlers/swagger.go internal/pkg/handlers/types.go

debug:
	go build ./cmd/azbrowse &&  dlv exec ./azbrowse --headless --listen localhost:2345 --api-version 2

run: checks install
	azbrowse
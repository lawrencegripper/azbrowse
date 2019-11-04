.PHONY: checks test build
all: checks test build

setup:
	./scripts/install_ci_tools.sh

test:
	GO111MODULE=on go test -v -short ./...

integration:
	GO111MODULE=on go test -v -count=1 ./...

build: swagger-codegen test checks 
	GO111MODULE=on go build ./cmd/azbrowse

fmt:
	find . -name '*.go' | grep -v vendor | xargs gofmt -s -w

install:
	GO111MODULE=on go install ./cmd/azbrowse

checks:
	GO111MODULE=on golangci-lint run

ci-docker:
	docker run -it -e BUILD_NUMBER=999-localci -v /var/run/docker.sock:/var/run/docker.sock -v $(CURDIR):/go/src/github.com/lawrencegripper/azbrowse golang:1.11.5 bash -f /go/src/github.com/lawrencegripper/azbrowse/scripts/ci.sh

swagger-update:
	./scripts/update-swagger.sh
	
swagger-codegen:
	export GO111MODULE=on
	go run ./cmd/swagger-codegen/ --output-file ./internal/pkg/expanders/swagger-armspecs.generated.go 
	go run ./cmd/swagger-codegen/ 
	# Format the generated code
	gofmt -s -w internal/pkg/expanders/swagger-armspecs.generated.go
	gofmt -s -w internal/pkg/expanders/search.generated.go
	# Build the generated go files to check for any go build issues
	go build internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go 
	# Test the generated code initalizes
	go test -v internal/pkg/expanders/swagger-armspecs_test.go internal/pkg/expanders/swagger-armspecs.generated.go internal/pkg/expanders/swagger-armspecs.go internal/pkg/expanders/swagger.go internal/pkg/expanders/types.go

debug:
	GO111MODULE=on go build ./cmd/azbrowse &&  dlv exec ./azbrowse --headless --listen localhost:2345 --api-version 2

run: checks install
	azbrowse
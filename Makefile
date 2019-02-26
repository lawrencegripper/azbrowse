.PHONY: dep checks test build
all: dep checks test build

dep:
	dep ensure -v --vendor-only

test:
	go test -short ./...

integration:
	go test ./...

build: test dep checks
	go build ./cmd/azbrowse

install: 
	go install ./cmd/azbrowse

checks:
	gometalinter --vendor --disable-all --enable=vet --enable=gofmt --enable=golint --enable=deadcode --enable=varcheck --enable=structcheck --enable=misspell --deadline=15m ./...

ci-docker:
	docker run -it -e BUILD_NUMBER=${TRAVIS_BUILD_NUMBER} -v $(CURDIR):/go/src/github.com/lawrencegripper/azbrowse golang:1.10 bash -f /go/src/github.com/lawrencegripper/azbrowse/scripts/release.sh

swagger-codegen:
	go run cmd/swagger-codegen/main.go --output-file ./internal/pkg/handlers/swagger.generated.go 
	# Format the generated code
	gofmt -s -w internal/pkg/handlers/swagger.generated.go
	# Build the generated go files to check for any go build issues
	go build internal/pkg/handlers/swagger.generated.go internal/pkg/handlers/swagger.go internal/pkg/handlers/types.go 
	
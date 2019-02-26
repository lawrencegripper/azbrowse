#/bin/bash
set -e

# Download and install dep (https://github.com/golang/dep)
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Download and install golangci-lint (https://github.com/golangci/golangci-lint)
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin
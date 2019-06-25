#!/bin/bash
set -e

# Download and install golangci-lint (https://github.com/golangci/golangci-lint)
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(go env GOPATH)/bin
#!/bin/bash
set -e

# Download and install golangci-lint (https://github.com/golangci/golangci-lint)
curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sed 's/tar -/tar --no-same-owner -/g' | sh -s -- -b $(go env GOPATH)/bin
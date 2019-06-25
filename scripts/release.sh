#!/usr/bin/env bash
set -e
cd `dirname $0`

if [ -z "$BUILD_NUMBER" ]; then
    echo "Envvar 'BUILD_NUMBER' must be set for this script to work correctly. When building locally for debugging/testing this script is not needed use 'go build' instead."
    exit 1
fi 

git tag -f v1.1.$BUILD_NUMBER
export GOVERSION=$(go version)
export GO111MODULE=on

echo "->Move to root directory"
cd ../

echo "->Use make build to codegen, lint and check"
make build
if [[ $(git diff --stat) != '' ]]; then
  echo "Ditry GIT: Failing as swagger-generated caused changes, please run `make swagger-update` and `make swagger-generate` and commit changes for build to pass"
  exit 1
else
  echo '`swagger-gen` ran and no changes detected in code: Success'
fi

echo "->Run go releaser"
if [ -z ${PUBLISH} ]; then
  echo "Running with --skip-publish as PUBLISH not set"
  curl -sL https://git.io/goreleaser | bash -s -- --skip-publish --rm-dist
else 
  echo "Publishing"
  curl -sL https://git.io/goreleaser | bash
fi
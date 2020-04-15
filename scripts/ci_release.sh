#!/usr/bin/env bash
set -e

cd `dirname $0`

echo ""
echo "-------> RELEASE PROCESS"
echo ""
# Fail if build number not set
if [ -z "$BUILD_NUMBER" ]; then
    echo "Envvar 'BUILD_NUMBER' must be set for this script to work correctly. When building locally for debugging/testing this script is not needed use 'go build' instead."
    exit 1
fi 

# If running inside CI login to docker
if [ -z ${IS_CI} ]; then
  echo "Not running in circle, skipping CI setup"
else 
  echo "Publishing"
  if [ -z $IS_PR ] && [[ $BRANCH == "refs/heads/master" ]]; then
    export PUBLISH=true
    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
    echo "On master setting PUBLISH=true"
  else 
    echo "Skipping publish as is from PR: $PR_NUMBER or not 'refs/heads/master' BRANCH: $BRANCH"
  fi
fi

# Set version for release (picked up later by goreleaser)
git tag -f v1.2.$BUILD_NUMBER

export GOVERSION=$(go version)

echo "->Move to root directory"
cd ../

echo "->Use make build to codegen, lint and check"
make build

if [[ $(git diff --stat) != '' ]]; then
  echo "--> Ditry GIT: Failing as swagger-generated caused changes, please run 'make swagger-update' and 'make swagger-generate' and commit changes for build to pass"
  git status
  sleep 1
  exit 1
else
  echo "'swagger-gen' ran and no changes detected in code: Success"
fi

echo "->Run go releaser"
if [ -z ${PUBLISH} ]; then
  echo "Running with --skip-publish as PUBLISH not set"
  goreleaser --skip-publish --rm-dist
else 
  echo "Publishing"
  goreleaser
  echo "Pushing update to devcontainer image to speed up next build"
  docker push "$DEV_CONTAINER_TAG"
fi
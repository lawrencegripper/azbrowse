#!/usr/bin/env bash
set -e

cd `dirname $0`

print_header () {
  echo ""
  echo "$(tput setaf 12)$(tput bold)------------------------------------------------------- $(tput sgr0)"
  echo "$(tput setaf 12)$(tput bold)--------------->$(tput sgr0) $1 "
  echo "$(tput setaf 12)$(tput bold)------------------------------------------------------- $(tput sgr0)"
  echo ""
}

print_header "RELEASE PROCESS"
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
    echo "On master setting PUBLISH=true"
    export PUBLISH=true
    
    echo "Docker login"
    echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin

    echo "Snapcraft Login"
    echo $SNAPCRAFT_LOGIN | base64 -d > snap.login
    snapcraft login --with snap.login
    # cleanup login file
    rm snap.login
  else 
    echo "Skipping publish as is from PR: $PR_NUMBER or not 'refs/heads/master' BRANCH: $BRANCH"
  fi
fi

# Set version for release (picked up later by goreleaser)
git tag -f v2.0.$BUILD_NUMBER

export GOVERSION=$(go version)

print_header "Use make build to codegen, lint and check"

cd ../
GO_BINARY=richgo make build

print_header "Check codegen results haven't changed checkedin code"
if [[ $(git diff --stat) != '' ]]; then
  echo "--> Ditry GIT: Failing as swagger-generated caused changes, please run 'make swagger-update' and 'make swagger-generate' and commit changes for build to pass"
  git status
  sleep 1
  exit 1
else
  echo "'swagger-gen' ran and no changes detected in code: Success"
fi

make docs-update
if [[ $(git diff --stat) != '' ]]; then
  echo "--> Ditry GIT: Commandline args changed but 'make docs-update' hasn't been run. Please run it and commit changes."
  git status
  sleep 1
  exit 1
fi

print_header "Run Integration tests on fake display"

Xvfb :99 -ac -screen 0 "$XVFB_RES" -nolisten tcp $XVFB_ARGS &
XVFB_PROC=$!
sleep 1
export DISPLAY=:99
exitcodefile=$(mktemp)
logfile=$(mktemp)

echo "Run make integration in Xterm"
xterm -e sh -c 'make integration > '"$logfile"'; echo $? > '"$exitcodefile"
echo "Tests finished"
cat $exitcodefile
exitcode=$(cat "$exitcodefile")
rm "$exitcodefile"

if [[ $exitcode == "0" ]]; then
  cat "$logfile"

  echo "Tests passed"
else
  cat "$logfile"
  go version
  
  echo "Tests returned exit code: $exitcode"
  echo "Tests FAILED. Logs:"
  exit 1
fi 
kill $XVFB_PROC

print_header "Run go releaser"

# Workaround concurrency bug which intermittently causes failure building snap.
mkdir -p $HOME/.cache/snapcraft/download
mkdir -p $HOME/.cache/snapcraft/stage-packages

if [ -z ${PUBLISH} ]; then
  echo "Running with --skip-publish as PUBLISH not set"
  goreleaser --skip-publish --rm-dist
else 
  echo "Publishing"
  goreleaser
  echo "Pushing update to devcontainer image to speed up next build"
  docker push "$DEV_CONTAINER_TAG"
fi
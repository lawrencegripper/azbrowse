#!/bin/bash
set -e
BUILD_NUMBER=1 IS_CI=true BRANCH=$(git branch --show-current) CMD="ruby ${PWD}/scripts/release.rb" make devcontainer-run
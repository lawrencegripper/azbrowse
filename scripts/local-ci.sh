#!/bin/bash
set -e
export BUILD_NUMBER=1
export IS_CI=true
export BRANCH=$(git branch --show-current) 
ruby ${PWD}/scripts/release.rb
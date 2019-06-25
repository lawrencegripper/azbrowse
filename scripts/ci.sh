#!/usr/bin/env bash
set -e
cd `dirname $0`

echo "->Installing tools"
bash -f ./install_ci_tools.sh
bash -f ./install_release_tools.sh

if [ -z ${CIRCLECI} ]; then
  echo "Not running in circle, skipping cirlce setup"
else 
  echo "Publishing"
  if [ -z $CIRCLE_PR_NUMBER ] && [[ $CIRCLE_BRANCH == "master" ]]; then
    export PUBLISH=true
    docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD
    echo "On master setting PUBLISH=true"
  else 
    echo "Skipping publish as is from PR: $CIRCLE_PR_NUMBER or not master BRANCH: $CIRCLE_BRANCH"
  fi
fi

bash -f ./release.sh

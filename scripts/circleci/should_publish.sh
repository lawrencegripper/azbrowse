#!/usr/bin/env bash
set -e

if [ -z $CIRCLE_PR_NUMBER ] && [[ $CIRCLE_BRANCH == "master" ]]; then
    'export PUBLISH=true' >> $BASH_ENV
    echo "On master setting PUBLISH=true"
else 
    echo "Skipping publish as is from PR: $CIRCLE_PR_NUMBER or not master BRANCH: $CIRCLE_BRANCH"
fi
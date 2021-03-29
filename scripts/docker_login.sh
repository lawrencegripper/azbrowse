#!/bin/bash
set -e

# If running inside CI login to docker
if [ -z ${IS_CI} ]; then
    echo "Not running in CI env so skipping docker login"
else 
    echo "Docker login:"
    echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin
fi
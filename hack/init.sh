#!/bin/bash

SCRIPT=$(readlink -f $0)
BASE_DIR=`dirname ${SCRIPT}`

pushd "$BASE_DIR"

terraform init
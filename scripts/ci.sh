#!/usr/bin/env bash
set -e
cd `dirname $0`

bash -f ./install_release_tools.sh
bash -f ./release.sh
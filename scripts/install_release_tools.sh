#/bin/bash
set -e

# Install go releaser
if [ "$(uname)" == "Darwin" ]; then
    echo "Installing go releaser from brew"
    brew install goreleaser/tap/goreleaser

elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
    echo "Installing go releaser from Snap"
    snap install goreleaser --classic
fi
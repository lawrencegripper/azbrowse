#/bin/bash
set -e

apt-get install rpm snapd
snap install snapcraft --classic
snap install goreleaser --classic
#!/bin/bash
set -e
cd "$(dirname "$0")"
export DEBIAN_FRONTEND=noninteractive
# Install xterm and server config for use in testing in CI
apt update
apt install xorg openbox xserver-xorg-video-dummy -y
apt install xterm -y

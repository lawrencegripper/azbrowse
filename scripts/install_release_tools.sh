#!/bin/bash
set -e
echo "->Installing release tools"
apt-get update
apt-get install -y rpm

echo "->Install Docker"
bash -c "cd /tmp && curl -fsSLO https://download.docker.com/linux/static/stable/x86_64/docker-18.06.3-ce.tgz && tar --strip-components=1 -xvzf docker-18.06.3-ce.tgz -C /usr/local/bin"
#/bin/bash
set -e

# Download and install dep (https://github.com/golang/dep)
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# Download and install gometalinter (https://github.com/alecthomas/gometalinter)
go get -u github.com/alecthomas/gometalinter
gometalinter --install
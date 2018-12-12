#/bin/bash
set -e

# Download and install dep
curl -sI https://github.com/golang/dep/releases/latest | grep -Fi Location  | tr -d '\r' | sed "s/tag/download/g" | awk -F " " '{ print $2 "/dep-linux-amd64"}' | wget --output-document=$GOPATH/bin/dep -i -
chmod +x $GOPATH/bin/dep
# Install linters 
go get -u github.com/alecthomas/gometalinter
gometalinter --install
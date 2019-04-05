#/bin/bash
set -e

# Install linters 
go get -u github.com/alecthomas/gometalinter
gometalinter --install
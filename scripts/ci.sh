docker run -it -e BUILD_NUMBER=$TRAVIS_BUILD_NUMBER -v $PWD:/go/src/github.com/lawrencegripper/azbrowse golang:1.10 bash -f /go/src/github.com/lawrencegripper/azbrowse/scripts/build.sh

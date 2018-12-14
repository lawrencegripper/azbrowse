#!/usr/bin/env bash
set -e


cd `dirname $0`
bash -f ./install.sh
cd ../

dep ensure 
gometalinter --vendor --disable-all --enable=vet --enable=gofmt --enable=golint --enable=deadcode --enable=varcheck --enable=structcheck --enable=misspell --deadline=15m ./...

VERSION="1.0.$TRAVIS_BUILD_NUMBER"
GIT_COMMIT=$(git rev-parse HEAD)
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GO_VERSION=$(go version | awk '{print $3}')

echo "Building version: $VERSION"

platforms=("linux/amd64" "windows/amd64" "windows/386" "darwin/amd64" "linux/386" "linux/arm")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    CGO_ENABLED=0 
    output_name='./bin/azbrowse-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi  
    echo "Building for $GOOS $GOARCH..."

    GOOS=$GOOS GOARCH=$GOARCH go build --ldflags "-w -s \
        -X github.com/lawrencegripper/azbrowse/version.BuildDataVersion=${VERSION} \
        -X github.com/lawrencegripper/azbrowse/version.BuildDataGitCommit=${GIT_COMMIT} \
        -X github.com/lawrencegripper/azbrowse/version.BuildDataBuildDate=${BUILD_DATE} \
        -X github.com/lawrencegripper/azbrowse/version.BuildDataGoVersion=${GO_VERSION}" \
        -a -installsuffix cgo -o $output_name

    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done

echo "Completed builds, for output see ./bin"

all: test build-linux build-darwin build-windows

build-linux:
	GOARCH=amd64 GOOS=linux go build -o bin/linux/amd64/armclient .

build-darwin:
	GOARCH=amd64 GOOS=darwin go build -o bin/darwin/amd64/armclient .

build-windows:
	GOARCH=amd64 GOOS=windows go build -o bin/windows/amd64/armclient.exe .

test:
	go test -v

clean:
	rm -rf dist/
	rm armclient
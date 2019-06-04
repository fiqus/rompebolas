all: build

setup:
	go get -u github.com/shurcooL/vfsgen

fmt:
	go fmt ./...

build:
	cd data && go generate -tags dev && cd -
	go build -ldflags '-X main.version=1.0' ./cmd/rompebolas/

run:
	go run ./cmd/rompebolas/

clean:
	go clean
	rm -f data/assets_vfsdata.go
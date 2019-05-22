lint:
	golint ./...

build:
	mkdir -p bin
	go build -o bin/lua main.go

test-all:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...


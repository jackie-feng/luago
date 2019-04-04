lint:
	golint ./...

build:
	mkdir -p bin
	go build -o bin/lua main.go

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...


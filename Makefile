all: build

default: build

build:
	go build -o bin/combinoctocat combinoctocat.go

vet:
	go vet ./...

test:
	go test -p 1 -cover -timeout 30s ./... && make vet

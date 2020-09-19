all: build

default: build

build:
	go build -o bin/combinoctocat combinoctocat.go

SHELL := bash
# VERSION := 'API version\:$(shell date -u +%Y%m%d.%H%M%S)\($(shell git rev-parse --short HEAD)\)'
VERSION := ($(shell git rev-parse --short HEAD))

export VERSION;

build:
	go build -v -ldflags "-X github.com/ivan-iver/hermes/lib.hash=${VERSION}" -o bin/hermes github.com/ivan-iver/hermes
	@cp templates/app.conf bin/

clean:
	rm -r bin/*

init:
	mkdir -p bin log
	go get -v -u ./...


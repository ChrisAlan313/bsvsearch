.Default_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run
.PHONY:lint

test: lint
	go test
.PHONY:test

run: lint
	go run bsvsearch.go bible.go specification.go server.go
.PHONY:run

build: lint
	go build bsvsearch.go bible.go specification.go server.go
.PHONY:build

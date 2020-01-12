.DEFAULT_GOAL := all

.PHONY: generate
generate:
	go generate ./...

.PHONY: clean
clean:
	go clean ./...
	rm -rf test/mocks/*

.PHONY: lint
lint:
	golint ./...

.PHONY: test-server
test-server:
	go test -v ./...

.PHONY: test-client
.ONESHELL:
test-client:
	go run cmd/main.go &
	PID=$$!
	cd test/client
	npm install
	npm test
	kill $$PID || true

.PHONY: test
test: test-server test-client

.PHONY: run
run:
	go run cmd/main.go

.PHONY: all
all: clean generate test

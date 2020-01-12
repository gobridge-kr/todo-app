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

ifneq ($(origin HEROKU_APP_NAME),undefined)
	OPTS := -a $(HEROKU_APP_NAME)
endif
.PHONY: deploy
deploy:
	heroku container:login
	heroku container:push $(OPTS) web
	heroku container:release $(OPTS) web

.PHONY: all
all: clean generate test

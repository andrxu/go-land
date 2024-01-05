SOURCE := ./...
MICROSERVICE := go-land
CONTAINER_IMAGE ?= andrxu/$(MICROSERVICE):latest

.DEFAULT_GOAL := build

GO=go

build:
	$(GO) build -v -o $(MICROSERVICE)
	$(GO) test -v -c -tags=functional -o $(MICROSERVICE)-functional.test
.PHONY: build

run: build
	$(GO) run main.go

lint:
	$(GO) vet $(SOURCE)
	test -z $(shell go fmt $(SOURCE))
.PHONY: lint

test: run
.PHONY: test

unit-test:
	$(GO) test -coverprofile=coverage.out $(SOURCE) -count=1
.PHONY: unit-test

docker-build:
	DOCKER_BUILDKIT=1 docker build --tag $(CONTAINER_IMAGE) .
.PHONY: docker-build 

SOURCE := ./...

GO=GOPRIVATE=github.com/andrxu/* go

build:
	$(GO) build -v $(SOURCE)
.PHONY: build

run:
	$(GO) run main.go

lint:
	$(GO) vet $(SOURCE)
	test -z $(shell go fmt $(SOURCE))
.PHONY: lint

unit-test:
	$(GO) test -coverprofile=coverage.out $(SOURCE) -count=1
.PHONY: unit-test



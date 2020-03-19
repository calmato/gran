GOLANGCI_VERSION := 1.22.2
PACKAGES := $(shell go list ./... | grep -v mock)

.PHONY: setup
setup:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: build
build:
	go build -o app ./cmd

.PHONY: fmt
fmt:
	! gofmt -d -s . | grep '^'

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	./bin/golangci-lint run -c .golangci.yml ./...

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: build
build:
	go build -o app ./cmd

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	./bin/golangci-lint run --enable-all ./...

.PHONY: test
test:
	go test -v -cover ./...

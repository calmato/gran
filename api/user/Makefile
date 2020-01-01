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
	go vet ./...

.PHONY: lint
lint:
	./bin/golangci-lint run --enable-all ./...

.PHONY: test
test:
	go test -v -cover ./...

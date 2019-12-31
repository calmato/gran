.PHONY: run
run:
	go run ./cmd/main.go

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint ./...

.PHONY: test
test:
	go test -v -cover ./...

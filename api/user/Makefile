GOLANGCI_VERSION := 1.22.2
REVEWDOG_VERSION := 0.9.17

.PHONY: setup
setup:
	wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v${GOLANGCI_VERSION}
	wget -O - -q https://raw.githubusercontent.com/reviewdog/reviewdog/master/install.sh | sh -s v${REVEWDOG_VERSION}
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

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
	./bin/golangci-lint run -c .golangci.yml ./...

.PHONY: test
test:
	go test -v -cover ./internal/...

# REVIEWDOG_ARG
# * local -> -diff='git diff master'
# * pull request -> -reporter='github-pr-review'
.PHONY: review
review:
	./bin/golangci-lint run -c .golangci.yml --out-format line-number ./... \
	| ./bin/reviewdog -f=golangci-lint --name=golangci-lint ${REVIEWDOG_ARG}

.PHONY: build-mocks
build-mocks:
	./bin/build-mocks.sh

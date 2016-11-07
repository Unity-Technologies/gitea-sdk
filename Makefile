IMPORT := github.com/go-gitea/go-sdk

PACKAGES ?= $(shell go list ./... | grep -v /vendor/)
GENERATE ?= github.com/go-gitea/go-sdk/gitea

.PHONY: all
all: clean test build

.PHONY: clean
clean:
	go clean -i ./...

generate:
	@which mockery > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/vektra/mockery/...; \
	fi
	go generate $(GENERATE)

.PHONY: fmt
fmt:
	go fmt $(PACKAGES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	@which golint > /dev/null; if [ $$? -ne 0 ]; then \
		go get -u github.com/golang/lint/golint; \
	fi
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: test
test:
	for PKG in $(PACKAGES); do go test -cover -coverprofile $$GOPATH/src/$$PKG/coverage.out $$PKG || exit 1; done;

.PHONY: build
build:
	go build ./gitea

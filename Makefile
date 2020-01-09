GO ?= go

.PHONY: all
all: clean test build

.PHONY: clean
clean:
	$(GO) clean -i ./...

.PHONY: fmt
fmt:
	find . -name "*.go" -type f ! -path "./vendor/*" ! -path "./benchmark/*" | xargs gofmt -s -w

.PHONY: vet
vet:
	cd gitea && $(GO) vet ./...

.PHONY: lint
lint:
	@echo 'make lint is depricated. Use "make revive" if you want to use the old lint tool, or "make golangci-lint" to run a complete code check.'

.PHONY: revive
revive:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml -exclude=./vendor/... ./... || exit 1

.PHONY: test
test:
	cd gitea && $(GO) test -cover -coverprofile coverage.out

.PHONY: bench
bench:
	cd gitea && $(GO) test -run=XXXXXX -benchtime=10s -bench=. || exit 1

.PHONY: build
build:
	cd gitea && $(GO) build

.PHONY: golangci-lint
golangci-lint:
	@hash golangci-lint > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		export BINARY="golangci-lint"; \
		curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s -- -b $(GOPATH)/bin v1.22.2; \
	fi
	golangci-lint run --timeout 5m
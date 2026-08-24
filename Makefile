GO  ?= go
BIN := gottage

.DEFAULT_GOAL := help
.PHONY: check fmt fmt-check vet tidy-check test test-race lint build run hooks clean help

## check      full local gate; run before every commit
check: fmt-check vet tidy-check test

## fmt        rewrite Go files with gofmt
fmt:
	$(GO) fmt ./...

## fmt-check  fail if any Go file is unformatted
fmt-check:
	@out="$$(gofmt -l .)"; \
	if [ -n "$$out" ]; then \
	  echo "gofmt: unformatted file(s):"; \
	  printf '%s\n' "$$out" | sed 's/^/  /'; \
	  echo "fix with: make fmt"; \
	  exit 1; \
	fi

## vet        go vet ./...
vet:
	$(GO) vet ./...

## tidy-check fail if go.mod / go.sum are not tidy
tidy-check:
	@$(GO) mod tidy -diff || { echo "not tidy - run: $(GO) mod tidy"; exit 1; }

## test       run the suite (~21s; pkg/art alone is ~17s)
test:
	$(GO) test ./... -count=1

## test-race  race detector (slow: pkg/art is ~222s under -race)
test-race:
	$(GO) test -race ./... -count=1

## lint       golangci-lint; optional, NOT part of `check`
lint:
	@command -v golangci-lint >/dev/null 2>&1 || { \
	  echo "golangci-lint is not installed."; \
	  echo "install with: brew install golangci-lint"; \
	  echo "(it is optional - 'make check' does not require it)"; \
	  exit 1; \
	}
	golangci-lint run

## build      build ./gottage (gitignored)
build:
	$(GO) build -o $(BIN) .

## run        go run .
run:
	$(GO) run .

## hooks      activate the committed pre-commit hook in this clone
hooks:
	@git config core.hooksPath hooks
	@echo "core.hooksPath=hooks - pre-commit hook active"

## clean      remove build outputs
clean:
	@rm -f $(BIN) coverage.out coverage.html
	@echo "cleaned"

## help       list targets
help:
	@echo "targets:"
	@grep -E '^## ' $(MAKEFILE_LIST) | sed 's/^## /  /'

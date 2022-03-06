GO ?= go

.PHONY: test
test:
	$(GO) test ./...

.PHONY: build
build:
	$(GO) build ./...

.PHONY: fmt
fmt:
	$(GO) fmt ./...

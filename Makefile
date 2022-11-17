ROOT_PKG ?= "github.com/pthomison/go-reddit"
LIST_PKG := $(shell go list $(ROOT_PKG)/...)

# Tests
TEST_TIMEOUT ?= 20

.DEFAULT_GOAL := usage

# Print colorized log
define log
	echo "\n\033[1;32m--- [$(@)] $(1) ---\033[0m\n"
endef

.PHONY: all
all: lint fmt vet test test-coverage

.PHONY: usage
usage:
	@echo "make [all|fmt|vet|lint|test|test-coverage]"

.PHONY: fmt
fmt:
	@$(call log,"Running formatter")
	@go fmt $(LIST_PKG)

.PHONY: vet
vet:
	@$(call log,"Running vet")
	@go vet -all $(LIST_PKG)

.PHONY: lint
lint:
	@$(call log,"Running linter")
	@staticcheck $(LIST_PKG)

.PHONY: test
test: fmt vet lint
	@$(call log,"Running tests")
	@go test -v -race -short -timeout $(TEST_TIMEOUT)s $(ARGS) $(LIST_PKG)

.PHONY: test-coverage
test-coverage: fmt vet lint
	@$(call log,"Running tests with coverage")
	@go test -v -race -short -timeout $(TEST_TIMEOUT)s $(ARGS) -coverprofile=coverage.out $(LIST_PKG)
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out -o coverage.html

# pthomison release logic

CURRENT_TAG=$(shell git tag --sort version:refname | tail -n 1)

first_release:
	git tag v0.0.1
	git push origin v0.0.1

rev_release:
	git tag $(shell exoskeleton rev -i $(CURRENT_TAG))
	git push origin $(shell exoskeleton rev -i $(CURRENT_TAG))

release:
	@if [[ "$(CURRENT_TAG)" == "" ]]; then $(MAKE) first_release; else $(MAKE) rev_release; fi





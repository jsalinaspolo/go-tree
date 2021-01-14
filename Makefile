#!/usr/bin/env make
VERSION := $(shell go run .)

SHELL := /bin/bash
.SHELLFLAGS = -ec

.PHONY: run
run:
	go run main.go

.PHONY: install
install:
	go install

.PHONY: fmt
fmt:
	go fmt ./... -v

.PHONY: lint
lint:
	@golangci-lint run --max-same-issues=0 -v

.PHONY: test
test:
	go test -cover ./...

.PHONY: release-cli-test
release-cli-test:
	curl -sL https://git.io/goreleaser | bash -s -- --rm-dist --snapshot --skip-publish


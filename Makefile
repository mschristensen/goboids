SHELL:=/bin/bash

cmd=goboids

.PHONY: build
build:
	@go build -o ./bin/$(cmd) ./cmd/$(cmd)/main.go

.PHONY: run
run:
	@./bin/$(cmd)
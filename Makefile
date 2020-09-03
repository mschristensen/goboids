SHELL:=/bin/bash

cmd=goboids

.PHONY: build
build:
	@go build -o ./bin/$(cmd) ./cmd/$(cmd)/main.go

.PHONY: run
run:
	@./bin/$(cmd)

.PHONY: run_watch
run_watch:
	@git ls-files | entr -r go run ./cmd/$(cmd)/main.go
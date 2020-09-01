SHELL:=/bin/bash

.PHONY: build
build:
	@go build -o ./bin/goboids ./main.go

.PHONY: run
run:
	@./bin/goboids
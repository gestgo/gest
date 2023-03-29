# SHELL=/bin/bash -e -o pipefail
# PWD = $(shell pwd)
fmt: ## Formats all code with go fmt
	@go fmt ./...

run: fmt ## Run a controller from your host
	@go run ./src/main.go
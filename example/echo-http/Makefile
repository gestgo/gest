# SHELL=/bin/bash -e -o pipefail
# PWD = $(shell pwd)
base_path?=github.com/gestgo/gest/example/src/module
#module_name?= noName
fmt: ## Formats all code with go fmt
	@go fmt ./...

run: fmt ## Run a controller from your host
	@go run ./src/main.go
swag: fmt ## Run a controller from your host
	@ ./cmd/swagger init --dir ./src/
gen_module:
	@ cookiecutter   /home/phongthien/Desktop/start-up/gest/template/module  name=$(name) base_path=$(base_path) --no-input --output-dir ./src/module/test

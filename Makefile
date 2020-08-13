.PHONY: lint setup

lint: ## Linting the codebase
	golint -set_exit_status ./...

run-example: ## Run the example
	go run _examples/main.go _examples/elasticsearch_recipe.go

setup: ## Setup modules
	go get -u golang.org/x/lint/golint

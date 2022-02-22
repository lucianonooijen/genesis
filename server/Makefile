help: ## Shows all commands
	@echo 'All Makefile commands:'
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: bootstrap
bootstrap: ## Installs Go packages and sets defaults
	@echo "Installing Go packages using 'go install'"
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/cosmtrek/air@latest
	go install go101.org/golds@latest
	go install golang.org/x/lint/golint@latest
	go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	go install github.com/jstemmer/go-junit-report@latest
	go install github.com/t-yuki/gocover-cobertura@latest
	go install github.com/swaggo/swag/cmd/swag@latest


.PHONY: bootstrap-files
bootstrap-files: ## Copies all of the files
	@echo "Copying files..."
	cp config.example.yml config.yml
	cp sqreen.yaml.example sqreen.yaml

.PHONY: dev
dev: ## Runs the Gin HTTP server
	air

.PHONY: build
build: ## Runs the Golang build for production, including the Sqreen intrumentation TODO: Improve setup
	./bin/build

.PHONY: swagger
swagger: ## Generates Swagger documentation based on HTTP Handler comments
	./bin/genswagger

.PHONY: sql
sql: ## Checks SQL for syntax and type errors, then generates Go code for ./SQL/ files
	@sqlc compile
	@sqlc generate
	@sqlc compile

.PHONY: test
test: ## Runs the unit tests
	go test ./...

.PHONY: fmt
fmt: ## Code style formatting and linting check for the project
	goimports -w -local git.bytecode.nl/bytecode/genesis/server .
	go fmt ./...
	go vet ./...
	golint ./...
	@echo "Checking linting:"
	golangci-lint run --sort-results

.PHONY: docs
docs: ## Runs documentation website
	golds .

.PHONY: ci
ci: ## Runs the CI tasks
	@echo "Checking if golangci-lint is installed"
	@golangci-lint --version
	@echo "Running linting and static code analysis"
	golangci-lint run --sort-results --timeout 1h
	go vet ./...
	golint ./...
	goimports -e -d -local git.bytecode.nl/bytecode/genesis/server .
	@echo "Running tests with coverage"
	go test -coverprofile=coverage.txt -covermode count -v 2>&1 ./... | go-junit-report -set-exit-code > junit.xml
	gocover-cobertura < coverage.txt > cobertura.xml

.PHONY: migrateup
migrateup: ## Runs the database migrations up
	go run . migrate -d up --verbose

.PHONY: migratedown
migratedown: ## Runs the database migrations down
	go run . migrate -d down --verbose

.PHONY: migration
migration: ## Asks for the new migration name and creates migration files
	./bin/create_migration

.PHONY: newdb
newdb: ## Migrates db down and up - WARNING: ALL DB DATA IS DELETED
	make migratedown
	make migrateup

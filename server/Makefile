help: # Shows all commands
	@echo 'All Makefile commands'
	@cat Makefile | grep -v '^\W' | grep -v '^\.' | grep -v -e '^$$'

.PHONY: bootstrap
bootstrap: # Installs Go packages and sets defaults
	@echo "Installing Go packages 'go install'"
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/cosmtrek/air@latest
	@go install go101.org/golds@latest
	@go install golang.org/x/lint/golint@latest
	@go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
	@echo "Packages installed. Copying files..."
#	cp .env.example .env
#	cp sqreen.yaml.example sqreen.yaml

.PHONY: dev
dev: # Runs the Gin HTTP server
	air

.PHONY: build
build: # Runs the Golang build for production, including the Sqreen intrumentation
	./bin/build

.PHONY: sql
sql: # Checks SQL for syntax and type errors, then generates Go code for ./SQL/ files
	@sqlc compile
	@sqlc generate
	@sqlc compile

.PHONY: test
test: # Runs the unit tests
	go test ./...

.PHONY: lint
lint: # Code style formatting and linting check for the project
	goimports -w .
	go fmt ./...
	go vet ./...
	golint ./...

.PHONY: docs
docs: # Runs documentation website
	golds .

.PHONY: ci
ci: # Runs the CI tasks
	@echo "Checking if golangci-lint is installed"
	@golangci-lint --version
	@echo "Running golangci-lint"
	@golangci-lint run
	@echo "Running go vet"
	@go vet ./...
	@echo "Running golint"
	@golint ./...
	@echo "Running goimports checks (displaying changes, does not fail CI)"
	@goimports -e -d .

.PHONY: migrateup
migrateup: # Runs the database migrations up
	go run . migrate -d up --verbose

.PHONY: migratedown
migratedown: # Runs the database migrations down
	go run . migrate -d down --verbose

.PHONY: migration
migration: # Asks for the new migration name and creates migration files
	./bin/create_migration

.PHONY: newdb
newdb: # Migrates db down and up - WARNING: ALL DB DATA IS DELETED
	make migratedown
	make migrateup

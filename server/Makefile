help: # Shows all commands
	@echo 'All Makefile commands'
	@cat Makefile | grep -v '^\W'

.PHONY: bootstrap
bootstrap: # Installs Go packages and sets defaults
	go get go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/tools/cmd/goimports
	cp .env.example .env
	cp sqreen.yaml.example sqreen.yaml

.PHONY: dev
dev: # Runs the Gin HTTP server
	air

.PHONY: build
build: # Runs the Golang build for production, including the Sqreen intrumentation
	./bin/build

.PHONY: test
test: # Runs the unit tests
	go test ./...

.PHONY: fmt
fmt: # Code style formatting for the project
	goimports -w .
	go fmt ./...

.PHONY: ci
ci: # Runs the CI tasks # TODO: Implement
	# Run tests
	# Run linting checks
	# Run goimports checks
	@echo "Not implemented"
	exit 1

.PHONY: migrateup
migrateup: # Runs the database migrations up
	go run . migrate -d up --verbose

.PHONY: migratedown
migratedown: # Runs the database migrations down
	go run . migrate -d down --verbose

.PHONY: newmigration
newmigration: # Asks for the new migration name and creates migration files
	./bin/create_migration

.PHONY: newdb
newdb: # Migrates db down and up - WARNING: ALL DB DATA IS DELETED
	make migratedown
	make migrateup

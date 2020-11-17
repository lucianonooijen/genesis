help: # Shows all commands
	@echo 'All Makefile commands'
	@cat Makefile | grep -v '^\W'

.PHONY: bootstrap
bootstrap: # Copies the defaults
	cp .env.example .env
	cp sqreen.yaml.example sqreen.yaml

.PHONY: bootstrap
dev: # Runs the Gin HTTP server
	air

.PHONY: bootstrap
build: # Runs the Golang build for production, including the Sqreen intrumentation
	./bin/build

.PHONY: bootstrap
test: # Runs the unit tests
	go test ./...

.PHONY: bootstrap
fmt: # Code style formatting for the project
	go fmt ./...

.PHONY: bootstrap
ci: # Runs the CI tasks # TODO: Implement
	# Run tests
	# Run linting checks
	# Run goimports checks
	@echo "Not implemented"
	exit 1

.PHONY: bootstrap
migrateup: # Runs the database migrations up
	go run . migrate -d up --verbose

.PHONY: bootstrap
migratedown: # Runs the database migrations down
	go run . migrate -d down --verbose

.PHONY: bootstrap
newmigration: # Asks for the new migration name and creates migration files
	./bin/create_migration

.PHONY: bootstrap
newdb: # Migrates db down and up - WARNING: ALL DB DATA IS DELETED
	make migratedown
	make migrateup

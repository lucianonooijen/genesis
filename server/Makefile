help: # Shows all commands
	@echo 'All Makefile commands'
	@cat Makefile | grep -v '^\W'

.PHONY: bootstrap
bootstrap: # Installs Go packages and sets defaults
	@echo "Installing Go packages with GO111MODULE=off..."
	@GO111MODULE=off go get -u golang.org/x/tools/cmd/goimports
	@GO111MODULE=off go get -u github.com/cosmtrek/air
	@GO111MODULE=off go get -u go101.org/golds
	@GO111MODULE=off go get -u golang.org/x/lint/golint
	@echo "Packages installed. Copying files..."
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

.PHONY: lint
lint: # Code style formatting and linting check for the project
	goimports -w .
	go fmt ./...
	golint ./...

.PHONY: docs
docs: # Runs documentation website # TODO: Improve
	golds .

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

# Genesis backend server

## Requirements

* Go 1.15 or higher
* [Air](https://github.com/cosmtrek/air), [Goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) and [Golds](https://github.com/go101/golds) in your $PATH (`go get`-ted by `make bootstrap`)
* PostgreSQL

## Installation

* Run `make bootstrap` and set the correct details in `.env`
* Copy the `.env.example` file to `.env` and set the correct credentials
* Run `make migrateup` to run all database migrations

## Development and commands

To run the application, run `make dev`.

For all commands, run `make` and all commands will show in your console.

## Application architecture and import boundaries

TODO

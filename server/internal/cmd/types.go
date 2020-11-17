package cmd

import "git.bytecode.nl/bytecode/genesis/internal/infrastructure/migrator"

// ApplicationClosures is a collection of closures that will be called based on the CLI options
type ApplicationClosures struct {
	Migrate     func(migrator.Direction) error `validate:"required"`
	StartServer func() error                   `validate:"required"`
}

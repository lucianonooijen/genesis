package interactors

import (
	"database/sql"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/data/mailer"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/passhash"
)

// Services contains all the shared services in the application.
type Services struct {
	// Config
	Config config.Config `validate:"required"`

	// Infrastructure
	BaseLogger *zap.Logger   `validate:"required"`
	JWT        jwt.Util      `validate:"required"`
	PassHash   passhash.Util `validate:"required"`
	DBConn     *sql.DB       `validate:"required"`

	// Data
	Mailer   mailer.Mailer     `validate:"required"`
	Database *database.Queries `validate:"required"`
}

package interactors

import (
	"database/sql"

	"go.uber.org/zap"

	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/mailer"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/pushnotifications"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/config"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/passhash"
)

// Services contains all the shared services in the application.
type Services struct {
	// Config
	Config config.Config `validate:"required"`

	// Infrastructure
	BaseLogger *zap.Logger    `validate:"required"`
	JWT        *jwt.Util      `validate:"required"`
	PassHash   *passhash.Util `validate:"required"`
	DBConn     *sql.DB        `validate:"required"`

	// Data
	Mailer            *mailer.Mailer                 `validate:"required"`
	Database          *database.Queries              `validate:"required"`
	PushNotifications *pushnotifications.PushService `validate:"required"`
}

package cmd

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/TheZeroSlave/zapsentry"
	"github.com/getsentry/sentry-go"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zapio"

	"git.bytecode.nl/bytecode/genesis/server/internal/constants"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/mailer"
	"git.bytecode.nl/bytecode/genesis/server/internal/data/pushnotifications"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/config"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/logger"
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/passhash"
	"git.bytecode.nl/bytecode/genesis/server/internal/interactors"
)

// nolint:wsl,funlen // with cuddles is better readable here, also allow long function length because it's the init func
func loadServices() *interactors.Services {
	var (
		services interactors.Services
		err      error
	)

	// Config
	// TODO: Change config to YAML
	services.Config, err = config.LoadConfig()
	panicOnErr(err)

	// Logger
	logBase, err := logger.New(services.Config.IsDevMode)
	panicOnErr(err)
	logInstance := logBase.Sugar()
	services.BaseLogger = logBase
	logMain := logInstance.Named("main_init")
	logMain.Info("Hello world. Config loaded, logger configured.")

	// Sentry setup
	logMain.Debug("setting up Sentry")
	err = sentry.Init(sentry.ClientOptions{
		Dsn:         services.Config.SentryDSN,
		Debug:       services.Config.SentryEnvironment == "development",
		Release:     constants.APIVersion,
		Environment: services.Config.SentryEnvironment,
		DebugWriter: &zapio.Writer{
			Log:   logBase.Named("sentry"),
			Level: zapcore.DebugLevel,
		},
	})
	panicOnErr(err)

	// Connect logger and Sentry
	logMain.Debug("connecting Sentry to BaseLogger")
	zapSentryConfig := zapsentry.Configuration{
		Level:             zapcore.ErrorLevel, // when to send message to sentry
		EnableBreadcrumbs: true,               // enable sending breadcrumbs to Sentry
		DisableStacktrace: false,              // we want stacktrace
		BreadcrumbLevel:   zapcore.InfoLevel,  // at what level should we sent breadcrumbs to sentry
	}
	core, err := zapsentry.NewCore(zapSentryConfig, zapsentry.NewSentryClientFromDSN(services.Config.SentryDSN))
	panicOnErr(err)
	services.BaseLogger = services.BaseLogger.With(zapsentry.NewScope())
	services.BaseLogger = zapsentry.AttachCoreToLogger(core, services.BaseLogger)

	// DBConn connection
	logMain.Debug("opening database connection")
	services.DBConn, err = sql.Open("postgres", services.Config.DatabaseConnectionString())
	panicOnErr(err)
	err = services.DBConn.Ping()
	if err != nil {
		logMain.Error("error pinging database", err)
		panic(err)
	}

	// JWT
	// TODO: make a refresh system and new login on expired token in app, so the JWTs can have a validity of month max for security reasons
	logMain.Debug("Building JWT instance for users")
	century := time.Hour * 24 * 365 * 100                                                      // nolint:gomnd // 24*365*100 speaks for itself
	services.JWT, err = jwt.New(services.Config.JWTSecret, constants.JwtSubjectUsers, century) // make JWTs valid for one year
	panicOnErr(err)

	// Password hasher
	logMain.Debug("Building password hasher instance")
	services.PassHash = passhash.New()

	// Static file URL
	staticFileURLBase := fmt.Sprintf("%s%s%s", services.Config.ServerHostname, constants.BasePathAPI, constants.APIStaticPath)
	logMain.Info(fmt.Sprintf("Serving static files from %s", staticFileURLBase))

	// Mailer
	logMain.Debug("Building mailer instance")
	services.Mailer, err = mailer.New(logBase, services.Config.EmailSenderEmail, services.Config.EmailSenderName, services.Config.SendinblueAPIKey, staticFileURLBase)
	panicOnErr(err)

	// Push notification instance
	logMain.Debug("Building push notification service instance")
	services.PushNotifications, err = pushnotifications.New(&services.Config, logBase)
	panicOnErr(err)

	// Database instance
	logMain.Debug("Building database queryer instance")
	services.Database = database.New(services.DBConn)

	// Validating the Applications interactor instance
	validate := validator.New()
	err = validate.Struct(services)
	panicOnErr(err)

	return &services
}

// panicOnErr is a helper in the init phase of the project to avoid error checks clutter.
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

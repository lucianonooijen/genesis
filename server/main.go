/*
Copyright © 2020 Bytecode Digital Agency B.V.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"

	"git.bytecode.nl/bytecode/genesis/internal/cmd"
	"git.bytecode.nl/bytecode/genesis/internal/constants"
	"git.bytecode.nl/bytecode/genesis/internal/data/database"
	"git.bytecode.nl/bytecode/genesis/internal/data/mailer"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/jwt"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/logger"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/migrator"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/passhash"
	"git.bytecode.nl/bytecode/genesis/internal/interactors"
	"git.bytecode.nl/bytecode/genesis/internal/server"
)

// TODO: Rework this and use this in cmd/ that returns the interactor,
// only load CLI thingy here, nothing else

func main() {
	var services interactors.Services

	// Config
	// TODO: Change config to YAML
	c, err := config.LoadConfig()
	panicOnErr(err)
	services.Config = c

	// Logger
	logBase, err := logger.New(c.IsDevMode)
	panicOnErr(err)
	services.BaseLogger = logBase
	logInstance := logBase.Sugar()
	logMain := logInstance.Named("main_init")
	logMain.Info("Hello world. Config loaded, logger configured.")

	// DbConn connection
	logMain.Debug("opening database connection")
	services.DbConn, err = sql.Open("postgres", c.DatabaseConnectionString())
	panicOnErr(err)
	err = services.DbConn.Ping()
	if err != nil {
		logMain.Error("error pinging database", err)
		panic(err)
	}

	// JWT
	logMain.Debug("Building JWT instance for users")
	services.JWT, err = jwt.New(c.JWTSecret, "users", time.Hour*24*365) // Valid for 1 year
	panicOnErr(err)

	// Password hasher
	logMain.Debug("Building password hasher instance")
	services.PassHash = passhash.New()

	// Static file URL
	staticFileURLBase := fmt.Sprintf("%s%s%s", c.ServerHostname, constants.BasePathAPI, constants.APIStaticPath)
	logMain.Info(fmt.Sprintf("Serving static files from %s", staticFileURLBase))

	// Mailer
	logMain.Debug("Building mailer instance")
	services.Mailer, err = mailer.New(logBase, c.EmailSenderEmail, c.EmailSenderName, c.SendinblueAPIKey, staticFileURLBase)
	panicOnErr(err)

	// Database instance
	logMain.Debug("Building database queryer instance")
	services.Database = database.New(services.DbConn)

	// Validating the Applications interactor instance
	validate := validator.New()
	err = validate.Struct(services)
	panicOnErr(err)

	/**
	 * Below are the instances passed to the CLI
	 */

	// Build server instance
	logMain.Info("Building server instance")
	s, err := server.New(&services)
	panicOnErr(err)

	// Migrator
	logMain.Debug("Building migrator")
	migrate := migrator.New(services.DbConn, c.DatabaseName, getMigrationSourceURL())

	// Build application closures
	logMain.Debug("Assembling ApplicationClosures")
	apps := cmd.ApplicationClosures{
		Migrate:     migrate,
		StartServer: s.Start,
	}

	// Pass to cmd package
	logMain.Info("Passing everything into the CLI handler and starting application")
	cmd.Execute(c, apps)
}

// panicOnErr is a helper in the init phase of the project to avoid error checks clutter
func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

// TODO: Embed migrations into the binary
func getMigrationSourceURL() string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat("./migrations"); os.IsNotExist(err) {
		log.Fatal(errors.New("migrations directory not present"))
	}
	return "file://" + workingDir + "/migrations"
}

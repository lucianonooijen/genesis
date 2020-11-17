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
	"errors"
	"os"
	"time"

	"git.bytecode.nl/bytecode/genesis/internal/utils/logger"

	"git.bytecode.nl/bytecode/genesis/internal/cmd"
	"git.bytecode.nl/bytecode/genesis/internal/data/jwt"
	"git.bytecode.nl/bytecode/genesis/internal/data/passhash"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/migrator"
	"git.bytecode.nl/bytecode/genesis/internal/server"
)

var log = logger.New("main")

// TODO: Create cleaner startup procedure, but only when really necessary

func main() {
	/**
	 * Step 1: load config and configure the logger
	 */
	c, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	if err = logger.Configure(c.IsDevMode, c.SentryDSN, c.SentryEnvironment); err != nil {
		panic(err)
	}
	log.Info("Config loaded, logger configured. Hi there!")

	/**
	 * Step 2: build infrastructure instances
	 */
	log.Info("Building infrastructure instances")
	log.Trace("Building migrator")
	migrate := migrator.New(c.DatabaseConnectionString(), "postgres", c.DatabaseName, getMigrationSourceURL())

	/**
	 * Step 3: build the data instances
	 */
	log.Info("Building data instances")
	log.Trace("Building JWT instance for users")
	_, err = jwt.New(c.JWTSecret, "users", time.Hour*24*365) // Valid for 1 year
	exitOnErr(err)
	log.Trace("Building password hasher instance")
	_ = passhash.New()

	/**
	 * Step 4: build the domain instances
	 */
	log.Info("Building domain instances")
	// TODO: Implement

	/**
	 * Step 5: build the application closures
	 */
	log.Info("Building application closures")
	log.Trace("Building server instance")
	s, err := server.New(server.Requirements{
		Debug: c.IsDevMode,
		Port:  c.ServerPort,
	})
	log.Trace("Assembling ApplicationClosures")
	apps := cmd.ApplicationClosures{
		Migrate:     migrate,
		StartServer: s.Start,
	}

	/**
	 * Step 6: run cmd.Execute to run the correct part of the application
	 */
	log.Info("Passing everything into the CLI handler and starting application")
	cmd.Execute(c, apps)
}

// exitOnErr is a helper in the init phase of the project to avoid error checks clutter
func exitOnErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

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

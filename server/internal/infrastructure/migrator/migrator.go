package migrator

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"

	//"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	// Needs side effects from golang-migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"
	// Needs side effects from lib/pg
	_ "github.com/lib/pq"
)

// Direction indicates the migration direction
type Direction int

// UpAll and DownAll indicate the migration direction
const (
	UpAll Direction = iota
	DownAll
)

// New returns a migrator closure that will accept UpAll or DownAll to up or down migrate the database
func New(dbString string, dbType string, dbName string, migrationScriptsPath string) func(Direction) error {
	return func(direction Direction) error {
		// TODO: Add support for migrating +1/-1 or to a specific migration
		// Direction check
		if direction != UpAll && direction != DownAll {
			return fmt.Errorf("migration direction should be 'UpAll' or 'DownAll'")
		}

		// Build database instance and migrator
		db, err := sql.Open(dbType, dbString)
		if err != nil {
			return err
		}
		defer db.Close()

		// Generating migration instance
		driver, err := postgres.WithInstance(db, &postgres.Config{})
		if err != nil {
			return err
		}
		migrator, err := migrate.NewWithDatabaseInstance(migrationScriptsPath, dbName, driver)
		if err != nil {
			return err
		}

		// Do the actual migration
		if direction == UpAll {
			return handleMigratorErrors(migrator.Up())
		}
		if direction == DownAll {
			return handleMigratorErrors(migrator.Down())
		}
		return errors.New("you should not see this")
	}
}

func handleMigratorErrors(err error) error {
	if err == migrate.ErrNoChange { // Do not report error when no database change
		return nil
	}
	return err
}

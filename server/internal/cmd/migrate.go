package cmd

import (
	"git.bytecode.nl/bytecode/genesis/server/internal/infrastructure/migrator"

	"github.com/spf13/cobra"
)

var (
	dbMigrateDirection string
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Runs database migrations",
	Long:  `Runs the PostgreSQL database migration scripts`,
	Run: func(cmd *cobra.Command, args []string) {
		s := loadServices()
		l := s.BaseLogger.Named("migrate_cmd")

		migrate := migrator.New(s.DBConn, s.Config.DatabaseName)

		if dbMigrateDirection == "up" {
			l.Info("Migrating to latest...")
			if err := migrate(migrator.UpAll); err != nil {
				panic(err)
			}
			l.Info("Migrations successful")
			return
		}

		if dbMigrateDirection == "down" {
			l.Info("Migrating all the way down...")
			if err := migrate(migrator.DownAll); err != nil {
				panic(err)
			}
			l.Info("Migrations successful")
			return
		}

		panic("invalid migration direction")
	},
}

func init() { // nolint:gochecknoinits // needed for sane Cobra use
	migrateCmd.Flags().StringVarP(&dbMigrateDirection, "direction", "d", "up|down", "The direction for the database migrations")

	// Required flags
	if err := migrateCmd.MarkFlagRequired("direction"); err != nil {
		panic(err)
	}
}

package cmd

import (
	"fmt"

	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/migrator"

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
		performDataChecks()
		if dbMigrateDirection == "up" {
			fmt.Println("Migrating to latest...")
			if err := apps.Migrate(migrator.UpAll); err != nil {
				panic(err)
			}
			fmt.Println("Migrations successful")
			return
		}
		if dbMigrateDirection == "down" {
			fmt.Println("Migrating all the way down...")
			if err := apps.Migrate(migrator.DownAll); err != nil {
				panic(err)
			}
			fmt.Println("Migrations successful")
			return
		}
		panic("invalid migration direction")
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&dbMigrateDirection, "direction", "d", "up|down", "The direction for the database migrations")

	// Required flags
	if err := migrateCmd.MarkFlagRequired("direction"); err != nil {
		panic(err)
	}

}

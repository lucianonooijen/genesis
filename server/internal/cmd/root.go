package cmd

import (
	"git.bytecode.nl/bytecode/genesis/internal/infrastructure/config"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cobra"
)

var (
	verbose = false
	apps    = ApplicationClosures{}
	c       = config.Config{}
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genesisbackend",
	Short: "The Genesis platform backend application",
	Long: `The Genesis platform backend application including the REST API server and migrator.
Application must be run with a 'migrations/' directory and '.env'-file in the same folder`,
}

// Execute runs Cobra and starts/runs the correct application/jobs
func Execute(cfg config.Config, applications ApplicationClosures) {
	apps = applications
	c = cfg
	performDataChecks()
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	// Root flags
	// TODO: implement verbose output
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output (not implemented yet)")

	// Add commands
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrateCmd)
}

func performDataChecks() {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		panic(err)
	}
	err = validate.Struct(apps)
	if err != nil {
		panic(err)
	}
}

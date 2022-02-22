package cmd

import (
	"github.com/spf13/cobra"
)

var (
	verbose = false
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "genesisbackend",
	Short: "The Genesis platform backend application",
	Long: `The Genesis platform backend application including the REST API server and migrator.
Application must be run with a 'migrations/' directory and 'config.yml'-file in the same folder (or environment variables set)`,
}

// Execute runs Cobra and starts/runs the correct application/jobs.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() { // nolint:gochecknoinits // needed for sane Cobra use
	// Root flags
	// TODO: implement verbose output
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output (not implemented yet)")

	// Add commands
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(migrateCmd)
}

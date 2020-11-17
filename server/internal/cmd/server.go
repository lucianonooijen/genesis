package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Running the backend server application",
	Long:  `The application running the REST API server`,
	Run: func(cmd *cobra.Command, args []string) {
		performDataChecks()
		err := apps.StartServer()
		if err != nil {
			panic(err)
		}
	},
}

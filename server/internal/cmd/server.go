package cmd

import (
	"github.com/spf13/cobra"

	"git.bytecode.nl/bytecode/genesis/internal/server"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Running the backend server application",
	Long:  `The application running the REST API server`,
	Run: func(cmd *cobra.Command, args []string) {
		services := loadServices()

		s, err := server.New(services)
		if err != nil {
			panic(err)
		}

		err = s.Start()
		if err != nil {
			panic(err)
		}
	},
}

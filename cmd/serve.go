package cmd

import (
	"github.com/spf13/cobra"
	"github.com/arisro/talos/cmd/server"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a HTTP service.",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
}
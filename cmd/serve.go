package cmd

import (
	"github.com/arisro/talos/cmd/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts a HTTP service.",
	Run:   server.StartServer(),
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

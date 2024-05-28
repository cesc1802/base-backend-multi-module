package server

import (
	"log"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starting HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("start http server")
	},
}

func RegisterServer(root *cobra.Command) {
	root.AddCommand(serverCmd)
}

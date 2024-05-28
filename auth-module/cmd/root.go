package cmd

import (
	"log"

	"auth-module/cmd/migrate"
	"auth-module/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "auth-module",
	Short:            "This is authentication module",
	TraverseChildren: true,
}

func init() {
	server.RegisterServer(rootCmd)
	migrate.RegisterMigrate(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("application cannot start %v", err)
	}
}

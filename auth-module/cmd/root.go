package cmd

import (
	"log"

	"github.com/cesc1802/auth-module/cmd/migrate"
	"github.com/cesc1802/auth-module/cmd/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:              "github.com/cesc1802/auth-module",
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

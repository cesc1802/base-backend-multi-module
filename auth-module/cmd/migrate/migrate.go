package migrate

import (
	"log"

	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "",
}

var up = &cobra.Command{
	Use:   "up",
	Short: "This will be use to migrate new change",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("run migrate up")
	},
}

var down = &cobra.Command{
	Use:   "down",
	Short: "This will be use to rollback new change",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("run rollback")
	},
}

func RegisterMigrate(root *cobra.Command) {
	migrate.AddCommand(up, down)
	root.AddCommand(migrate)
}

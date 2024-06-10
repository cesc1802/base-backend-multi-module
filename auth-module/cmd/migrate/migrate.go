package migrate

import (
	"log"

	"github.com/cesc1802/auth-module/migration"
	"github.com/cesc1802/share-module/config"
	"github.com/cesc1802/share-module/system"
	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "",
}

var up = &cobra.Command{
	Use:   "up",
	Short: "This will be use to migrate new change",
	RunE: func(cmd *cobra.Command, args []string) error {

		cfg, err := config.LoadAppConfig(".")
		if err != nil {
			log.Fatalln(err)
			return err
		}
		sys := system.New(cfg, cmd.Parent().Name())

		if err := sys.MigrateDB(migration.FS); err != nil {
			log.Fatalln(err)
			return err
		}
		return nil
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

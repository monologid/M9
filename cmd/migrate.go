package cmd

import (
	"strings"

	"github.com/monologid/m9/config"
	"github.com/monologid/m9/db"
	"github.com/monologid/m9/login"
	"github.com/spf13/cobra"
)

var migrate = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate schemas into database.",
	Run: func(cmd *cobra.Command, args []string) {
		config.New(configPath)

		db.New(config.C.Database.Engine, config.C.Database.URL)
		db.SetVerbose(verbose)

		switch strings.ToLower(module) {
		case "login":
			login.Migrate()
		default:
			login.Migrate()
		}

		log("Database schemas have been migrated.")
	},
}

func init() {
	migrate.PersistentFlags().StringVarP(&module, "module", "m", "all", "Module name (e.g. login)")
	migrate.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "Verbose all database query")
}

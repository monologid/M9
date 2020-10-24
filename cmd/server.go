package cmd

import (
	"github.com/monologid/m9/config"
	"github.com/monologid/m9/db"
	"github.com/monologid/m9/httpsvr"
	"github.com/monologid/m9/login"
	"github.com/spf13/cobra"
)

var server = &cobra.Command{
	Use:   "server",
	Short: "HTTP server",
}

var start = &cobra.Command{
	Use:   "start",
	Short: "Start for starting HTTP server.",
	Run: func(cmd *cobra.Command, args []string) {
		config.New(configPath)

		db.New(config.C.Database.Engine, config.C.Database.URL)
		db.SetVerbose(verbose)

		svr := httpsvr.New().Initialize()

		if len(args) == 0 {
			login.New(svr.Server())

			svr.Start(verbose)
			return
		}

		moduleName := args[0]

		switch moduleName {
		case "login":
			login.New(svr.Server())
		}

		svr.Start(verbose)
	},
}

func init() {
	start.PersistentFlags().StringVarP(&configPath, "config-path", "c", "/etc/m9", "Set config path")
	start.PersistentFlags().BoolVarP(&verbose, "verbose", "", false, "Verbose all database query")
}

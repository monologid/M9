package cmd

import (
	"github.com/monologid/m9/config"
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
		svr := httpsvr.New().Initialize()

		if len(args) == 0 {
			login.New(svr.Server())

			svr.Start()
			return
		}

		moduleName := args[0]

		switch moduleName {
		case "login":
			login.New(svr.Server())
		}

		svr.Start()
	},
}

func init() {
	start.PersistentFlags().StringVarP(&configPath, "config-path", "c", "/etc/m9", "Set config path")
}

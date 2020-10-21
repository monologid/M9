package cmd

import (
	"github.com/spf13/cobra"
)

var (
	module         string
	configPath     string
	databaseEngine string
	databaseURL    string
	debug          bool
)

// Initialize executes the cmd package
func Initialize() {
	commands := &cobra.Command{
		Use:   "m9",
		Short: "M9 is an open-source software project by Monolog written in Go.",
	}

	server.AddCommand(start)

	commands.AddCommand(server)
	commands.Execute()
}

// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// pluginCmd represents the plugin command
var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Plugin management for the Project Flogo CLI",
	Run:   runPlugin,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(pluginCmd)
}

// runPlugin is the actual execution of the command
func runPlugin(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe Plugin command supports the plugin management capabilities of the Project Flogo CLI.\nThe commands available are:\n\n")

	// Print all subcommands
	for _, command := range cmd.Commands() {
		if command.Use != "help [command]" {
			fmt.Printf("%s %s\n", fgutil.RightPadToLen(command.Use, ".", 25), command.Short)
		}
	}

	fmt.Printf("\nRun 'flogo help plugin [command]' for more details\n\n")
}

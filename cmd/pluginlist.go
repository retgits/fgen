// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// pluginListCmd represents the plugin command
var pluginListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all installed plugins",
	Run:   runPluginList,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	pluginCmd.AddCommand(pluginListCmd)
}

// runPluginList is the actual execution of the command
func runPluginList(cmd *cobra.Command, args []string) {
	commandList := rootCmd.Commands()
	fmt.Printf("Your Project Flogo CLI has the following plugins installed:\n\n")
	for _, command := range commandList {
		if command.Use != "help [command]" {
			getCommands(command, 0)
		}
	}
	fmt.Printf("\n")
}

func getCommands(cmd *cobra.Command, depth int) {
	// Print the current command
	fmt.Printf("%s %s\n", fgutil.RightPadToLen(fgutil.LeftPad(cmd.Use, "  ", depth), ".", 25), cmd.Short)

	// Print all subcommands
	subCommands := cmd.Commands()
	if len(subCommands) > 0 {
		for _, command := range subCommands {
			if command.Use != "help [command]" {
				getCommands(command, depth+1)
			}
		}
	}
}

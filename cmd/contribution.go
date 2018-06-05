// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"

	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// contributionCmd represents the xxx command
var contribCmd = &cobra.Command{
	Use:   "contrib",
	Short: "Lists all contrib actions available to Project Flogo",
	Run:   runContrib,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	rootCmd.AddCommand(contribCmd)
}

// runContrib is the actual execution of the command
func runContrib(cmd *cobra.Command, args []string) {
	fmt.Printf("\nThe contrib command supports the contribution capabilities of the Project Flogo CLI.\nThe commands available are:\n\n")

	// Print all subcommands
	for _, command := range cmd.Commands() {
		if command.Use != "help [command]" {
			fmt.Printf("%s %s\n", fgutil.RightPadToLen(command.Use, ".", 25), command.Short)
		}
	}

	fmt.Printf("\nRun 'flogo help plugin [command]' for more details\n\n")
}

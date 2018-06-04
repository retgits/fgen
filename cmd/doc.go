package cmd

/*
Each command that is implemented has a default structure so that the layout is the same:

// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/retgits/fgen/gen"
	"github.com/spf13/cobra"
)

// xxxCmd represents the xxx command
var xxxCmd = &cobra.Command{
	Use:   "xxx",
	Short: "XXX does...",
	Run:   runXXX,
}

// Flags
var (

)

// Variables
var (

)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(xxxCmd)
}

// runXXX is the actual execution of the command
func runXXX(cmd *cobra.Command, args []string) {

}


*/

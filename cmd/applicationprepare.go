// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/spf13/cobra"
)

// appPrepareCmd represents the xxx command
var appPrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare the flogo application",
	Run:   runPrepareApp,
}

// Flags
var (
	appPrepareOptimize bool
	appPrepareEmbed    bool
)

// Variables
var ()

// init registers the command and flags
func init() {
	appCmd.AddCommand(appPrepareCmd)
	appPrepareCmd.Flags().BoolVarP(&appPrepareOptimize, "optimize", "o", false, "Optimize the preparation")
	appPrepareCmd.Flags().BoolVarP(&appPrepareEmbed, "embed", "e", false, "Embed configuration into the application")
}

// runPrepareApp is the actual execution of the command
func runPrepareApp(cmd *cobra.Command, args []string) {
	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	options := &app.PrepareOptions{OptimizeImports: appPrepareOptimize, EmbedConfig: appPrepareEmbed}
	err = app.PrepareApp(app.SetupExistingProjectEnv(appDir), options)
	if err != nil {
		fmt.Printf("Error while preparing app: %s\n\n", err.Error())
	}
}

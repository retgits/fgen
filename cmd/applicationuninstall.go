// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/spf13/cobra"
)

// appUninstallCmd represents the xxx command
var appUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall an app contribution",
	Run:   runUninstallApp,
}

// Flags
var (
	appUninstallName string
)

// Variables
var ()

// init registers the command and flags
func init() {
	appCmd.AddCommand(appUninstallCmd)
	appUninstallCmd.Flags().StringVarP(&appUninstallName, "name", "n", "", "The name of the contribution (required)")
	appUninstallCmd.MarkFlagRequired("name")
}

// runUninstallApp is the actual execution of the command
func runUninstallApp(cmd *cobra.Command, args []string) {
	contribPath := appUninstallName

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	err = app.UninstallDependency(app.SetupExistingProjectEnv(appDir), contribPath)
	if err != nil {
		fmt.Printf("Error while uninstalling contribution: %s\n\n", err.Error())
	}
}

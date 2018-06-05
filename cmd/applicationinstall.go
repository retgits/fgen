// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/spf13/cobra"
)

// appInstallCmd represents the xxx command
var appInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install an app contribution",
	Run:   runInstallApp,
}

// Flags
var (
	appInstallName    string
	appInstallVersion string
	appInstallPalette bool
)

// Variables
var ()

// init registers the command and flags
func init() {
	appCmd.AddCommand(appInstallCmd)
	appInstallCmd.Flags().StringVarP(&appInstallVersion, "version", "v", "", "Specify the version of the contribution (optional)")
	appInstallCmd.Flags().StringVarP(&appInstallName, "name", "n", "", "The name of the contribution (required)")
	appInstallCmd.Flags().BoolVarP(&appInstallPalette, "palette", "p", false, "Install palette file")
	appInstallCmd.MarkFlagRequired("name")
}

// runInstallApp is the actual execution of the command
func runInstallApp(cmd *cobra.Command, args []string) {
	contribPath, version := splitVersion(appInstallName)

	if len(appInstallVersion) != 0 {
		version = appInstallVersion
	}

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	if appInstallPalette {
		err = app.InstallPalette(app.SetupExistingProjectEnv(appDir), contribPath)
		if err != nil {
			fmt.Printf("Error while installing contribution: %s\n\n", err.Error())
		}
	}

	err = app.InstallDependency(app.SetupExistingProjectEnv(appDir), contribPath, version)
	if err != nil {
		fmt.Printf("Error while installing contribution: %s\n\n", err.Error())
	}
}

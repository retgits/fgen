// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/device"

	"github.com/spf13/cobra"
)

// deviceInstallCmd represents the xxx command
var deviceInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install a device contribution",
	Run:   runDeviceInstall,
}

// Flags
var (
	deviceInstallVersion string
	deviceInstallName    string
)

// Variables
var ()

// init registers the command and flags
func init() {
	deviceCmd.AddCommand(deviceInstallCmd)
	deviceInstallCmd.Flags().StringVarP(&deviceInstallVersion, "version", "v", "", "Specify the version of the contribution (optional)")
	deviceInstallCmd.Flags().StringVarP(&deviceInstallName, "name", "n", "", "The name of the contribution (required)")
	deviceInstallCmd.MarkFlagRequired("name")
}

// runDeviceInstall is the actual execution of the command
func runDeviceInstall(cmd *cobra.Command, args []string) {
	contribPath, version := splitVersion(deviceInstallName)

	if len(deviceInstallVersion) != 0 {
		version = deviceInstallVersion
	}

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	err = device.InstallContribution(device.SetupExistingDeviceProjectEnv(appDir), contribPath, version)
	if err != nil {
		fmt.Printf("Error while installing contribution: %s\n\n", err.Error())
	}
}

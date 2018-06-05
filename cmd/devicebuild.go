// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/device"

	"github.com/spf13/cobra"
)

// deviceBuildCmd represents the xxx command
var deviceBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a device application",
	Run:   runDeviceBuild,
}

// Flags
var (
	skipPrepare bool
)

// Variables
var ()

// init registers the command and flags
func init() {
	deviceCmd.AddCommand(deviceBuildCmd)
	deviceBuildCmd.Flags().BoolVarP(&skipPrepare, "no-gen", "g", false, "Only perform the build, without performing the generation of metadata")
}

// runDeviceBuild is the actual execution of the command
func runDeviceBuild(cmd *cobra.Command, args []string) {
	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	options := &device.BuildOptions{SkipPrepare: skipPrepare, PrepareOptions: &device.PrepareOptions{}}
	err = device.BuildDevice(device.SetupExistingDeviceProjectEnv(appDir), options)
	if err != nil {
		fmt.Printf("Error while building device project: %s\n\n", err.Error())
	}
}

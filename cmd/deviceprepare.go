// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/device"

	"github.com/spf13/cobra"
)

// devicePrepareCmd represents the xxx command
var devicePrepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Prepare the device application",
	Run:   runDevicePrepare,
}

// Flags
var (
	devicePrepareOptimize bool
	devicePrepareEmbed    bool
)

// Variables
var ()

// init registers the command and flags
func init() {
	deviceCmd.AddCommand(devicePrepareCmd)
	devicePrepareCmd.Flags().BoolVarP(&devicePrepareOptimize, "optimize", "o", false, "Optimize the preparation")
	devicePrepareCmd.Flags().BoolVarP(&devicePrepareEmbed, "embed", "e", false, "Embed configuration into the application")
}

// runDevicePrepare is the actual execution of the command
func runDevicePrepare(cmd *cobra.Command, args []string) {
	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	options := &device.PrepareOptions{}
	err = device.PrepareDevice(device.SetupExistingDeviceProjectEnv(appDir), options)
	if err != nil {
		fmt.Printf("Error while preparing device application: %s\n\n", err.Error())
	}
}

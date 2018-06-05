// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/device"

	"github.com/spf13/cobra"
)

// deviceUploadCmd represents the xxx command
var deviceUploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload the device application",
	Run:   runDeviceUpload,
}

// Flags
var ()

// Variables
var ()

// init registers the command and flags
func init() {
	deviceCmd.AddCommand(deviceUploadCmd)
}

// runDeviceUpload is the actual execution of the command
func runDeviceUpload(cmd *cobra.Command, args []string) {
	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	err = device.UploadDevice(device.SetupExistingDeviceProjectEnv(appDir))
	if err != nil {
		fmt.Printf("Error while uploading device application: %s\n\n", err.Error())
	}
}

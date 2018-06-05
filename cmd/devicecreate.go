// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/retgits/fgen/device"

	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// deviceCreateCmd represents the xxx command
var deviceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a device project",
	Run:   runDeviceCreate,
}

// Flags
var (
	deviceJSONFile   string
	deviceCreateName string
)

// Variables
var (
	tplSimpleDevice = `{
		"name": "mydevice",
		"type": "flogo:device",
		"version": "0.0.1",
		"description": "My flogo device application description",
		"device_profile": "github.com/TIBCOSoftware/flogo-contrib/device/profile/feather_m0_wifi",
		"mqtt_enabled":true,
		"settings": {
		  "mqtt:server":"192.168.1.50",
		  "mqtt:port":"1883",
		  "mqtt:user":"",
		  "mqtt:pass":"",
		  "wifi:ssid":"mynetwork",
		  "wifi:password": "mypass"
		},
		"triggers": [
		  {
			"id": "mqtt_trigger",
			"ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/device-mqtt",
			"actionId": "pin_on",
			"settings": {
			  "topic": "recievetopic"
			}
		  }
		],
		"actions": [
		  {
			"id": "pin_on",
			"ref": "github.com/TIBCOSoftware/flogo-contrib/action/device-activity",
			"data": {
			  "activity": {
				"ref": "github.com/TIBCOSoftware/flogo-contrib/activity/device-pin",
				"settings": {
				  "pin": "A1",
				  "digital": "true",
				  "value": "HIGH"
				}
			  }
			}
		  }
		]
	  }`
)

// init registers the command and flags
func init() {
	deviceCmd.AddCommand(deviceCreateCmd)
	deviceCreateCmd.Flags().StringVarP(&deviceJSONFile, "file", "f", "", "Specify the device.json to create device project from (optional)")
	deviceCreateCmd.Flags().StringVarP(&deviceCreateName, "name", "n", "", "The name of the device project (required)")
	deviceCreateCmd.MarkFlagRequired("name")
}

// runDeviceCreate is the actual execution of the command
func runDeviceCreate(cmd *cobra.Command, args []string) {
	var deviceJson string
	var err error

	if len(deviceJSONFile) != 0 {

		if fgutil.IsRemote(deviceJSONFile) {

			deviceJson, err = fgutil.LoadRemoteFile(deviceJSONFile)
			if err != nil {
				fmt.Printf("Error loading device file: '%s' - %s\n\n", deviceJSONFile, err.Error())
				os.Exit(2)
			}
		} else {
			deviceJson, err = fgutil.LoadLocalFile(deviceJSONFile)
			if err != nil {
				fmt.Printf("Error loading device file: '%s' - %s\n\n", deviceJSONFile, err.Error())
				os.Exit(2)
			}
		}
	} else {
		deviceJson = tplSimpleDevice
	}

	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error getting current directory: %s\n\n", err.Error())
		os.Exit(2)
	}

	deviceDir := filepath.Join(currentDir, deviceCreateName)

	err = device.CreateDevice(device.SetupNewDeviceProjectEnv(), deviceJson, deviceDir, deviceCreateName)
	if err != nil {
		fmt.Printf("Error while creating device project: %s\n\n", err.Error())
	}
}

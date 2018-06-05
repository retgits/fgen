// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/retgits/fgen/app"
	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// appCreateCmd represents the xxx command
var appCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Flogo application",
	Run:   runCreateApp,
}

// Flags
var (
	appCreateFLV    string
	appCreateVendor string
	appCreateFile   string
	appCreateName   string
)

// Variables
var (
	tplSimpleApp = `{
		"name": "AppName",
		"type": "flogo:app",
		"version": "0.0.1",
		"description": "My flogo application description",
		"triggers": [
		  {
			"id": "my_rest_trigger",
			"ref": "github.com/TIBCOSoftware/flogo-contrib/trigger/rest",
			"settings": {
			  "port": "9233"
			},
			"handlers": [
			  {
				"actionId": "my_simple_flow",
				"settings": {
				  "method": "GET",
				  "path": "/test"
				}
			  }
			]
		  }
		],
		"actions": [
		  {
			"id": "my_simple_flow",
			"name": "my simple flow",
			"ref": "github.com/TIBCOSoftware/flogo-contrib/action/flow",
			"data": {
			  "flow": {
				"name": "my simple flow",
				"attributes": [],
				"rootTask": {
				  "id": 1,
				  "type": 1,
				  "tasks": [
					{
					  "id": 2,
					  "type": 1,
					  "activityRef": "github.com/TIBCOSoftware/flogo-contrib/activity/log",
					  "name": "log",
					  "attributes": [
						{
						  "name": "message",
						  "value": "Simple Log",
						  "type": "string"
						}
					  ]
					}
				  ],
				  "links": [
				  ]
				}
			  }
			}
		  }
		]
	  }`
)

// init registers the command and flags
func init() {
	appCmd.AddCommand(appCreateCmd)
	appCreateCmd.Flags().StringVar(&appCreateFLV, "flv", "", " The flogo dependency constraints as comma separated value (for example github.com/TIBCOSoftware/flogo-lib@0.0.0,github.com/TIBCOSoftware/flogo-contrib@0.0.0)")
	appCreateCmd.Flags().StringVar(&appCreateFile, "file", "", "The flogo.json to create project from")
	appCreateCmd.Flags().StringVar(&appCreateVendor, "vendor", "", "Copy sources from an existing vendor directory")
	appCreateCmd.Flags().StringVar(&appCreateName, "name", "", "The name of the app (required)")
	appCreateCmd.MarkFlagRequired("name")
}

// runCreateApp is the actual execution of the command
func runCreateApp(cmd *cobra.Command, args []string) {
	var appJson string
	var appName string
	var err error

	if len(appCreateFile) != 0 {

		if fgutil.IsRemote(appCreateFile) {

			appJson, err = fgutil.LoadRemoteFile(appCreateFile)
			if err != nil {
				fmt.Printf("Error loading app file: '%s' - %s\n\n", appCreateFile, err.Error())
				os.Exit(2)
			}
		} else {
			appJson, err = fgutil.LoadLocalFile(appCreateFile)
			if err != nil {
				fmt.Printf("Error loading app file: '%s' - %s\n\n", appCreateFile, err.Error())
				os.Exit(2)
			}
		}
	} else {
		appName = appCreateName
		appJson = tplSimpleApp
	}

	currentDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	appDir := filepath.Join(currentDir, appName)

	err = app.CreateApp(app.SetupNewProjectEnv(), appJson, appDir, appName, appCreateVendor, appCreateFLV)
	if err != nil {
		fmt.Printf("Error while creating app project: %s\n\n", err.Error())
	}
}

// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/retgits/fgen/config"
	"github.com/spf13/cobra"
)

// appListCmd represents the xxx command
var appListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed contributions",
	Run:   runListApp,
}

// Flags
var (
	appListJSON bool
	appListType string
)

// Variables
var ()

// Constants
const (
	ctActions   = "actions"
	ctTriggers  = "triggers"
	ctActivites = "activities"
)

// init registers the command and flags
func init() {
	appCmd.AddCommand(appListCmd)
	appListCmd.Flags().StringVarP(&appListType, "type", "t", "", "The type of contribution you want to list (\"actions\"|\"triggers\"|\"activities\") (required)")
	appListCmd.Flags().BoolVarP(&appListJSON, "json", "j", false, "Generate output as json")
	appListCmd.MarkFlagRequired("type")
}

// runListApp is the actual execution of the command
func runListApp(cmd *cobra.Command, args []string) {
	var cType config.ContribType
	listCT := appListType

	switch listCT {
	case ctActions:
		cType = config.ACTION
	case ctTriggers:
		cType = config.TRIGGER
	case ctActivites:
		cType = config.ACTIVITY
	case "flow-models":
		cType = config.FLOW_MODEL
	default:
		fmt.Printf("Error: Unknown contribution type - %s\n\n", listCT)
		os.Exit(2)
	}

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	dependencies, err := app.ListDependencies(app.SetupExistingProjectEnv(appDir), cType)

	if err != nil {
		fmt.Printf("Error while getting dependencies: %s", err.Error())
		os.Exit(2)
	}

	if appListJSON {
		depJson, err := json.MarshalIndent(dependencies, "", "  ")
		if err != nil {
			fmt.Printf("Error while marshalling JSON: %s", err.Error())
			os.Exit(2)
		}
		fmt.Println(string(depJson))
	} else {
		byType := make(map[string][]string)

		//aggregate by ContribType
		for _, dependency := range dependencies {

			switch dependency.ContribType {
			case config.ACTION:
				byType[ctActions] = append(byType[ctActions], dependency.Ref)
			case config.TRIGGER:
				byType[ctTriggers] = append(byType[ctTriggers], dependency.Ref)
			case config.ACTIVITY:
				byType[ctActivites] = append(byType[ctActivites], dependency.Ref)
			default:
				byType[dependency.ContribType.String()] = append(byType[dependency.ContribType.String()], dependency.Ref)
			}
		}

		for ct, refs := range byType {

			fmt.Printf("%s:\n", ct)

			for _, ref := range refs {

				fmt.Printf("  %s\n", ref)
			}
		}
	}
}

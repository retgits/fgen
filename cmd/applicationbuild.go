// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/spf13/cobra"
)

// appBuildCmd represents the xxx command
var appBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a Flogo application",
	Run:   runBuildApp,
}

// Flags
var (
	appBuildOptimize       bool
	appBuildNoGeneration   bool
	appBuildGenerationOnly bool
	appBuildEmbedConfig    bool
	appBuildShim           string
	appBuildDocker         string
)

// Variables
var ()

// init registers the command and flags
func init() {
	appCmd.AddCommand(appBuildCmd)
	appBuildCmd.Flags().BoolVar(&appBuildOptimize, "o", false, "optimize build")
	appBuildCmd.Flags().BoolVar(&appBuildEmbedConfig, "e", false, "embed config")
	appBuildCmd.Flags().BoolVar(&appBuildNoGeneration, "nogen", false, "no generation")
	appBuildCmd.Flags().BoolVar(&appBuildGenerationOnly, "gen", false, "only generation")
	appBuildCmd.Flags().StringVar(&appBuildShim, "shim", "", "trigger shim")
	appBuildCmd.Flags().StringVar(&appBuildDocker, "docker", "", "build docker")
}

// runBuildApp is the actual execution of the command
func runBuildApp(cmd *cobra.Command, args []string) {
	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	// Validate exclusive params
	if appBuildGenerationOnly && appBuildNoGeneration {
		fmt.Printf("Error: nogen and gen flags are mutually exclusive, please choose just one\n\n")
		os.Exit(2)
	}

	options := &app.BuildOptions{SkipPrepare: false, NoGeneration: appBuildNoGeneration, GenerationOnly: appBuildGenerationOnly, BuildDocker: appBuildDocker, PrepareOptions: &app.PrepareOptions{OptimizeImports: appBuildOptimize, EmbedConfig: appBuildEmbedConfig, Shim: appBuildShim}}
	err = app.BuildApp(app.SetupExistingProjectEnv(appDir), options)
	if err != nil {
		fmt.Printf("Error while building app project: %s\n\n", err.Error())
	}
}

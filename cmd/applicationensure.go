// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"

	"github.com/retgits/fgen/app"
	"github.com/retgits/fgen/dep"
	"github.com/spf13/cobra"
)

// appEnsureCmd represents the xxx command
var appEnsureCmd = &cobra.Command{
	Use:   "ensure",
	Short: "Ensure gets a project into a complete, reproducible, and likely compilable state",
	Run:   runEnsureApp,
}

// Flags
var (
	ensureAdd        string
	ensureUpdate     bool
	ensureNoVendor   bool
	ensureVerbose    bool
	ensureVendorOnly bool
)

// Variables
var ()

// init registers the command and flags
func init() {
	appCmd.AddCommand(appEnsureCmd)
	appEnsureCmd.Flags().StringVar(&ensureAdd, "add", "", "add new dependencies, or populate Gopkg.toml with constraints for existing dependencies (default: false)")
	appEnsureCmd.Flags().BoolVar(&ensureUpdate, "update", false, "update the named dependencies (or all, if none are named) in Gopkg.lock to the latest allowed by Gopkg.toml (default: false)")
	appEnsureCmd.Flags().BoolVar(&ensureNoVendor, "no-vendor", false, " update Gopkg.lock (if needed), but do not update vendor/ (default: false)")
	appEnsureCmd.Flags().BoolVar(&ensureVerbose, "verbose", false, "enable verbose logging (default: false)")
	appEnsureCmd.Flags().BoolVar(&ensureVendorOnly, "vendor-only", false, "populate vendor/ from Gopkg.lock without updating it first (default: false)")
}

// runEnsureApp is the actual execution of the command
func runEnsureApp(cmd *cobra.Command, args []string) {

	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	// Create args
	ensureArgs := []string{}
	if len(ensureAdd) > 0 {
		ensureArgs = append(ensureArgs, "-add", ensureAdd)
	}
	if ensureUpdate {
		ensureArgs = append(ensureArgs, "-update")
	}
	if ensureVerbose {
		ensureArgs = append(ensureArgs, "-v")
	}
	if ensureNoVendor {
		ensureArgs = append(ensureArgs, "-no-vendor")
	} else if ensureVendorOnly {
		ensureArgs = append(ensureArgs, "-vendor-only")
	}

	depManager := dep.New(app.SetupExistingProjectEnv(rootDir))

	err = depManager.Ensure(ensureArgs...)
	if err != nil {
		fmt.Printf("Error while running dep ensure: %s\n\n", err.Error())
	}
}

// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	toml "github.com/pelletier/go-toml"
	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// pluginUninstallCmd represents the plugin command
var pluginUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall and remove a plugin from the Project Flogo CLI",
	Run:   runPluginUninstall,
}

// Flags
var (
	uninstallPluginRepoURL string
)

// Variables
var ()

// init registers the command and flags
func init() {
	pluginCmd.AddCommand(pluginUninstallCmd)
	pluginUninstallCmd.Flags().StringVarP(&uninstallPluginRepoURL, "repo", "r", "", "The repository of the plugin you want to uninstall (required)")
	pluginUninstallCmd.MarkFlagRequired("repo")
}

// runPluginUninstall is the actual execution of the command
func runPluginUninstall(cmd *cobra.Command, args []string) {
	fmt.Printf("\nUninstall plugin from the Project Flogo CLI\n")
	fmt.Printf("Using repository: %s\n", uninstallPluginRepoURL)

	rawRepoURL := strings.Replace(uninstallPluginRepoURL, "github.com", "raw.githubusercontent.com", 1)
	rawRepoURL = fmt.Sprintf("%s/master", rawRepoURL)

	// Get the plugin.toml content
	pluginTomlContent, err := fgutil.LoadRemoteFile(fmt.Sprintf("%s/plugin.toml", rawRepoURL))
	if err != nil {
		fmt.Printf("Error while getting plugin.toml: %s", err.Error())
		os.Exit(2)
	}

	// Load the content into a TOML tree
	config, err := toml.Load(pluginTomlContent)
	if err != nil {
		fmt.Printf("Error converting plugin.toml: %s\n", err.Error())
		os.Exit(2)
	}

	// Get the correct key
	queryResult := config.Get("plugin")
	if queryResult == nil {
		fmt.Printf("Unknown error occured, no plugin found\n")
		os.Exit(2)
	}

	// Get the plugin structure
	pluginMap := queryResult.([]*toml.Tree)[0].ToMap()

	// Get the GOPATH
	out, err := exec.Command("go", "env", "GOPATH").Output()
	if err != nil {
		log.Fatal(err)
	}
	gopath := strings.TrimSuffix(string(out), "\n")

	// Getting plugin files
	fmt.Printf("Found plugin: %s\n", pluginMap["name"])
	fmt.Printf("Remove files...\n")
	for _, filename := range pluginMap["files"].([]interface{}) {
		localPath := filepath.Join(gopath, goPathCLI, filename.(string))
		fgutil.DeleteFile(localPath)
	}

	// Install the new version of the commandline
	cmdExec := exec.Command("go", "install", "./...")
	cmdExec.Dir = filepath.Join(gopath, goPathCLI)
	cmdExec.Env = append(os.Environ())

	out, err = cmdExec.Output()
	if err != nil {
		fmt.Printf("Error while executing command: %s", err.Error())
		os.Exit(2)
	}

	fmt.Printf("Removed plugin!\n\n")
}

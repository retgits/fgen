// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/retgits/fgen/gen"
	"github.com/spf13/cobra"
)

// contribCreateCmd represents the generate command
var contribCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Generate a new contribution for Project Flogo",
	Run:   runContribCreate,
}

// Flags
var (
	genType string
	genName string
)

// Variables
var (
	generators = make(map[string]gen.CodeGenerator)
)

// init registers the command and flags
func init() {
	contribCmd.AddCommand(contribCreateCmd)
	contribCreateCmd.Flags().StringVarP(&genType, "type", "t", "", "The type of contribution you want to generate (\"action\"|\"activity\"|\"flowmodel\"|\"trigger\") (required)")
	contribCreateCmd.Flags().StringVarP(&genName, "name", "n", "", "The name you want to give your contribution (required)")
	contribCreateCmd.MarkFlagRequired("type")
	contribCreateCmd.MarkFlagRequired("name")
}

// runContribCreate is the actual execution of the command
func runContribCreate(cmd *cobra.Command, args []string) {
	contribution := genType
	name := genName

	generators["action"] = &gen.ActionGenerator{}
	generators["trigger"] = &gen.TriggerGenerator{}
	generators["activity"] = &gen.ActivityGenerator{}
	generators["flowmodel"] = &gen.FlowModelGenerator{}

	generator, exists := generators[contribution]

	if exists {

		data := struct {
			Name string
		}{
			name,
		}

		currentDir, _ := os.Getwd()
		basePath := path.Join(currentDir, name)

		if _, err := os.Stat(basePath); err == nil {
			fmt.Printf("Error: Cannot create project, directory '%s' already exists\n\n", name)
			os.Exit(2)
		}

		os.MkdirAll(basePath, 0777)

		err := generator.Generate(basePath, data)

		if err != nil {
			fmt.Printf("Error generating contribution: %s\n\n", err.Error())
			os.Exit(2)
		}

	} else {
		fmt.Printf("Error: unknown contribution type %q\n", contribution)
		fmt.Printf("Run 'flogo generate help' for a list available types\n\n")
	}
}

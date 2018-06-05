// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	toml "github.com/pelletier/go-toml"
	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// contribSearchCmd represents the generate command
var contribSearchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for Project Flogo contributions",
	Run:   runContribSearch,
}

// Flags
var (
	contribSearchType   string
	contribSearchString string
)

// Variables
var (
	filterByType   = false
	filterByString = false
)

// Constants
const (
	typeAll      = "all"
	typeActivity = "activity"
	typeTrigger  = "trigger"
	tomlURL      = "https://raw.githubusercontent.com/TIBCOSoftware/flogo/master/showcases/data/items.toml"
	tomlKey      = "items"
)

// init registers the command and flags
func init() {
	contribCmd.AddCommand(contribSearchCmd)
	contribSearchCmd.Flags().StringVarP(&contribSearchType, "type", "t", "", "The type you're looking for (\"all\"|\"activity\"|\"trigger\") (required)")
	contribSearchCmd.Flags().StringVarP(&contribSearchString, "string", "s", "", "The search string you want to use (optional)")
	//contribSearchCmd.MarkFlagRequired("type")
}

// runContribSearch is the actual execution of the command
func runContribSearch(cmd *cobra.Command, args []string) {
	// Break if the type is not known
	switch contribSearchType {
	case typeActivity,
		typeTrigger:
		filterByType = true
	case typeAll:
		filterByType = false
	default:
		fmt.Printf("Error: Unknown type - %s\n\n", contribSearchType)
	}

	if len(contribSearchString) > 0 {
		filterByString = true
	}

	if !filterByType && !filterByString {
		fmt.Printf("Error: Neither type or string flags are specified\n\n")
		os.Exit(2)
	}

	// Get the FAR content
	content, err := fgutil.LoadRemoteFile(tomlURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n\n", err.Error())
		os.Exit(2)
	}

	// Find artifacts
	datamap, err := searchContent(content, filterByType, filterByString, contribSearchType, contribSearchString)
	if err != nil {
		fmt.Printf("Error: %s\n\n", err.Error())
		os.Exit(2)
	}

	// Print a table
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Type", "Description", "URL", "Author"})

	for _, v := range datamap {
		table.Append(v)
	}

	table.Render()
}

func searchContent(content string, filterByType bool, filterByString bool, searchType string, searchString string) ([][]string, error) {
	// Read the content
	config, err := toml.Load(content)
	if err != nil {
		return nil, err
	}

	// Get the correct key
	queryResult := config.Get(tomlKey)
	if queryResult == nil {
		return nil, fmt.Errorf("Unknown error occured, no items found in FAR")
	}

	// Prepare a result structure
	resultArray := queryResult.([]*toml.Tree)
	datamap := make([][]string, 0)
	for _, val := range resultArray {
		tempVal := val.ToMap()
		if filterByType && filterByString {
			if containsKey(tempVal, "type", searchType) && containsValue(tempVal, searchString) {
				datamap = append(datamap, []string{tempVal["name"].(string), tempVal["type"].(string), tempVal["description"].(string), tempVal["url"].(string), tempVal["author"].(string)})
			}
		} else if filterByType {
			if containsKey(tempVal, "type", searchType) {
				datamap = append(datamap, []string{tempVal["name"].(string), tempVal["type"].(string), tempVal["description"].(string), tempVal["url"].(string), tempVal["author"].(string)})
			}
		} else if filterByString {
			if containsValue(tempVal, searchString) {
				datamap = append(datamap, []string{tempVal["name"].(string), tempVal["type"].(string), tempVal["description"].(string), tempVal["url"].(string), tempVal["author"].(string)})
			}
		}
	}

	return datamap, nil
}

func containsKey(datamap map[string]interface{}, key string, value string) bool {
	if _, ok := datamap[key]; ok {
		if datamap[key] == value {
			return true
		}
	}
	return false
}

func containsValue(datamap map[string]interface{}, value string) bool {
	for key := range datamap {
		if strings.Contains(datamap[key].(string), value) {
			return true
		}
	}
	return false
}

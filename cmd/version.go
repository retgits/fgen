// Package cmd defines and implements command-line commands and flags
// used by flogo. Commands and flags are implemented using Cobra.
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	toml "github.com/pelletier/go-toml"
	"github.com/retgits/fgen/fgutil"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display the current version of the CLI",
	Run:   runVersion,
}

// Types that represent the Gopkg.lock file
type rawLock struct {
	SolveMeta solveMeta          `toml:"solve-meta"`
	Projects  []rawLockedProject `toml:"projects"`
}

type solveMeta struct {
	InputsDigest    string `toml:"inputs-digest"`
	AnalyzerName    string `toml:"analyzer-name"`
	AnalyzerVersion int    `toml:"analyzer-version"`
	SolverName      string `toml:"solver-name"`
	SolverVersion   int    `toml:"solver-version"`
}

type rawLockedProject struct {
	Name     string   `toml:"name"`
	Branch   string   `toml:"branch,omitempty"`
	Revision string   `toml:"revision"`
	Version  string   `toml:"version,omitempty"`
	Source   string   `toml:"source,omitempty"`
	Packages []string `toml:"packages"`
}

// Constants
const (
	lockName = "Gopkg.lock"
)

// init registers the command and flags
func init() {
	rootCmd.AddCommand(versionCmd)
}

// runVersion is the actual execution of the command
func runVersion(cmd *cobra.Command, args []string) {
	cmdExec := exec.Command("git", "describe", "--tags")
	cmdExec.Dir = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "TIBCOSoftware", "flogo-cli")
	cmdExec.Env = append(os.Environ())

	out, err := cmdExec.Output()
	if err != nil {
		fmt.Printf("Error while executing command: %s", err.Error())
		os.Exit(2)
	}
	re := regexp.MustCompile("\\n")
	fc := re.ReplaceAllString(string(out), "")

	fmt.Printf("flogo cli version [%s]\n", fc)

	appDir, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error: Unable to determine working directory\n\n")
		os.Exit(2)
	}

	project := SetupExistingProjectEnv(appDir)

	config, err := toml.LoadFile(filepath.Join(project.GetAppDir(), lockName))

	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		raw := rawLock{}
		err := config.Unmarshal(&raw)
		if err != nil {
			fmt.Printf("Unable to parse the lock as TOML\n")
		}

		for _, v := range raw.Projects {
			if fgutil.CaseInsensitiveContains(v.Name, "flogo") {
				if v.Version == "" {
					fmt.Printf("Your project uses %s branch %s and revision %s\n", v.Name, v.Branch, v.Revision)
				} else {
					fmt.Printf("Your project uses %s version %s\n", v.Name, v.Version)
				}
			}
		}
	}
}

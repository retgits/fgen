package device

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SetupNewDeviceProjectEnv() Project {
	return NewPlatformIoProject()
}

func SetupExistingDeviceProjectEnv(appDir string) Project {

	project := NewPlatformIoProject()

	if err := project.Init(appDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing flogo device project: %s\n\n", err.Error())
		os.Exit(2)
	}

	if err := project.Open(); err != nil {
		fmt.Fprintf(os.Stderr, "Error opening flogo device project: %s\n\n", err.Error())
		os.Exit(2)
	}

	return project
}

func splitVersion(t string) (path string, version string) {

	idx := strings.LastIndex(t, "@")

	version = ""
	path = t

	if idx > -1 {
		v := t[idx+1:]

		if isValidVersion(v) {
			version = v
			path = t[0:idx]
		}
	}

	return path, version
}

//todo validate that "s" a valid semver
func isValidVersion(s string) bool {

	if s == "" {
		//assume latest version
		return true
	}

	if s[0] == 'v' && len(s) > 1 && isNumeric(string(s[1])) {
		return true
	}

	if isNumeric(string(s[0])) {
		return true
	}

	return false
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

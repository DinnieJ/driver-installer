package command

import (
	"fmt"
	"regexp"
)

type ChromeCmdRunner struct {
	Runner
}

var ERR_NO_VERSION_FOUND = fmt.Errorf("found no version in binary")

var ERR_NO_CHROME_BINARY_FOUND = fmt.Errorf("no chrome binary found")

func (c *ChromeCmdRunner) FetchChromeVersion() (string, error) {
	strOutput, err := c.Run("google-chrome-stable", "--version")
	if err != nil {
		return "", ERR_NO_CHROME_BINARY_FOUND
	}
	versionReg := regexp.MustCompile(`(?m)(\d*\.\d*\.\d*\.\d*)`)

	for _, match := range versionReg.FindAllString(string(strOutput), -1) {
		return match, nil
	}
	return "", ERR_NO_VERSION_FOUND
}

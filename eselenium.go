package seleniumembedded

import (
	"github.com/DinnieJ/selenium-embedded/internal/command"
	. "github.com/DinnieJ/selenium-embedded/internal/driver"
)

var NewChromeSeleniumDriver = func(folderPath string, arch string) *ChromeSeleniumDriver {
	chromeCmdRunner := &command.ChromeCmdRunner{Runner: &command.CommandRunner{}}
	version, err := chromeCmdRunner.FetchChromeVersion()
	if err != nil {
		return nil
	}

	// arch = fmt.Sprintf("%s_%s", runtime.GOOS, arch)
	return &ChromeSeleniumDriver{
		ChromeDriverDownloader: NewChromeDriver(version, folderPath, "linux64"),
		ChromeCmdRunner:        chromeCmdRunner,
	}
}

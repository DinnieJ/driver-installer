package seleniumembedded

import (
	"github.com/DinnieJ/driver-installer/internal/command"
	. "github.com/DinnieJ/driver-installer/internal/driver"
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
	}
}

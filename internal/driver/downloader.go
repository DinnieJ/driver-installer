package driver

import (
	"github.com/DinnieJ/driver-installer/internal/command"
)

type DriverInstaller interface {
	Install() error
	GetDriverPath() string
}

type ChromeSeleniumDriver struct {
	ChromeDriverDownloader
}

var chromeCmdRunner command.ChromeCmdRunner = command.ChromeCmdRunner{Runner: &command.CommandRunner{}}

func (c *ChromeSeleniumDriver) Install() error {
	if c.IsDriverDownloaded() {
		return nil
	}
	if err := c.DownloadDriver(); err != nil {
		return err
	}
	if err := c.UnzipDriver(); err != nil {
		return err
	}
	return nil
}

func (c *ChromeSeleniumDriver) GetDriverPath() string {
	return c.GetDriverPath()
}

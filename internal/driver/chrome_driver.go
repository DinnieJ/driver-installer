package driver

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"archive/zip"
)

type ChromeDriverDownloader interface {
	DownloadDriver() error
	UnzipDriver() error
	GetDriverPath() string
	IsDriverDownloaded() bool
}

type ChromeDriver struct {
	Version    string
	FolderPath string
	Arch       string

	driverPath string
}

const DRIVER_FILENAME = "chromedriver"
const DRIVER_DEFAULT_PATH = "/tmp/eselenium/"
const DRIVER_ZIP_NAME = "chromedriver_linux64.zip"

var ERROR_UNABLE_TO_DOWNLOAD_DRIVER = fmt.Errorf("unable to download driver")
var ERROR_UNABLE_TO_CREATE_NEW_DRIVER_FILE = fmt.Errorf("unable to create driver file")
var ERROR_FAILED_TO_DOWNLOAD_DRIVER = fmt.Errorf("failed to download driver")
var ERROR_FAILED_TO_OPEN_ZIP_FILE = fmt.Errorf("failed to open zip file")
var ERROR_UNABLE_TO_OPEN_DRIVER_FILE = fmt.Errorf("unable to open driver file")
var ERROR_UNABLE_TO_EXTRACT_DRIVER_FILE = fmt.Errorf("unable to extract driver file")

func NewChromeDriver(version string, folderPath string, arch string) *ChromeDriver {
	if folderPath == "" {
		folderPath = DRIVER_DEFAULT_PATH
	}
	return &ChromeDriver{
		Version:    version,
		FolderPath: folderPath,
		Arch:       arch,
		driverPath: filepath.Join(folderPath, DRIVER_FILENAME),
	}
}

func (c *ChromeDriver) DownloadDriver() error {
	downloadUrl := fmt.Sprintf("https://storage.googleapis.com/chrome-for-testing-public/%s/%s/chromedriver-%s.zip", c.Version, c.Arch, c.Arch)
	resp, err := http.Get(downloadUrl)
	if err != nil || resp.StatusCode != 200 {
		return ERROR_UNABLE_TO_DOWNLOAD_DRIVER
	}
	// driverDest := "/tmp/chromedriver_linux64.zip"
	f, err := os.CreateTemp(os.TempDir(), DRIVER_ZIP_NAME)
	if err != nil {
		return ERROR_UNABLE_TO_CREATE_NEW_DRIVER_FILE
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		return ERROR_FAILED_TO_DOWNLOAD_DRIVER
	}
	defer f.Close()
	defer resp.Body.Close()

	return nil
}

func (c *ChromeDriver) UnzipDriver() error {
	r, err := zip.OpenReader(filepath.Join(c.FolderPath, DRIVER_ZIP_NAME))
	if err != nil {
		return ERROR_FAILED_TO_OPEN_ZIP_FILE
	}
	for _, f := range r.File {
		if f.Name != DRIVER_FILENAME {
			continue
		}
		rc, err := f.Open()
		if err != nil {
			return ERROR_UNABLE_TO_OPEN_DRIVER_FILE
		}
		defer rc.Close()
		out, err := os.Create(c.driverPath)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, rc)
		if err != nil {
			return ERROR_UNABLE_TO_EXTRACT_DRIVER_FILE
		}
		break
	}
	return nil
}

func (c *ChromeDriver) GetDriverPath() string {
	return c.driverPath
}

func (c *ChromeDriver) IsDriverDownloaded() bool {
	_, err := os.Stat(c.driverPath)
	if err == nil || !os.IsNotExist(err) {
		return true
	}
	return false
}

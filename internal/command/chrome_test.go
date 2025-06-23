package command_test

import (
	"testing"

	"github.com/DinnieJ/selenium-embedded/internal/command"
	. "github.com/DinnieJ/selenium-embedded/internal/command/mock"
)

func TestGetVersionChromeSuccess(t *testing.T) {
	mockCmdRunner := new(MockCmdRunner)
	mockCmdRunner.On("Run", "google-chrome-stable", "--version").Return("1.2.3.4", nil)
	chromeCmdRunner := &command.ChromeCmdRunner{
		Runner: mockCmdRunner,
	}

	a, err := chromeCmdRunner.FetchChromeVersion()
	if err != nil {
		t.Fatal(err)
	}
	if a != "1.2.3.4" {
		t.Fatalf("expected 1.2.3.4, got %s", a)
	}
}

func TestGetVersionChromeWhenFindNoBinary(t *testing.T) {
	mockCmdRunner := new(MockCmdRunner)
	mockCmdRunner.On("Run", "google-chrome-stable", "--version").Return("", command.ERR_NO_CHOME_BINARY_FOUND)
	chromeCmdRunner := &command.ChromeCmdRunner{
		Runner: mockCmdRunner,
	}

	_, err := chromeCmdRunner.FetchChromeVersion()
	if err != command.ERR_NO_CHOME_BINARY_FOUND {
		t.Fatalf("expected ERR_NO_CHOME_BINARY_FOUND, got %s", err)
	}
}

func TestGetVersionChromeWhenVersionStringInvalid(t *testing.T) {
	mockCmdRunner := new(MockCmdRunner)
	mockCmdRunner.On("Run", "google-chrome-stable", "--version").Return("Invalid string", nil)
	chromeCmdRunner := &command.ChromeCmdRunner{
		Runner: mockCmdRunner,
	}

	_, err := chromeCmdRunner.FetchChromeVersion()
	if err != command.ERR_NO_VERSION_FOUND {
		t.Fatalf("expected ERR_NO_VERSION_FOUND, got %s", err)
	}
}

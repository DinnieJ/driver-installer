package command_test

import (
	"strings"
	"testing"

	"github.com/DinnieJ/selenium-embedded/internal/command"
)

func TestCommandRunnerSuccess(t *testing.T) {
	cmdRunner := new(command.CommandRunner)

	output, err := cmdRunner.Run("echo", "hello world !")
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}
	if strings.TrimSpace(output) != "hello world !" {
		t.Fatalf("expected 'hello world !', got %s", output)
	}
}

func TestCommandRunnerFail(t *testing.T) {
	cmdRunner := new(command.CommandRunner)
	output, err := cmdRunner.Run("nonexistent")
	if err == nil {
		t.Fatalf("expected error, got %s", output)
	}
}

package command

import (
	"os/exec"
)

type Runner interface {
	Run(args ...string) (string, error)
}

type CommandRunner struct{}

func (c *CommandRunner) Run(args ...string) (string, error) {

	command := exec.Command(args[0], args[1:]...)

	output, err := command.Output()

	if err != nil {
		return "", err
	}

	return string(output), nil
}

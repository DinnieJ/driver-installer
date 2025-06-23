package command

import "github.com/stretchr/testify/mock"

type MockCmdRunner struct {
	mock.Mock
}

func (c *MockCmdRunner) Run(args ...string) (string, error) {
	infArgs := make([]interface{}, len(args))
	for i, arg := range args {
		infArgs[i] = arg
	}
	r := c.Called(infArgs...)
	return r.String(0), r.Error(1)
}

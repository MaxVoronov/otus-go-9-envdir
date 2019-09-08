package service

import (
	"os"
	"os/exec"
	"strings"
)

type Executor struct {
	*exec.Cmd
}

func NewExecutor(name string, envList *EnvironmentList) *Executor {
	args := strings.Split(name, " ")
	executor := &Executor{exec.Command(args[0])}
	executor.Env = envList.Vars
	executor.Stdout = os.Stdout
	executor.Stderr = os.Stderr

	if len(args) > 1 {
		executor.Args = append(executor.Args, args[1:]...)
	}

	return executor
}

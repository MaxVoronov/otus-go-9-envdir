package cmd

import (
	"errors"
	"github.com/maxvoronov/otus-go-9-envdir/service"
	"os"
)

func Execute() error {
	args := os.Args
	if len(args) < 3 {
		return errors.New("you must pass 2 arguments")
	}

	envList := service.NewEnvironmentList()
	if err := envList.ReadFromDirectory(args[1]); err != nil {
		return err
	}

	executor := service.NewExecutor(args[2], envList)
	if err := executor.Run(); err != nil {
		return err
	}

	return nil
}

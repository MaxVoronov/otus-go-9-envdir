package service

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type EnvironmentList struct {
	Vars []string
}

func NewEnvironmentList() *EnvironmentList {
	return &EnvironmentList{}
}

func (e *EnvironmentList) ReadFromDirectory(dir string) error {
	absPath, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	files, err := ioutil.ReadDir(absPath)
	if err != nil {
		return err
	}

	e.Vars = make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		buf, err := ioutil.ReadFile(filepath.Join(absPath, file.Name()))
		if err != nil {
			return err
		}
		e.Vars = append(e.Vars, fmt.Sprintf("%s=%s", file.Name(), string(buf)))
	}

	return nil
}

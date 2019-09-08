package main

import (
	"fmt"
	"github.com/maxvoronov/otus-go-9-envdir/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	if err := createCommandsStructure().Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

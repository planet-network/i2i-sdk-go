package main

import (
	"fmt"

	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/spf13/cobra"
)

func cfgInit(cmd *cobra.Command, args []string) {
	application, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := application.Initialize(); err != nil {
		fail(err)
	}

	fmt.Println("done")
}

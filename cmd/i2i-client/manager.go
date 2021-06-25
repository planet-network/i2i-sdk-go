package main

import (
	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/spf13/cobra"
)

func managerQuickOrder(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

}

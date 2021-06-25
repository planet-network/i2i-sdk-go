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

func cfgSetActive(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	if err := appHandler.NodeSetDefault(args[0]); err != nil {
		fail(err)
	}
}

func cfgList(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	cfg := appHandler.Config()

	fmt.Println("nodes:")
	for k, _ := range cfg.Nodes {
		if k == cfg.SelectedNode {
			fmt.Printf(" -> %s\n", k)
		} else {
			fmt.Printf("    %s\n", k)
		}
	}

}

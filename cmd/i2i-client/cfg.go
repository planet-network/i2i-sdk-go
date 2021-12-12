package main

import (
	"fmt"
	"os"

	"github.com/planet-network/i2i-sdk-go/app"
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

func completion(cmd *cobra.Command, args []string) {
	switch args[0] {
	case "bash":
		if err := cmd.Root().GenBashCompletion(os.Stdout); err != nil {
			fail(err)
		}
	case "zsh":
		if err := cmd.Root().GenZshCompletion(os.Stdout); err != nil {
			fail(err)
		}
	case "fish":
		if err := cmd.Root().GenFishCompletion(os.Stdout, true); err != nil {
			fail(err)
		}
	case "powershell":
		if err := cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout); err != nil {
			fail(err)
		}
	}
}

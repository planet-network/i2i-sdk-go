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

	nodePlan, err := cmd.Flags().GetString(flagPlan)
	if err != nil {
		fail(err)
	}

	hostingAddr, err := cmd.Flags().GetString(flagHosting)
	if err != nil {
		fail(err)
	}

	nodeName, err := cmd.Flags().GetString(flagName)
	if err != nil {
		fail(err)
	}

	//initializeAs, err := cmd.Flags().GetString(flagInitializeAs)
	//if err != nil {
	//	fail(err)
	//}

	node := &app.Node{
		Meta: app.NodeMeta{
			Hosting: app.NodeHosting{
				Plan:           nodePlan,
				ManagerAddress: hostingAddr,
			},
		},
		Name: nodeName,
	}

	if err := appHandler.NodeCreateWithKeychain(node); err != nil {
		fail(err)
	}

}

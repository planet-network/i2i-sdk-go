package main

import (
	"fmt"

	"github.com/planet-platform/i2i-sdk-go/manager"
	"github.com/spf13/cobra"
)

func register(cmd *cobra.Command, args []string) {

}

func order(cmd *cobra.Command, args []string) {

}

func managerPlans(cmd *cobra.Command, args []string) {
	hostingAddr, err := cmd.Flags().GetString(flagHosting)
	if err != nil {
		fail(err)
	}

	managerClient := manager.NewClient(manager.ClientOpt{Address: hostingAddr})

	plans, err := managerClient.PlanList()
	if err != nil {
		fail(err)
	}

	fmt.Println(plans)
}

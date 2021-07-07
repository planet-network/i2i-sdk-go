package main

import (
	"fmt"

	"github.com/planet-platform/i2i-sdk-go/manager"

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

func cfgShow(cmd *cobra.Command, args []string) {
	var selectedNode string

	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	if len(args) > 0 {
		selectedNode = args[0]
	}

	config := appHandler.Config()
	if len(args) == 0 {
		selectedNode = config.SelectedNode
	}

	if selectedNode == "" {
		fail("active node not set")
	}

	node, ok := config.Nodes[selectedNode]
	if !ok {
		fail("node: ", selectedNode, " not found")
	}

	fmt.Println("Node:")
	fmt.Println("  name    :", selectedNode)
	fmt.Println("  address :", node.Meta.NodeAddress)
	fmt.Println("  acl     :", node.Meta.APIToken)
	fmt.Println("Hosting details:")
	fmt.Println("  manager      :", node.Meta.Hosting.ManagerAddress)
	fmt.Println("  plan         :", node.Meta.Hosting.Plan)
	fmt.Println("  unlock token :", node.Meta.Hosting.UnlockToken)
	fmt.Println("  client ID    :", node.Meta.Hosting.ClientID)
	fmt.Println("  password     :", node.Meta.Hosting.Password)
}

func cfgDelete(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fail("missing node name")
	}

	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	node, err := appHandler.NodeByName(args[0])
	if err != nil {
		fail(err)
	}

	if node.Meta.Hosting.ClientID != "" {
		client := manager.NewClient(manager.ClientOpt{
			Address:  node.Meta.Hosting.ManagerAddress,
			ClientID: node.Meta.Hosting.ClientID,
			Password: node.Meta.Hosting.Password,
		})

		if err := client.NodeRemove(); err != nil {
			fmt.Println(err)
		}

		if err := client.ClientRemove(); err != nil {
			fmt.Println(err)
		}
	}

	if err := appHandler.RemoveNode(args[0]); err != nil {
		fail(err)
	}

}

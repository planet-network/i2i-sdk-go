package main

import (
	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func connectionAdd(cmd *cobra.Command, args []string) {
	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	err = i2iClient.ConnectionAdd(args[0], args[1])
	if err != nil {
		fail(err)
	}
}

func connectionAddLocal(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	nodeA, err := appHandler.NodeWithKeychain(args[0])
	if err != nil {
		fail(err)
	}

	nodeB, err := appHandler.NodeWithKeychain(args[1])
	if err != nil {
		fail(err)
	}

	i2iClientA := client.New(client.Opt{
		Token:    nodeA.Meta.Hosting.UnlockToken,
		Address:  nodeA.Meta.NodeAddress,
		Acl:      nodeA.Meta.APIToken,
		Keychain: nodeA.Keychain,
	})

	i2iClientB := client.New(client.Opt{
		Token:    nodeB.Meta.Hosting.UnlockToken,
		Address:  nodeB.Meta.NodeAddress,
		Acl:      nodeB.Meta.APIToken,
		Keychain: nodeB.Keychain,
	})

	i2iClientA.Info()
	i2iClientB.Info()

}

func connectionList(cmd *cobra.Command, args []string) {
	profile, err := cmd.Flags().GetString(flagProfile)
	if err != nil {
		fail(err)
	}

	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	i2iInfo, err := i2iClient.ConnectionList(profile)
	if err != nil {
		fail(err)
	}

	printResult(i2iInfo)
}

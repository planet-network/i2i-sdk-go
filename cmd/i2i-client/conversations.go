package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func conversations(cmd *cobra.Command, args []string) {
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

	c, err := i2iClient.Conversations("")
	if err != nil {
		fail(err)
	}
	printResult(c)
}

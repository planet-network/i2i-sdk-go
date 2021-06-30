package main

import (
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func info(cmd *cobra.Command, args []string) {
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

	i2iInfo, err := i2iClient.Info()
	if err != nil {
		fail(err)
	}

	printResult(i2iInfo)
}

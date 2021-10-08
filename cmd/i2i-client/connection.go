package main

import (
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func connectionAdd(cmd *cobra.Command, args []string) {

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

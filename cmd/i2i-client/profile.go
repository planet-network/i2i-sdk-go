package main

import (
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func profileList(cmd *cobra.Command, args []string) {
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

	list, err := i2iClient.ProfileList()
	if err != nil {
		fail(err)
	}

	printResult(list)
}

func profileAdd(cmd *cobra.Command, args []string) {
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

	profile, err := i2iClient.DMeProfileAdd(&client.DMeProfileInput{ProfileName: args[0]})
	if err != nil {
		fail(err)
	}

	printResult(profile)
}

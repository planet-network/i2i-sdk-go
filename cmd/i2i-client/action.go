package main

import (
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func actionList(cmd *cobra.Command, args []string) {
	actionType, err := cmd.Flags().GetString(flagType)
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

	if actionType == "connection-request" {
		actions, err := i2iClient.FriendRequests()
		if err != nil {
			fail(err)
		}
		printResult(actions)
		return
	}

	actions, err := i2iClient.InteractiveActions()
	if err != nil {
		fail(err)
	}

	printResult(actions)
}

func actionUpdate(cmd *cobra.Command, args []string) {
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

	var action client.InterActiveAction

	switch args[0] {
	case "accept", "ACCEPT":
		action = client.InterActiveActionAccept
	case "deny", "DENY":
		action = client.InterActiveActionDeny
	default:
		fail("invalid action")
	}

	err = i2iClient.InteractiveActionUpdate(args[0], action)
	if err != nil {
		fail(err)
	}
}

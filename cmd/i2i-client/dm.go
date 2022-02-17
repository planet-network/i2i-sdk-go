package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func dmView(cmd *cobra.Command, args []string) {
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

	c, err := i2iClient.DirectMessage(&client.MessageViewInput{
		Conversation: args[0],
		Count:        100,
	})
	if err != nil {
		fail(err)
	}
	printResult(c)
}

func dmSend(cmd *cobra.Command, args []string) {
	reply, err := cmd.Flags().GetString(flagReply)
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

	input := &client.DirectMessageInput{
		Destination: args[0],
		Content:     args[1],
		Reply:       nil,
		Attachments: nil,
	}

	if reply != "" {
		input.Reply = &reply
	}

	c, err := i2iClient.SendDirectMessage(input)
	if err != nil {
		fail(err)
	}
	printResult(c)
}

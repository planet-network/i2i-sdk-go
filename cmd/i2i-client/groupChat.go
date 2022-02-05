package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func groupList(cmd *cobra.Command, args []string) {
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

	list, err := i2iClient.GroupChatList("")
	if err != nil {
		fail(err)
	}

	printResult(list)
}

func groupCreate(cmd *cobra.Command, args []string) {
	var (
		name       = args[0]
		profile    = ""
		publicKey  = client.RandomString32()
		privateKey = client.RandomString32()
	)

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

	chat, err := i2iClient.GroupChatCreate(&client.GroupchatInput{
		Profile:          &profile,
		GroupDisplayName: &name,
		ChatPublicKey:    &publicKey,
		ChatPrivateKey:   &privateKey,
	})
	if err != nil {
		fail(err)
	}

	printResult(chat)
}

func groupAddParticipant(cmd *cobra.Command, args []string) {
	var (
		chatID      = args[0]
		participant = args[1]
	)

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

	chat, err := i2iClient.GroupChatAddUser(&client.GroupchatAddUser{
		ChatPublicKey:    chatID,
		UserSignatureKey: []*string{&participant},
	})
	if err != nil {
		fail(err)
	}

	printResult(chat)

}

func groupLeave(cmd *cobra.Command, args []string) {
	var (
		chatID = args[0]
	)

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

	id, err := i2iClient.GroupChatLeave(chatID)
	if err != nil {
		fail(err)
	}

	printResult(id)
}

func groupMessageShow(cmd *cobra.Command, args []string) {
	var (
		chatID = args[0]
	)

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

	id, err := i2iClient.GroupChat(&client.MessageViewInput{
		Conversation: chatID,
		Count:        200,
	})
	if err != nil {
		fail(err)
	}

	printResult(id)
}

func groupMessageSend(cmd *cobra.Command, args []string) {
	var (
		chatID  = args[0]
		content = args[1]
	)

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

	id, err := i2iClient.GroupSendMessage(&client.GroupMessageInput{
		Destination: chatID,
		Content:     content,
	})
	if err != nil {
		fail(err)
	}

	printResult(id)

}

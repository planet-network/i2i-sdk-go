package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func dmeInfo(cmd *cobra.Command, args []string) {
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

	dmeInfo, err := i2iClient.DMeInfo()
	if err != nil {
		fail(err)
	}

	printResult(dmeInfo)
}

func dmeUpdate(cmd *cobra.Command, args []string) {
	firstName, err := cmd.Flags().GetString(flagFirstName)
	if err != nil {
		fail(err)
	}

	surname, err := cmd.Flags().GetString(flagSurname)
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

	input := &client.DMeInput{}
	if firstName != "" {
		input.FirstName = &firstName
	}
	if surname != "" {
		input.Surname = &surname
	}

	err = i2iClient.DMeUpdate(input)
	if err != nil {
		fail(err)
	}
}

package main

import (
	"github.com/planet-platform/i2i-sdk-go/app"

	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func aclList(cmd *cobra.Command, args []string) {
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

	aclList, err := i2iClient.AclList()
	if err != nil {
		fail(err)
	}

	printResult(aclList)
}

func aclAdd(cmd *cobra.Command, args []string) {
	scope, err := cmd.Flags().GetBool(flagPrivateScope)
	if err != nil {
		fail(err)
	}

	uuid, err := cmd.Flags().GetString(flagUUID)
	if err != nil {
		fail(err)
	}

	aclName, err := cmd.Flags().GetString(flagName)
	if err != nil {
		fail(err)
	}

	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	config := appHandler.Config()
	if config.SelectedNode == "" {
		fail("no active node set")
	}

	node, err := appHandler.NodeDefaultWithKeychain()
	if err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      node.Meta.APIToken,
		Keychain: node.Keychain,
	})

	input := &client.ACLInput{
		UUID:           uuid,
		Name:           &aclName,
		PrivatePlScope: &scope,
	}

	acl, err := i2iClient.AclAdd(input)
	if err != nil {
		fail(err)
	}

	node.Meta.APIToken = acl.Authorization

	if err := appHandler.NodeUpdate(node); err != nil {
		fail(err)
	}

	printResult(acl)
}

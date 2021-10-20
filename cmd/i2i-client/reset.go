package main

import (
	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func reset(cmd *cobra.Command, args []string) {
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

	err = i2iClient.Reset()
	if err != nil {
		fail(err)
	}

	node.Meta.APIToken = ""
	if err := appHandler.NodeUpdate(node); err != nil {
		fail(err)
	}
}

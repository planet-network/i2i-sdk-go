package main

import (
	"os"
	"runtime"

	"github.com/planet-network/i2i-sdk-go/app"
	"github.com/planet-network/i2i-sdk-go/client"
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

func aclRemove(cmd *cobra.Command, args []string) {
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

	aclList, err := i2iClient.AclRemove(args[0])
	if err != nil {
		fail(err)
	}

	printResult(aclList)
}

func aclAdd(cmd *cobra.Command, args []string) {
	var (
		applicationName = "i2i-client-go"
		version         = client.Version
		osName          = runtime.GOOS
	)

	scope, err := cmd.Flags().GetString(flagPrivateScope)
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
		ApplicationName:    &applicationName,
		ApplicationVersion: &version,
		OsName:             &osName,
	}

	hostName, err := os.Hostname()
	if err == nil {
		input.DeviceName = &hostName
	}

	if scope != "" {
		input.PrivatePlScopeName = &scope
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

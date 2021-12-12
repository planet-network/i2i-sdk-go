package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func plVerify(cmd *cobra.Command, args []string) {
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

	report, err := i2iClient.PlVerify()
	if err != nil {
		fail(err)
	}

	printResult(report)
}

func relationsList(cmd *cobra.Command, args []string) {
	var (
		filter *client.RelationFilterInput
		scope  string
	)

	as, err := cmd.Flags().GetString(flagAs)
	if err != nil {
		fail(err)
	}

	if len(args) > 0 {
		scope = args[0]
	}

	if as != "" || scope != "" {
		filter = &client.RelationFilterInput{}
	}

	if as != "" {
		filter.As = &as
	}

	if scope != "" {
		filter.Scope = &scope
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

	relations, err := i2iClient.PlRelationList(filter)
	if err != nil {
		fail(err)
	}

	printResult(relations)
}

func instancesList(cmd *cobra.Command, args []string) {
	var (
		filter *client.InstanceFilterInput
		scope  string
	)

	as, err := cmd.Flags().GetString(flagAs)
	if err != nil {
		fail(err)
	}

	if len(args) > 0 {
		scope = args[0]
	}

	if as != "" || scope != "" {
		filter = &client.InstanceFilterInput{}
	}

	if as != "" {
		filter.As = &as
	}

	if scope != "" {
		filter.Scope = &scope
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

	instances, err := i2iClient.PlInstanceList(filter)
	if err != nil {
		fail(err)
	}

	printResult(instances)
}

func scopeList(cmd *cobra.Command, args []string) {
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

	scopes, err := i2iClient.PlScopeList()
	if err != nil {
		fail(err)
	}

	printResult(scopes)
}

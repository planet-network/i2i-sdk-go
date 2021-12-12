package main

import (
	"strings"

	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func initialize(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fail("missing initialization type")
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

	as := strings.ToUpper(args[0])
	if err := i2iClient.Initialize(as); err != nil {
		fail(err)
	}
}

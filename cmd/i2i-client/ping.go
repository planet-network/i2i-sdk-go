package main

import (
	"fmt"

	"github.com/planet-platform/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func ping(cmd *cobra.Command, args []string) {
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

	if err := i2iClient.Ping(); err != nil {
		fail(err)
	}

	fmt.Println("pong")
}

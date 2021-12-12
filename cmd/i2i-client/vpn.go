package main

import (
	"fmt"

	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func vpnCreate(cmd *cobra.Command, args []string) {
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

	config, err := i2iClient.VnfWireguardCreate(&client.WireguardConfigInput{Name: args[0]})
	if err != nil {
		fail(err)
	}

	printResult(config)
}

func vpnStart(cmd *cobra.Command, args []string) {
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

	config, err := i2iClient.VnfWireguardStart(args[0])
	if err != nil {
		fail(err)
	}

	printResult(config)
}

func vpnStop(cmd *cobra.Command, args []string) {
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

	config, err := i2iClient.VnfWireguardStop(args[0])
	if err != nil {
		fail(err)
	}

	printResult(config)
}

func vpnCreatePeerConfig(cmd *cobra.Command, args []string) {
	wgFormat, err := cmd.Flags().GetBool(flagWireguardFormat)
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

	config, err := i2iClient.VnfWireguardCreatePeerConfig(&client.WireguardPeerInput{
		NetworkName: args[0],
		PeerName:    args[1],
	})
	if err != nil {
		fail(err)
	}

	if wgFormat {
		makeWgConfig(config)
	} else {
		printResult(config)
	}

}

func makeWgConfig(c *client.WireguardPeerConfig) {
	tpl := `
[Interface]
Address = %s
PrivateKey = %s
DNS = 1.1.1.1

[Peer]
PublicKey = %s
Endpoint = %s
AllowedIPs = %s
`

	fmt.Printf(tpl,
		c.Address,
		c.PrivateKey,
		c.PeerPublicKey,
		c.Endpoint,
		c.AllowedIps,
	)
}

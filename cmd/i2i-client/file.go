package main

import (
	"github.com/planet-network/i2i-sdk-go/client"
	"github.com/spf13/cobra"
)

func fileUpload(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.FileUpload(args[0])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

func fileStat(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.File(args[0])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

func fileRename(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.FileRename(args[0], args[1])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

func fileDownload(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.FileDownload(args[0], args[1])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

func fileList(cmd *cobra.Command, args []string) {
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

	fileList, err := i2iClient.FileList()
	if err != nil {
		fail(err)
	}
	printResult(fileList)
}

func fileRemove(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.FileRemove(args[0])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

func fileTransfer(cmd *cobra.Command, args []string) {
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

	fileInfo, err := i2iClient.FileTransfer(args[0], args[1])
	if err != nil {
		fail(err)
	}
	printResult(fileInfo)
}

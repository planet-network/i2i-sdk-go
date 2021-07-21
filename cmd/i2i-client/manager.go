package main

import (
	"fmt"
	"time"

	"github.com/planet-platform/i2i-sdk-go/client"

	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/planet-platform/i2i-sdk-go/manager"
	"github.com/spf13/cobra"
)

func managerQuickOrder(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	nodePlan, err := cmd.Flags().GetString(flagPlan)
	if err != nil {
		fail(err)
	}

	hostingAddr, err := cmd.Flags().GetString(flagHosting)
	if err != nil {
		fail(err)
	}

	password, err := cmd.Flags().GetString(flagPassword)
	if err != nil {
		fail(err)
	}

	nodeName, err := cmd.Flags().GetString(flagName)
	if err != nil {
		fail(err)
	}

	//initializeAs, err := cmd.Flags().GetString(flagInitializeAs)
	//if err != nil {
	//	fail(err)
	//}

	node := &app.Node{
		Meta: app.NodeMeta{
			Type: app.NodeTypeHosted,
			Hosting: app.NodeHosting{
				Plan:           nodePlan,
				ManagerAddress: hostingAddr,
			},
		},
		Name: nodeName,
	}

	managerClient := manager.NewClient(manager.ClientOpt{Address: hostingAddr})
	fmt.Println("[] Registering new client")
	clientID, err := managerClient.Register(password)
	if err != nil {
		fail(err)
	}
	node.Meta.Hosting.Password = password
	node.Meta.Hosting.ClientID = clientID
	node.Meta.Hosting.UnlockToken = client.RandomString32()

	fmt.Println("[] Log into manager with client ID: ", clientID)
	if err := managerClient.Login(); err != nil {
		fail(err)
	}

	fmt.Println("[] Ordering i2i")
	reply, err := managerClient.NodeOrder(&manager.NodeOrderRequest{
		Token: node.Meta.Hosting.UnlockToken,
		Plan:  nodePlan,
	})
	if err != nil {
		fail(err)
	}
	fmt.Println("[] New node ordered", reply.OrderID)

	if err := managerClient.WaitNodeProvisioned(time.Second * 3); err != nil {
		fail(err)
	}

	nodeDetails, err := managerClient.NodeShow()
	if err != nil {
		fail(err)
	}

	fmt.Println("[] Node acquired", nodeDetails.ApiAddress)

	node.Meta.NodeAddress = nodeDetails.ApiAddress
	node.Meta.APIToken = nodeDetails.Token

	if err := appHandler.NodeCreateWithKeychain(node); err != nil {
		fail(err)
	}

	i2iClient := client.New(client.Opt{
		Token:    node.Meta.Hosting.UnlockToken,
		Address:  node.Meta.NodeAddress,
		Acl:      "",
		Keychain: node.Keychain,
	})

	fmt.Println("[] Unlocking node")
	if err := i2iClient.Unlock(); err != nil {
		fail(err)
	}

}

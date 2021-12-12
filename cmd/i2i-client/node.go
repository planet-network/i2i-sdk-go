package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/planet-network/i2i-sdk-go/client"

	"github.com/planet-network/i2i-sdk-go/manager"

	"github.com/planet-network/i2i-sdk-go/app"
	"github.com/spf13/cobra"
)

func nodeList(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	cfg := appHandler.Config()

	fmt.Println("nodes:")
	for k, node := range cfg.Nodes {
		if k == cfg.SelectedNode {
			fmt.Printf(" * %s\n", nodeString(node))
		} else {
			fmt.Printf("   %s\n", nodeString(node))
		}
	}
}

func nodeInfo(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	config := appHandler.Config()

	if config.SelectedNode == "" {
		fail("active node not set")
	}

	node, ok := config.Nodes[config.SelectedNode]
	if !ok {
		fail("node: ", config.SelectedNode, " not found")
	}

	fmt.Println("Node:")
	fmt.Println("  name    :", config.SelectedNode)
	fmt.Println("  address :", node.Meta.NodeAddress)
	fmt.Println("  acl     :", node.Meta.APIToken)

	if node.Meta.Hosting.ManagerAddress != "" {
		fmt.Println("Hosting details:")
		fmt.Println("  manager      :", node.Meta.Hosting.ManagerAddress)
		fmt.Println("  plan         :", node.Meta.Hosting.Plan)
		fmt.Println("  unlock token :", node.Meta.Hosting.UnlockToken)
		fmt.Println("  client ID    :", node.Meta.Hosting.ClientID)
		fmt.Println("  password     :", node.Meta.Hosting.Password)
	}
}

func nodeAdd(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	if err := appHandler.NodeAddNoKeychain(&app.Node{
		Name:        args[0],
		HasKeychain: false,
		Meta: app.NodeMeta{
			Type:        "manual",
			NodeAddress: args[1],
		},
	}); err != nil {
		fail(err)
	}
}

func nodeExec(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	nodeName, err := cmd.Flags().GetString(flagName)
	if err != nil {
		fail(err)
	}

	i2iPath, err := cmd.Flags().GetString(flagI2iPath)
	if err != nil {
		fail(err)
	}

	port, err := cmd.Flags().GetInt(flagPort)
	if err != nil {
		fail(err)
	}

	if nodeName == "" {
		fail("missing node name")
	}

	if i2iPath == "" {
		fail("missing i2i path")
	}

	if !appHandler.NodeExist(nodeName) {
		err := appHandler.NodeCreateWithKeychain(&app.Node{
			Name: nodeName,
			Meta: app.NodeMeta{
				NodeAddress: fmt.Sprintf("0.0.0.0:%d", port),
				Type:        app.NodeTypeLocal,
				LocalExec: app.NodeLocalExec{
					I2IPath: i2iPath,
				},
			},
		})
		if err != nil {
			fail(err)
		}
	}

	executor, err := appHandler.Executor(nodeName, port)
	if err != nil {
		fail(err)
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := executor.Run(); err != nil {
			fail(err)
		}
	}()

	<-sigs
	if err := executor.Stop(); err != nil {
		fail(err)
	}
}

func nodeOrder(cmd *cobra.Command, args []string) {
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

func nodeRemove(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fail("missing node name")
	}

	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	node, err := appHandler.NodeByName(args[0])
	if err != nil {
		fail(err)
	}

	if node.Meta.Hosting.ClientID != "" {
		client := manager.NewClient(manager.ClientOpt{
			Address:  node.Meta.Hosting.ManagerAddress,
			ClientID: node.Meta.Hosting.ClientID,
			Password: node.Meta.Hosting.Password,
		})

		if err := client.Login(); err != nil {
			fmt.Println(err)
		}

		if err := client.NodeRemove(); err != nil {
			fmt.Println(err)
		}

		if err := client.ClientRemove(); err != nil {
			fmt.Println(err)
		}
	}

	if err := appHandler.RemoveNode(args[0]); err != nil {
		fail(err)
	}
}

func nodeSetActive(cmd *cobra.Command, args []string) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	if err := appHandler.NodeSetDefault(args[0]); err != nil {
		fail(err)
	}
}

func nodeShow(cmd *cobra.Command, args []string) {
	var selectedNode string

	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	if len(args) > 0 {
		selectedNode = args[0]
	}

	config := appHandler.Config()
	if len(args) == 0 {
		selectedNode = config.SelectedNode
	}

	if selectedNode == "" {
		fail("active node not set")
	}

	node, ok := config.Nodes[selectedNode]
	if !ok {
		fail("node: ", selectedNode, " not found")
	}

	fmt.Println("Node:")
	fmt.Println("  name    :", selectedNode)
	fmt.Println("  address :", node.Meta.NodeAddress)
	fmt.Println("  acl     :", node.Meta.APIToken)
	fmt.Println("Hosting details:")
	fmt.Println("  manager      :", node.Meta.Hosting.ManagerAddress)
	fmt.Println("  plan         :", node.Meta.Hosting.Plan)
	fmt.Println("  unlock token :", node.Meta.Hosting.UnlockToken)
	fmt.Println("  client ID    :", node.Meta.Hosting.ClientID)
	fmt.Println("  password     :", node.Meta.Hosting.Password)
}

func nodeString(n *app.Node) string {
	typeStr := fmt.Sprintf("[%s]", n.Meta.Type)
	return fmt.Sprintf("%-9s %s", typeStr, n.Name)
}

func nodeUpdate(cmd *cobra.Command, args []string) {
	node, err := activeNode()
	if err != nil {
		fail(err)
	}

	if node.Meta.Hosting.ClientID == "" {
		fail("missing client id")
	}

	managerClient := manager.NewClient(manager.ClientOpt{
		Address:  node.Meta.Hosting.ManagerAddress,
		ClientID: node.Meta.Hosting.ClientID,
		Password: node.Meta.Hosting.Password,
	})

	if err := managerClient.NodeUpdate(); err != nil {
		fail(err)
	}

}

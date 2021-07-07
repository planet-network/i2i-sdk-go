package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/planet-platform/i2i-sdk-go/app"
	"github.com/spf13/cobra"
)

func execute(cmd *cobra.Command, args []string) {
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

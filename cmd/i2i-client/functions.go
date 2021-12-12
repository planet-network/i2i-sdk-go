package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/planet-network/i2i-sdk-go/app"
)

// fail prints message and exists with code 1
func fail(msg ...interface{}) {
	fmt.Println(msg...)
	os.Exit(1)
}

func activeNode() (*app.Node, error) {
	appHandler, err := app.NewApp()
	if err != nil {
		fail(err)
	}

	if err := appHandler.LoadConfig(); err != nil {
		fail(err)
	}

	config := appHandler.Config()
	if config.SelectedNode == "" {
		return nil, fmt.Errorf("no active node set")
	}

	node, err := appHandler.NodeDefaultWithKeychain()
	if err != nil {
		fail(err)
	}

	return node, nil
}

func printResult(r interface{}) {
	data, _ := json.Marshal(r)
	fmt.Println(string(data))
}

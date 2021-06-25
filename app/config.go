package app

import (
	"encoding/json"
	"io/ioutil"
)

// Config contains i2i parameters
type Config struct {
	SelectedNode string           `json:"selected_node"`
	Nodes        map[string]*Node `json:"address"`
}

type Node struct {
	Name           string `json:"name"`
	Plan           string `json:"plan"`
	ManagerAddress string `json:"manager_address"`
	NodeAddress    string `json:"node_address"`
	UnlockToken    string `json:"unlock_token"`
	APIToken       string `json:"api_token"`
}

func loadConfig(filename string, config *Config) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, config)
}

func (c *Config) Store(filename string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0600)
}

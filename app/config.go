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
	ManagerAddress string `json:"manager_address"`
	NodeAddress    string `json:"node_address"`
	UnlockToken    string `json:"unlock_token"`
	APIToken       string `json:"api_token"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(data, config)

	return config, err
}

func (c *Config) Store(filename string) error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0600)
}

package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/planet-platform/i2i-sdk-go/client"
)

// Config contains i2i parameters
type Config struct {
	SelectedNode string           `json:"selected_node"`
	Nodes        map[string]*Node `json:"nodes"`
}

type NodeHosting struct {
	Plan           string `json:"plan"`
	ManagerAddress string `json:"manager_address"`
	UnlockToken    string `json:"unlock_token"`
	ClientID       string `json:"client_id"`
	Password       string `json:"password"`
}

type Node struct {
	// unique name of the node
	Name string `json:"name"`
	// HasKeychain describes if node has keychain set
	HasKeychain bool `json:"has_keychain"`
	// keychain of the i2i node
	Keychain *client.FullKeychain `json:"-"`
	// i2i metadata, storable on the disk
	Meta NodeMeta `json:"meta"`
}

type NodeMeta struct {
	Hosting     NodeHosting `json:"hosting"`
	NodeAddress string      `json:"node_address"`
	APIToken    string      `json:"api_token"`
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

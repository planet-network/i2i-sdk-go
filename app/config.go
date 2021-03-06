package app

import (
	"encoding/json"
	"io/ioutil"

	"github.com/planet-network/i2i-sdk-go/client"
)

// Config contains i2i parameters
type Config struct {
	SelectedNode    string              `json:"selected_node"`
	SelectedManager string              `json:"selected_manager"`
	Nodes           map[string]*Node    `json:"nodes"`
	Managers        map[string]*Manager `json:"managers"`
}
type Manager struct {
	Address  string `json:"address"`
	Token    string `json:"token"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type NodeHosting struct {
	Plan           string `json:"plan"`
	ManagerAddress string `json:"manager_address"`
	UnlockToken    string `json:"unlock_token"`
	ClientID       string `json:"client_id"`
	Password       string `json:"password"`
}

type NodeLocalExec struct {
	I2IPath string `json:"i2i_path"`
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

const (
	NodeTypeLocal  = "local"
	NodeTypeHosted = "hosted"
)

type NodeMeta struct {
	LocalExec   NodeLocalExec `json:"local_exec"`
	Hosting     NodeHosting   `json:"hosting"`
	NodeAddress string        `json:"node_address"`
	APIToken    string        `json:"api_token"`
	Type        string        `json:"type"`
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

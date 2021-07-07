package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/planet-platform/i2i-sdk-go/client"
)

const (
	i2iClientDir   = ".i2i-client"
	nodesDirName   = "nodes"
	configFilename = "config"
)

type App struct {
	config         Config
	workdir        string
	configFilePath string
	nodesDataPath  string
}

func NewApp() (*App, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	return &App{
		config:         Config{Nodes: map[string]*Node{}},
		workdir:        filepath.Join(homeDir, i2iClientDir),
		nodesDataPath:  filepath.Join(homeDir, i2iClientDir, nodesDirName),
		configFilePath: filepath.Join(homeDir, i2iClientDir, configFilename),
	}, nil
}

func (a *App) Initialize() error {
	if err := os.Mkdir(a.workdir, 0700); err != nil {
		return err
	}

	if err := os.Mkdir(a.nodesDataPath, 0700); err != nil {
		return err
	}

	return a.config.Store(a.configFilePath)
}

func (a *App) LoadConfig() error {
	return loadConfig(a.configFilePath, &a.config)
}

func (a *App) Config() Config {
	return a.config
}

func (a *App) WriteConfig() error {
	return a.config.Store(a.configFilePath)
}

func (a *App) NodeUpdate(node *Node) error {
	a.config.Nodes[node.Name] = node
	return a.config.Store(a.configFilePath)
}

func (a *App) keychainPath(name string) string {
	return filepath.Join(a.workdir, nodesDirName, name, "keychain.i2i")
}

func (a *App) keychainSourceFile(name string) string {
	return fmt.Sprintf("file://%s", a.keychainPath(name))
}

func (a *App) workdirPath(name string) string {
	return filepath.Join(a.workdir, nodesDirName, name)
}

func (a *App) NodeCreateWithKeychain(node *Node) error {
	if _, ok := a.config.Nodes[node.Name]; ok {
		return fmt.Errorf("already exist")
	}

	nodeDir := a.workdirPath(node.Name)
	if err := os.Mkdir(nodeDir, 0700); err != nil {
		return err
	}

	keychain, err := client.GenerateFullKeychain()
	if err != nil {
		return err
	}

	if err := keychain.SaveToFileSafe(a.keychainPath(node.Name)); err != nil {
		return err
	}
	node.Keychain = keychain
	node.HasKeychain = true

	a.config.Nodes[node.Name] = node

	if a.config.SelectedNode == "" {
		a.config.SelectedNode = node.Name
	}

	return a.config.Store(a.configFilePath)
}

func (a *App) NodeSetDefault(name string) error {
	if _, ok := a.config.Nodes[name]; !ok {
		return fmt.Errorf("not found")
	}
	a.config.SelectedNode = name
	return a.config.Store(a.configFilePath)
}

func (a *App) NodeDefaultWithKeychain() (*Node, error) {
	node, ok := a.config.Nodes[a.config.SelectedNode]
	if !ok {
		return nil, fmt.Errorf("node %q not found", a.config.SelectedNode)
	}

	if node.HasKeychain {
		keychain, err := client.LoadFullKeychainFromFile(a.keychainPath(a.config.SelectedNode))
		if err != nil {
			return nil, err
		}
		node.Keychain = keychain
	}

	return node, nil
}

func (a *App) NodeByName(name string) (*Node, error) {
	node, ok := a.config.Nodes[name]
	if !ok {
		return nil, fmt.Errorf("node %q not found", a.config.SelectedNode)
	}

	if node.HasKeychain {
		keychain, err := client.LoadFullKeychainFromFile(a.keychainPath(a.config.SelectedNode))
		if err != nil {
			return nil, err
		}
		node.Keychain = keychain
	}

	return node, nil
}

func (a *App) NodeExist(name string) bool {
	_, ok := a.config.Nodes[name]
	return ok
}

func (a *App) RemoveNode(name string) error {
	_, ok := a.config.Nodes[name]
	if !ok {
		return fmt.Errorf("node %q not found", a.config.SelectedNode)
	}

	if a.config.SelectedNode == name {
		a.config.SelectedNode = ""
	}

	delete(a.config.Nodes, name)
	if a.config.SelectedNode == name {
		a.config.SelectedNode = ""
	}

	if err := a.WriteConfig(); err != nil {
		return err
	}

	return os.RemoveAll(a.workdirPath(name))
}

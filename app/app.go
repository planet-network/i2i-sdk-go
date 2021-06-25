package app

import (
	"fmt"
	"os"
	"path/filepath"
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

func (a *App) nodePut(node *Node) error {
	a.config.Nodes[node.Name] = node
	return a.config.Store(a.configFilePath)
}

func (a *App) NodeSetDefault(name string) error {
	if _, ok := a.config.Nodes[name]; !ok {
		return fmt.Errorf("not found")
	}
	a.config.SelectedNode = name
	return a.config.Store(a.configFilePath)
}

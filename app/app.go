package app

import (
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

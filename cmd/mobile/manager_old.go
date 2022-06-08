package mobile

import (
	"crypto/rand"
	"fmt"
	"github.com/planet-network/i2i-sdk-go/manager"
)

type Manager struct {
	client *manager.Client
}

func NewManager() *Manager {
	return &Manager{
		client: manager.NewClient(manager.ClientOpt{
			Address: "https://v2.vphi.io",
		}),
	}
}

func NewManagerAuth(clientID, password string) *Manager {
	return &Manager{
		client: manager.NewClient(manager.ClientOpt{
			Address:  "https://v2.vphi.io",
			ClientID: clientID,
			Password: password,
		}),
	}
}

// Register registers client to i2i-manager, it returns string- client ID, and error if any
func (m *Manager) Register(password string) (string, error) {
	return m.client.Register(password)
}

func (m *Manager) Login() error {
	return m.client.Login()
}

func GenerateToken() string {
	var key [32]byte
	_, _ = rand.Read(key[:])
	return fmt.Sprintf("%x", key[:])
}

func (m *Manager) NodeOder(token string) error {
	_, err := m.client.NodeOrder(&manager.NodeOrderRequest{
		Token: token,
		Plan:  "infinite",
	})

	return err
}

type Node struct {
	ID         string
	ApiAddress string
	Plan       string
	Token      string
	ValidUntil int64
	CreatedAt  int64
	Live       bool
}

func (m *Manager) NodeShow() (*Node, error) {
	node, err := m.client.NodeShow()
	if err != nil {
		return nil, err
	}

	return &Node{
		ID:         node.ID,
		ApiAddress: node.ApiAddress,
		Plan:       node.Plan,
		Token:      node.Token,
		ValidUntil: node.ValidUntil,
		CreatedAt:  node.CreatedAt,
		Live:       node.Live,
	}, nil
}

func (m *Manager) NodeRemove() error {
	return m.client.NodeRemove()
}

func (m *Manager) ClientRemove() error {
	return m.client.ClientRemove()
}

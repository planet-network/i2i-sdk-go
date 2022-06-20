package pc

import (
	"crypto/rand"
	"fmt"
	"github.com/planet-network/i2i-sdk-go/manager"
	"strconv"
	"time"
)

type Manager struct {
	restClient    *RestClient
	managerClient *manager.Client
}

func NewManager(r *RestClient) *Manager {
	return &Manager{restClient: r}
}

func (m *Manager) NodeOrder(project string) error {
	var (
		password = "1234123412341234qwertyuiop"
	)

	if _, err := m.restClient.TableList(); err != nil {
		return err
	}

	m.managerClient = manager.NewClient(manager.ClientOpt{
		Address: "https://v2.vphi.io",
	})

	if _, err := m.managerClient.Register(password); err != nil {
		return err
	}

	if err := m.managerClient.Login(); err != nil {
		return err
	}

	token := GenerateToken()

	if _, err := m.managerClient.NodeOrder(&manager.NodeOrderRequest{
		Token: token,
		Plan:  "infinite",
	}); err != nil {
		return err
	}

	time.Sleep(time.Second * 3)

	node, err := m.managerClient.NodeShow()
	if err != nil {
		return err
	}

	table := fmt.Sprintf("app.i2i.%s", project)

	if err := m.restClient.DataAdd(table, "plan", node.Plan); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "address", node.ApiAddress); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "id", node.ID); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "token", node.Token); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "live", strconv.FormatBool(node.Live)); err != nil {
		return err
	}

	return nil
}

func GenerateToken() string {
	var key [32]byte
	_, _ = rand.Read(key[:])
	return fmt.Sprintf("%x", key[:])
}

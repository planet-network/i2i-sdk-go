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

func (m *Manager) NodeOrder(project string, password string) error {
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

	if err := m.restClient.DataAdd(table, "client_id", m.managerClient.ClientID()); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "password", m.managerClient.Password()); err != nil {
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

	if err := m.restClient.DataAdd(table, "valid_until", strconv.FormatInt(node.ValidUntil, 10)); err != nil {
		return err
	}

	return nil
}

func (m *Manager) NodeGet(project string) (*manager.CustomerNode, error) {
	table := fmt.Sprintf("app.i2i.%s", project)

	clientID, err := m.restClient.DataGet(table, "client_id")
	if err != nil {
		return nil, err
	}

	password, err := m.restClient.DataGet(table, "password")
	if err != nil {
		return nil, err
	}

	if m.managerClient == nil {
		m.managerClient = manager.NewClient(manager.ClientOpt{
			Address:  "https://v2.vphi.io",
			ClientID: string(clientID.Value),
			Password: string(password.Value),
		})

		if err := m.managerClient.Login(); err != nil {
			return nil, err
		}
	}

	plan, err := m.restClient.DataGet(table, "plan")
	if err != nil {
		return nil, err
	}

	address, err := m.restClient.DataGet(table, "address")
	if err != nil {
		return nil, err
	}

	id, err := m.restClient.DataGet(table, "id")
	if err != nil {
		return nil, err
	}

	token, err := m.restClient.DataGet(table, "token")
	if err != nil {
		return nil, err
	}

	liveRaw, err := m.restClient.DataGet(table, "live")
	if err != nil {
		return nil, err
	}
	live, err := strconv.ParseBool(string(liveRaw.Value))
	if err != nil {
		return nil, err
	}

	validUntilRaw, err := m.restClient.DataGet(table, "valid_until")
	if err != nil {
		return nil, err
	}

	validUntil, err := strconv.ParseInt(string(validUntilRaw.Value), 10, 64)
	if err != nil {
		return nil, err
	}

	return &manager.CustomerNode{
		ID:         string(id.Value),
		ApiAddress: string(address.Value),
		Plan:       string(plan.Value),
		Token:      string(token.Value),
		ValidUntil: validUntil,
		Live:       live,
	}, nil
}

func (m *Manager) NodeDelete(project string) error {
	table := fmt.Sprintf("app.i2i.%s", project)

	clientID, err := m.restClient.DataGet(table, "client_id")
	if err != nil {
		return err
	}

	password, err := m.restClient.DataGet(table, "password")
	if err != nil {
		return err
	}

	if m.managerClient == nil {
		m.managerClient = manager.NewClient(manager.ClientOpt{
			Address:  "https://v2.vphi.io",
			ClientID: string(clientID.Value),
			Password: string(password.Value),
		})

		if err := m.managerClient.Login(); err != nil {
			return err
		}
	}

	if err := m.managerClient.NodeRemove(); err != nil {
		return err
	}

	if err := m.managerClient.ClientRemove(); err != nil {
		return err
	}

	if err := m.restClient.DataAdd(table, "deleted", strconv.FormatBool(true)); err != nil {
		return err
	}

	return nil
}

func GenerateToken() string {
	var key [32]byte
	_, _ = rand.Read(key[:])
	return fmt.Sprintf("%x", key[:])
}

package mobile

import (
	"context"
	"fmt"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/planet-network/i2i-manager/client"
	"github.com/planet-network/i2i-manager/client/auth"
	"github.com/planet-network/i2i-manager/client/plan"
	"github.com/planet-network/i2i-manager/client/user"
	"github.com/planet-network/i2i-manager/models"
	"time"
)

type Manager struct {
	client *client.I2iManager
	jwt    string
}

func NewManager() *Manager {

	return &Manager{
		client: client.NewHTTPClient(nil),
	}
}

func (m *Manager) Register(login string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err := m.client.Auth.PostV1UserRegister(&auth.PostV1UserRegisterParams{
		User: &models.ManagerUserRegistrationRequest{
			Password: login,
			Username: password,
		},
		Context: ctx,
	})

	return err
}

func (m *Manager) Login(login string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := m.client.Auth.PostV1UserLogin(&auth.PostV1UserLoginParams{
		User: &models.ManagerClientLoginRequest{
			Password: login,
			Username: password,
		},
		Context: ctx,
	})
	if err != nil {
		return err
	}

	cfg := client.DefaultTransportConfig()
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	transport.DefaultAuthentication = httptransport.BearerToken(response.Payload.Jwt)
	m.client.SetTransport(transport)

	m.jwt = response.Payload.Jwt

	return nil
}

//not working yet
func (m *Manager) PlanList() error {
	response, err := m.client.Plan.GetClientPlanList(plan.NewGetClientPlanListParams())
	if err != nil {
		return err
	}

	for _, i := range response.Payload {
		fmt.Println(i)
	}

	return nil
}

func (m *Manager) UserRemove() error {
	_, err := m.client.User.PostClientRemove(user.NewPostClientRemoveParams())
	return err
}

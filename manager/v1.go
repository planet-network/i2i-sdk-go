package manager

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

type ClientV1 struct {
	c   *client.I2iManager
	jwt string
}

func NewClientV1() *ClientV1 {

	return &ClientV1{
		c: client.NewHTTPClient(nil),
	}
}

func (c *ClientV1) Register(login string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	_, err := c.c.Auth.PostV1UserRegister(&auth.PostV1UserRegisterParams{
		User: &models.ManagerUserRegistrationRequest{
			Password: login,
			Username: password,
		},
		Context: ctx,
	})

	return err
}

func (c *ClientV1) Login(login string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	response, err := c.c.Auth.PostV1UserLogin(&auth.PostV1UserLoginParams{
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
	c.c.SetTransport(transport)

	c.jwt = response.Payload.Jwt

	return nil
}

func (c *ClientV1) PlanList() error {
	response, err := c.c.Plan.PostClientPlanList(plan.NewPostClientPlanListParams())
	if err != nil {
		return err
	}

	for _, i := range response.Payload {
		fmt.Println(i)
	}

	return nil
}

func (c *ClientV1) UserRemove() error {
	_, err := c.c.User.PostClientRemove(user.NewPostClientRemoveParams())
	return err
}

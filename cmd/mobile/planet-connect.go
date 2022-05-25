package mobile

import (
	"fmt"
	"github.com/planet-network/planet-connect/client"
	"github.com/planet-network/planet-connect/models"
	"time"
)

type PCClient struct {
	client *client.Client
}

func New() *PCClient {
	return &PCClient{
		client: client.New(),
	}
}

func (c *PCClient) Connect() error {
	if err := c.client.Connect("wss://pc.vlow.me/api"); err != nil {
		return err
	}

	go c.client.ReceiveLoop(32)

	return nil
}

func (c *PCClient) Register(login string, secret string, method string) error {
	id, err := c.client.Register(login, secret, method)
	if err != nil {
		return err
	}

	_, err = c.clientWaitReply(id, models.TypeRegister)
	return err
}

func (c *PCClient) Login(login string, secret string, method string) error {
	id, err := c.client.Login(login, secret, method)
	if err != nil {
		return err
	}

	_, err = c.clientWaitReply(id, models.TypeLogin)
	return err
}

func (c *PCClient) Logout() error {
	_, err := c.client.Logout()
	if err != nil {
		return err
	}

	return err
}

func (c *PCClient) PersonalDataAdd(key string, value string) error {
	id, err := c.client.PersonalDataAdd(&client.PersonalDataInput{
		Key:      []byte(key),
		Value:    []byte(value),
		DataFlag: client.DataFlagPrivate,
	})

	if err != nil {
		return err
	}

	_, err = c.clientWaitReply(id, models.TypePersonalDataAdd)
	return err
}

func (c *PCClient) PersonalDataGet(key string) (string, error) {
	id, err := c.client.PersonalDataGet([]byte(key))

	if err != nil {
		return "", err
	}

	reply, err := c.clientWaitReply(id, models.TypePersonalDataGet)
	if err != nil {
		return "", err
	}

	parsed, err := c.client.Parse(reply)
	if err != nil {
		return "", err
	}

	return string(parsed.Data.Value), nil
}

func (c *PCClient) clientWaitReply(id string, replyType models.CallType) (*models.Response, error) {
	var (
		timer = time.NewTimer(time.Minute * 15)
		r     = &models.Response{}
	)

	select {
	case <-timer.C:
		return nil, fmt.Errorf("reply timeout")
	case r = <-c.client.C:
	}

	if r.Code != 200 {
		return nil, fmt.Errorf("call failed: code: %d, err: %s", r.Code, r.Error)
	}

	if r.ID != id {
		return nil, fmt.Errorf("reply id=%s, expected: %s", r.ID, id)
	}

	if r.Type != replyType {
		return nil, fmt.Errorf("reply type=%s, expected: %s", r.Type.String(), replyType.String())
	}

	return r, nil
}

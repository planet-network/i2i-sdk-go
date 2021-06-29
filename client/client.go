package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	token      string
	address    string
	acl        string
	keychain   *FullKeychain
}

type Opt struct {
	Token    string
	Address  string
	Acl      string
	Keychain *FullKeychain
}

func New(opt Opt) *Client {
	return &Client{
		token:      opt.Token,
		address:    opt.Address,
		acl:        opt.Acl,
		keychain:   opt.Keychain,
		httpClient: &http.Client{},
	}
}

func (c *Client) SetKeychain(k *FullKeychain) {
	c.keychain = k
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) nodeAddress() string {
	return fmt.Sprintf("http://%s", c.address)
}

func (c *Client) nodeStateAddress() string {
	return fmt.Sprintf("http://%s/state", c.address)
}

func (c *Client) nodeGraphqlAddress() string {
	return fmt.Sprintf("http://%s/query", c.address)
}

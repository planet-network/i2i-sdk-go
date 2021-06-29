package client

import "fmt"

type Client struct {
	token    string
	address  string
	acl      string
	keychain *FullKeychain
}

type Opt struct {
	Token    string
	Address  string
	Acl      string
	Keychain *FullKeychain
}

func New(opt Opt) *Client {
	return &Client{
		token:    opt.Token,
		address:  opt.Address,
		acl:      opt.Acl,
		keychain: opt.Keychain,
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

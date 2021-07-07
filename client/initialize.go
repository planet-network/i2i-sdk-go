package client

import (
	"time"
)

func (c *Client) Initialize(as string) error {
	var response interface{}

	_, err := c.query(&query{
		query:     mutationInitialize,
		variables: map[string]interface{}{"type": as},
		timeout:   time.Second * 5,
		response:  response,
	})

	if err != nil {
		return err
	}

	return nil
}

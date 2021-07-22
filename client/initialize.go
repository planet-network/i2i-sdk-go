package client

import (
	"time"
)

// Initialize is function which must be called in a first place.
// It will determine the role of i2i and must be done carefully, because choice is immutable.
// Following arguments are acceptable:
// - DME
// - DORG
// - SUPERNODE
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

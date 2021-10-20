package client

import "time"

func (c *Client) Reset() error {
	_, err := c.query(&query{
		query:   mutationReset,
		timeout: time.Second * 2,
	})
	return err
}

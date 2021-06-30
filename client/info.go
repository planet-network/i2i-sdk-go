package client

import (
	"time"
)

func (c *Client) Info() (*Info, error) {
	response := struct {
		Info Info `json:"info"`
	}{}

	_, err := c.query(&query{
		query:     queryInfo,
		variables: nil,
		timeout:   time.Second,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Info, nil
}

package client

import "time"

func (c *Client) DMeInfo() (*DMeInfo, error) {
	response := struct {
		DMeInfo DMeInfo `json:"dMeInfo"`
	}{}

	_, err := c.query(&query{
		query:     queryDMeInfo,
		variables: nil,
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.DMeInfo, nil
}

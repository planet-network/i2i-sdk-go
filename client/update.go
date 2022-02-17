package client

import "time"

// DMeUpdate updates the Dme
func (c *Client) DMeUpdate(input *DMeInput) error {

	_, err := c.query(&query{
		query:     mutationDMeUpdate,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  nil,
	})
	return err
}

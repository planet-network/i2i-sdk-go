package client

import "time"

// DMeUpdate updates the Dme
func (c *Client) DMeUpdate(firstName string) error {
	var data = struct {
		FirstName string `json:"first_name"`
	}{
		FirstName: firstName,
	}
	_, err := c.query(&query{
		query:     mutationDMeUpdate,
		variables: map[string]interface{}{"input": data},
		timeout:   time.Second * 2,
		response:  &data,
	})
	return err
}
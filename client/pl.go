package client

import "time"

// PlDataRead reads data on remote i2i
func (c *Client) PlDataRead(input *PlDataReadInput) ([]string, error) {
	response := struct {
		Values []string `json:"plDataRead"`
	}{}

	_, err := c.query(&query{
		query:     queryPlDataRead,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Values, nil
}

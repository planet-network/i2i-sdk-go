package client

import "time"

func (c *Client) MixinGetID() (string, error) {
	response := struct {
		ID string `json:"mixinGetID"`
	}{}

	_, err := c.query(&query{
		query:     queryMixingGetId,
		variables: nil,
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return "", err
	}

	return response.ID, nil
}

func (c *Client) MixinSetID(id string) error {
	_, err := c.query(&query{
		query:     mutationMixinSetId,
		variables: map[string]interface{}{"id": id},
		timeout:   time.Second * 2,
		response:  nil,
	})

	return err
}

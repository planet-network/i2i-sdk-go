package client

import (
	"io/ioutil"
	"net/http"
)

func (c *Client) State() (string, error) {
	response, err := c.httpClient.Get(c.nodeStateAddress())
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", ErrHttpWithCode(response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

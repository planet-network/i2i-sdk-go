package client

import "time"

func (c *Client) VnfWireguardCreatePeerConfig(input *WireguardPeerInput) (*WireguardPeerConfig, error) {
	response := struct {
		Config WireguardPeerConfig `json:"vnfWireguardCreatePeerConfig"`
	}{}

	_, err := c.query(&query{
		query:     mutationVnfWireguardCreatePeerConfig,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 5,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Config, nil
}

func (c *Client) VnfWireguardStart(name string) (*WireguardConfig, error) {
	response := struct {
		Config WireguardConfig `json:"vnfWireguardStart"`
	}{}

	_, err := c.query(&query{
		query:     mutationVnfWireguardStart,
		variables: map[string]interface{}{"input": name},
		timeout:   time.Second * 10,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Config, nil
}

func (c *Client) VnfWireguardStop(name string) (*WireguardConfig, error) {
	response := struct {
		Config WireguardConfig `json:"vnfWireguardStop"`
	}{}

	_, err := c.query(&query{
		query:     mutationVnfWireguardStop,
		variables: map[string]interface{}{"input": name},
		timeout:   time.Second * 10,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Config, nil
}

func (c *Client) VnfWireguardCreate(input *WireguardConfigInput) (*WireguardConfig, error) {
	response := struct {
		Config WireguardConfig `json:"vnfWireguardCreate"`
	}{}

	_, err := c.query(&query{
		query:     mutationVnfWireguardCreate,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 10,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Config, nil
}

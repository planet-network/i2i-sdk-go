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

func (c *Client) PlScopeList() ([]string, error) {
	response := struct {
		Scopes []string `json:"plScopeList"`
	}{}

	_, err := c.query(&query{
		query:    queryPlScopeList,
		timeout:  time.Second * 2,
		response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Scopes, nil
}

func (c *Client) PlInstanceList(filter *InstanceFilterInput) ([]*Instance, error) {
	response := struct {
		Instances []*Instance `json:"plInstances"`
	}{}

	_, err := c.query(&query{
		query:     queryPlInstances,
		variables: map[string]interface{}{"filter": filter},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Instances, nil
}

func (c *Client) PlRelationList(filter *RelationFilterInput) ([]*Relation, error) {
	response := struct {
		Relations []*Relation `json:"plRelations"`
	}{}

	_, err := c.query(&query{
		query:     queryPlRelations,
		variables: map[string]interface{}{"filter": filter},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Relations, nil
}

func (c *Client) PlVerify() ([]*PlReport, error) {
	response := struct {
		Report []*PlReport `json:"plVerify"`
	}{}

	_, err := c.query(&query{
		query:    queryPlVerify,
		timeout:  time.Second * 2,
		response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Report, nil
}

func (c *Client) PlInstance(scope, id string) (*Instance, error) {
	response := struct {
		Instance *Instance `json:"plInstance"`
	}{}

	_, err := c.query(&query{
		query:     queryPlInstance,
		variables: map[string]interface{}{"scope": scope, "id": id},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Instance, nil
}

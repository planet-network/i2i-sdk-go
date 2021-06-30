package client

import (
	"time"
)

func (c *Client) AclAdd(aclInput *ACLInput) (*ACL, error) {
	response := struct {
		AclAdd ACL `json:"aclAdd"`
	}{}

	_, err := c.query(&query{
		query:     mutationAclAdd,
		variables: map[string]interface{}{"input": aclInput},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.AclAdd, nil
}

func (c *Client) AclList() ([]*ACL, error) {
	response := struct {
		AclList []*ACL `json:"aclList"`
	}{}

	_, err := c.query(&query{
		query:    queryAclList,
		timeout:  time.Second * 2,
		response: &response,
	})

	if err != nil {
		return nil, err
	}

	return response.AclList, nil
}

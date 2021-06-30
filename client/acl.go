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

package client

import (
	"time"
)

// AclAdd creates new access token, which is used to authorize application in i2i.
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

// AclList lists acl tokens known by i2i
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

func (c *Client) AclRemove(id string) (string, error) {
	response := struct {
		ID string `json:"aclRemove"`
	}{}

	_, err := c.query(&query{
		query:     mutationAclRemove,
		variables: map[string]interface{}{"id": id},
		timeout:   time.Second * 2,
		response:  &response,
	})

	if err != nil {
		return "", err
	}

	return response.ID, nil
}

package client

import "time"

func (c *Client) GroupChatList(profile string) ([]*GroupChat, error) {
	response := struct {
		GroupChats []*GroupChat `json:"groupChatList"`
	}{}

	_, err := c.query(&query{
		query:     queryGroupChatList,
		variables: map[string]interface{}{"profile": profile},
		timeout:   2 * time.Second,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.GroupChats, nil
}

func (c *Client) GroupChatCreate(input *GroupchatInput) (*GroupChat, error) {
	response := struct {
		GroupChat *GroupChat `json:"groupChatCreate"`
	}{}

	_, err := c.query(&query{
		query:     mutationGroupChatCreate,
		variables: map[string]interface{}{"input": input},
		timeout:   5 * time.Second,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.GroupChat, nil
}

func (c *Client) GroupChatAddUser(input *GroupchatAddUser) (*GroupChat, error) {
	response := struct {
		GroupChats *GroupChat `json:"addUserToGroupChat"`
	}{}

	_, err := c.query(&query{
		query:     mutationAddUserToGroupChat,
		variables: map[string]interface{}{"input": input},
		timeout:   5 * time.Second,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.GroupChats, nil
}

func (c *Client) GroupChatLeave(id string) (string, error) {
	response := struct {
		GroupChat string `json:"groupChatLeave"`
	}{}

	_, err := c.query(&query{
		query:     mutationGroupChatLeave,
		variables: map[string]interface{}{"input": id},
		timeout:   5 * time.Second,
		response:  &response,
	})

	if err != nil {
		return "", err
	}

	return response.GroupChat, nil
}

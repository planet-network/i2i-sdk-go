package client

import "time"

func (c *Client) SendDirectMessage(input *DirectMessageInput) (*DirectMessage, error) {
	response := struct {
		Message DirectMessage `json:"sendDirectMessage"`
	}{}

	_, err := c.query(&query{
		query:     mutationSendDirectMessage,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 5,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Message, nil
}

func (c *Client) DirectMessage(input *MessageViewInput) (*DirectMessagePage, error) {
	response := struct {
		Messages DirectMessagePage `json:"directMessage"`
	}{}

	_, err := c.query(&query{
		query:     queryDirectMessage,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 10,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return &response.Messages, nil
}

func (c *Client) Conversations(profile string) ([]*Conversation, error) {
	response := struct {
		Messages []*Conversation `json:"conversations"`
	}{}

	_, err := c.query(&query{
		query:     queryConversations,
		variables: map[string]interface{}{"profile": profile},
		timeout:   time.Second * 10,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Messages, nil
}

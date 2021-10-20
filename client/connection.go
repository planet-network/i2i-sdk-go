package client

import (
	"time"
)

func (c *Client) ConnectionAdd(profile string, publicKey string) error {
	input := ConnectionInput{
		Profile:   []string{profile},
		PublicKey: publicKey,
	}
	var response interface{}
	_, err := c.query(&query{
		query:     mutationAddConnection,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  &response,
	})
	return err
}

func (c *Client) FriendRequests() ([]FriendRequest, error) {
	response := struct {
		FriendRequests []FriendRequest `json:"interactiveActions"`
	}{}

	_, err := c.query(&query{
		query:     interactiveActionsFriendRequestQry,
		variables: nil,
		timeout:   time.Second * 2,
		response:  &response,
	})

	return response.FriendRequests, err
}

func (c *Client) InteractiveActions() ([]InteractiveAction, error) {
	response := struct {
		InteractiveActions []InteractiveAction `json:"interactiveActions"`
	}{}

	_, err := c.query(&query{
		query:     interactiveActionsQry,
		variables: nil,
		timeout:   time.Second * 2,
		response:  &response,
	})

	return response.InteractiveActions, err
}

type InterActiveAction uint

const (
	InterActiveActionAccept InterActiveAction = iota
	InterActiveActionDeny
)

func (i InterActiveAction) String() string {
	switch i {
	case InterActiveActionAccept:
		return "ACCEPT"
	case InterActiveActionDeny:
		return "DENY"
	default:
		return ""
	}
}

func (c *Client) InteractiveActionUpdate(id string, action InterActiveAction) error {
	input := NotificationAction{
		ID:     id,
		Action: action.String(),
	}

	_, err := c.query(&query{
		query:     mutationInteractiveAction,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  nil,
	})

	return err
}

func (c *Client) ConnectionList(profile string) ([]*Connection, error) {
	response := struct {
		Connections []*Connection `json:"connectionList"`
	}{}

	_, err := c.query(&query{
		query:     queryConnectionList,
		variables: map[string]interface{}{"profile": profile},
		timeout:   2 * time.Second,
		response:  &response,
	})

	if err != nil {
		return nil, err
	}

	return response.Connections, nil
}

package client

import (
	"errors"
	"time"
)

type ConnectionInput struct {
	//  ID is the database ID of the connection
	ID string `json:"ID"`
	// profile contains id of the profiles to which connection is related
	Profile []string `json:"profile"`
	// public_key if set will change public key of the connection
	PublicKey string `json:"public_key"`
	// signature_key if set will change signature key of the connection
	SignatureKey string `json:"signature_key,omitempty"`
	// display_name if set will change display name of the connection
	DisplayName string `json:"display_name"`
	// name if set will change name of the connection
	Name string `json:"name"`
	// surname if set will change surname of the connection
	Surname string `json:"surname"`
	// country if set will change country of the connection
	Country string `json:"country"`
	//  for internal usage
	Transactions string `json:"transactions"`
}

type FriendRequest struct {
	// id is the database id of the object
	ID string `json:"id"`
	// id is the database id of the object
	Source string `json:"source"`
	// full_name is complete set of entity name making request
	FullName string `json:"full_name"`
	// time is
	Time time.Time `json:"time"`
}

type InteractiveActions struct {
	FriendRequests []FriendRequest `json:"interactiveActions"`
}

type NotificationAction struct {
	ID     string `json:"id"`
	Action string `json:"action"`
}

func (c *Client) AddConnection(profile string, publickey string) error {
	input := ConnectionInput{
		Profile:   []string{profile},
		PublicKey: publickey,
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

func (c *Client) QueryFriendRequests() (string, error) {
	var response InteractiveActions
	_, err := c.query(&query{
		query:     interactiveActionsQry,
		variables: nil,
		timeout:   time.Second * 2,
		response:  &response,
	})
	if err != nil {
		return "", err
	}
	if len(response.FriendRequests) == 1 {
		return response.FriendRequests[0].ID, nil
	}
	if len(response.FriendRequests) > 1 {
		return "", errors.New("multiple friend requests")
	}
	return "", errors.New("no friend requests")
}

func (c *Client) InterActiveActionAccept(id string) error {
	var response interface{}
	input := NotificationAction{
		ID:     id,
		Action: "ACCEPT",
	}
	_, err := c.query(&query{
		query:     mutationInteractiveAction,
		variables: map[string]interface{}{"input": input},
		timeout:   time.Second * 2,
		response:  &response,
	})
	return err
}

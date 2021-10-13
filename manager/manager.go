package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	// client related endpoints for i2i-manager
	clientRegister   = "/client/register"
	clientLogin      = "/client/login"
	clientRemove     = "/client/remove"
	clientNodeOrder  = "/client/node/order"
	clientNodeRemove = "/client/node/remove"
	clientNodeShow   = "/client/node/show"
	clientPlanList   = "/client/plan/list"
)

// Client runs API calls on i2i-manager
type Client struct {
	address  string
	password string
	jwt      string
	clientID string
}

type ClientOpt struct {
	// Address is either http or https url
	// https://manager.example.com
	// http://1.2.3.4
	Address  string
	ClientID string
	Password string
	JWT      string
}

func NewClient(opt ClientOpt) *Client {
	return &Client{
		address:  opt.Address,
		jwt:      opt.JWT,
		clientID: opt.ClientID,
		password: opt.Password,
	}
}

type apiCall struct {
	method       string
	httpPath     string
	payload      interface{}
	response     interface{}
	responseCode int
}

func (c *Client) apiCallDo(call *apiCall) error {
	var (
		request *http.Request
		client  = &http.Client{}
	)

	managerUrl, err := url.Parse(c.address)
	if err != nil {
		return err
	}
	managerUrl.Path = call.httpPath

	if call.payload != nil {
		data, err := json.Marshal(call.payload)
		if err != nil {
			return err
		}
		request, err = http.NewRequest(call.method, managerUrl.String(), bytes.NewBuffer(data))
	} else {
		request, err = http.NewRequest(call.method, managerUrl.String(), nil)
	}
	if err != nil {
		return err
	}

	request.Header.Set("Authorization", c.jwt)

	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	call.responseCode = resp.StatusCode
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("call failed with code %d", resp.StatusCode)
	}

	if call.response == nil {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, call.response)
}

func (c *Client) PlanList() ([]*CustomerPlan, error) {
	var plans []*CustomerPlan

	err := c.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: clientPlanList,
		payload:  nil,
		response: &plans,
	})

	return plans, err
}

type RegisterRequest struct {
	Password string `json:"password"`
}

type RegisterResponse struct {
	ClientID string `json:"client_id"`
}

// Register registers new client for the manager.
// As response clientID is returned
func (c *Client) Register(password string) (string, error) {
	var (
		request  = &RegisterRequest{Password: password}
		response = &RegisterResponse{}
	)

	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: clientRegister,
		payload:  request,
		response: response,
	})
	c.clientID = response.ClientID
	c.password = password

	return response.ClientID, err
}

type LoginRequest struct {
	ClientID string `json:"client_id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	JWT string `json:"jwt"`
}

func (c *Client) Login() error {
	var (
		request = &LoginRequest{
			ClientID: c.clientID,
			Password: c.password,
		}
		response = &LoginResponse{}
	)

	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: clientLogin,
		payload:  request,
		response: response,
	})
	c.jwt = response.JWT

	return err
}

type NodeOrderRequest struct {
	Token string `json:"token"`
	Plan  string `json:"plan"`
}

type NodeOrderReply struct {
	OrderID    string `json:"order_id"`
	CustomerID string `json:"customer_id"`
}

func (c *Client) NodeOrder(order *NodeOrderRequest) (*NodeOrderReply, error) {
	var (
		response = &NodeOrderReply{}
	)

	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: clientNodeOrder,
		payload:  order,
		response: response,
	})

	return response, err
}

func (c *Client) NodeUpdate() error {

	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: fmt.Sprintf("/client/update/%s", c.clientID),
		payload:  nil,
		response: nil,
	})

	return err
}

func (c *Client) WaitNodeProvisioned(timeout time.Duration) error {
	time.Sleep(timeout)
	return nil
}

type CustomerNode struct {
	ID string `json:"node_id,omitempty"`
	// ApiAddress is the ip:port address of the graphql endpoint
	ApiAddress string `json:"api_address,omitempty"`
	// Plan contains name of the plan assigned to user
	Plan string `json:"plan,omitempty"`
	// Token is keychain unlocking token
	Token      string `json:"token,omitempty"`
	ValidUntil int64  `json:"valid_until,omitempty"`
	CreatedAt  int64  `json:"created_at,omitempty"`
	Live       bool   `json:"live"`
}

func (c *Client) NodeShow() (*CustomerNode, error) {
	node := &CustomerNode{}

	err := c.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: clientNodeShow,
		payload:  nil,
		response: node,
	})

	return node, err
}

func (c *Client) ClientRemove() error {
	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: clientRemove,
		payload:  nil,
		response: nil,
	})

	return err
}

func (c *Client) NodeRemove() error {
	err := c.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: clientNodeRemove,
		payload:  nil,
		response: nil,
	})

	return err
}

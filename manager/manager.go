package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	Address string
}

func NewClient(opt ClientOpt) *Client {
	return &Client{
		address: opt.Address,
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

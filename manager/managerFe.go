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
	feLogin        = "/fe/login"
	fePlanAdd      = "/fe/plan/add"
	fePlanRemove   = "/fe/plan/remove"
	fePlanList     = "/fe/plan/list"
	feConfigUpdate = "/fe/config/update"

	feConfigShow  = "/fe/config/show"
	feClientList  = "/fe/client/list"
	feVersionSet  = "/fe/version"
	clientVersion = "/client/version"
)

func feClientUpdate(id string) string {
	return fmt.Sprintf("/fe/client/update/%s", id)
}

// FeClient runs frontend API calls on i2i-manager
type FeClient struct {
	address   string
	login     string
	password  string
	jwt       string
	stripeKey string
}

type FeOpt struct {
	// Address is either http or https url
	// https://manager.example.com
	// http://1.2.3.4
	Address string
	// admin login
	Login string
	// admin password
	Password string
	// stripe secret key
	StripeKey string
	Jwt       string
}

func NewFeClient(opt FeOpt) *FeClient {
	return &FeClient{
		address:   opt.Address,
		login:     opt.Login,
		password:  opt.Password,
		stripeKey: opt.StripeKey,
		jwt:       opt.Jwt,
	}
}

type FeLoginRequest struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type FeLoginResponse struct {
	JWT string `json:"jwt"`
}

func (m *FeClient) Login() error {
	var (
		request = &FeLoginRequest{
			User:     m.login,
			Password: m.password,
		}

		response = &FeLoginResponse{}
	)

	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: feLogin,
		payload:  request,
		response: response,
	})

	m.jwt = response.JWT

	return err
}

func (m *FeClient) Jwt() string {
	return m.jwt
}

type FePlan struct {
	// unique plan ID
	ID string `json:"id"`
	// The plan’s name, meant to be displayable to the customer.
	Name string `json:"name"`
	// The plan’s description, meant to be displayable to the customer.
	Description string `json:"description"`
	Currency    string `json:"currency"`
	Price       int64  `json:"price"`
	// Duration is the duration of the plan in seconds, after which it needs to be paid
	// or will be deleted
	Duration int64 `json:"duration"`
	// StorageSize is size for the storage in bytes of the docker for i2i
	StorageSize int64 `json:"storage_size"`
}

func (m *FeClient) PlanAdd(plan *FePlan) error {
	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: fePlanAdd,
		payload:  plan,
		response: nil,
	})

	return err
}

type FePlanRemoveRequest struct {
	// unique plan ID
	ID string `json:"id"`
}

func (m *FeClient) PlanRemove(request *FePlanRemoveRequest) error {
	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: fePlanRemove,
		payload:  request,
		response: nil,
	})

	return err
}

type FePlans struct {
	Plans []*FePlan `json:"plans"`
}

func (m *FeClient) PlanList() (*FePlans, error) {
	response := &FePlans{}

	err := m.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: fePlanList,
		payload:  nil,
		response: response,
	})

	return response, err
}

type FeManagerClient struct {
	ID               string      `json:"ID"`
	StripeCostumerID string      `json:"stripe_costumer_id"`
	JoinedAt         string      `json:"joined_at"`
	Instance         *FeInstance `json:"instance"`
}

type FeInstance struct {
	HostAPIPort     int    `json:"host_api_port"`
	HostNetworkPort int    `json:"host_network_port"`
	Plan            string `json:"plan"`
	ValidUntil      string `json:"valid_until"`
	CreatedAt       string `json:"created_at"`
	Live            bool   `json:"live"`
}

func (m *FeClient) ClientList() ([]*FeManagerClient, error) {
	var response []*FeManagerClient

	err := m.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: feClientList,
		payload:  nil,
		response: &response,
	})

	return response, err
}

func (m *FeClient) ClientUpdate(id string) error {
	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: feClientUpdate(id),
		payload:  nil,
		response: nil,
	})

	return err
}

type FeVersion struct {
	Version string `json:"version"`
}

func (m *FeClient) VersionSet(version string) error {
	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: feVersionSet,
		payload:  &FeVersion{Version: version},
		response: nil,
	})

	return err
}

func (m *FeClient) VersionShow() (string, error) {
	version := &FeVersion{}

	err := m.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: clientVersion,
		payload:  nil,
		response: version,
	})

	return version.Version, err
}

// ConfigPayments contains payments data
type FeConfigPayments struct {
	StripeSecretKey string `json:"stripe_secret_key,omitempty"`
	StripePublicKey string `json:"stripe_public_key,omitempty"`
}

type FeConfigAdmin struct {
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
}

// Config stores configuration options for the manager
type FeConfig struct {
	// Address is the IP address of the manager
	Address string `json:"address,omitempty"`
	// DockerImage is the name of the docker image run by manager for clients
	DockerImage string `json:"docker_image,omitempty"`
	// Company represents owner company name
	Company   string           `json:"company,omitempty"`
	Payments  FeConfigPayments `json:"payments,omitempty"`
	AdminUser FeConfigAdmin    `json:"admin,omitempty"`
}

func (m *FeClient) ConfigGet() (*FeConfig, error) {
	response := &FeConfig{}

	err := m.apiCallDo(&apiCall{
		method:   http.MethodGet,
		httpPath: feConfigShow,
		payload:  nil,
		response: response,
	})

	return response, err
}

func (m *FeClient) ConfigUpdate(cfg *FeConfig) error {
	err := m.apiCallDo(&apiCall{
		method:   http.MethodPost,
		httpPath: feConfigUpdate,
		payload:  cfg,
		response: nil,
	})

	return err
}

func (m *FeClient) apiCallDo(call *apiCall) error {
	var (
		request *http.Request
		client  = &http.Client{}
	)

	managerUrl, err := url.Parse(m.address)
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

	request.Header.Set("Authorization", m.jwt)

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

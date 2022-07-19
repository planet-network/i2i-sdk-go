package mobile

import (
	"github.com/planet-network/i2i-sdk-go/pc"
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
)

type PCClient struct {
	client *pc.RestClient
}

func New() *PCClient {
	client, _ := pc.NewRestClient("https://pc.vlow.me")

	return &PCClient{
		client: client,
	}
}

func (c *PCClient) SetAuthorization(auth string) {
	c.client.SetAuthorization(auth)
}

func (c *PCClient) Register(login string, secret string, method string) error {
	_, err := c.client.Register(login, secret, method)
	if err != nil {
		return err
	}

	return err
}

type LoginResponse struct {
	Authorization string   `json:"authorization"`
	SecureRandom  [32]byte `json:"secure_random"`
}

func (c *PCClient) Login(login string, secret string) (*LoginResponse, error) {
	response, err := c.client.Login(login, secret)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Authorization: response.Authorization,
		SecureRandom:  response.SecureRandom,
	}, err
}

func (c *PCClient) Initialize(secret string) error {
	response, err := c.client.SecureRandom(secret)
	if err != nil {
		return err
	}

	var (
		derivedPassword = cryptography.DerivedPassword([]byte(secret))
		preMasterKey    = cryptography.CalculatePreMasterKey(derivedPassword)
		masterKey       = cryptography.CalculateMasterKey(preMasterKey, response.SecureRandom)
	)

	c.client.SetMasterKey(masterKey)

	return err
}

func (c *PCClient) DataAdd(table string, key string, value string) error {
	err := c.client.DataAdd(table, key, value)
	return err
}

type DataResponse struct {
	// Table is name of the table in which the data is stored
	Table []byte `json:"table"`
	// Key is data unique key per table data is stored in
	Key []byte `json:"key"`
	// Value is data value
	Value []byte `json:"value"`
	// CreatedAt is epoch time when data was added
	CreatedAt int64 `json:"created_at"`
	// UpdatedAt is epoch time when data was last time updated
	ModifiedAt int64 `json:"modified_at"`
}

func (c *PCClient) DataGet(table string, key string) (*DataResponse, error) {
	response, err := c.client.DataGet(table, key)

	if err != nil {
		return nil, err
	}

	parsed, err := c.client.ParseDataResponse(response)
	if err != nil {
		return nil, err
	}

	return &DataResponse{
		Table:      parsed.Table,
		Key:        parsed.Key,
		Value:      parsed.Value,
		CreatedAt:  parsed.CreatedAt,
		ModifiedAt: parsed.ModifiedAt,
	}, nil
}

func (c *PCClient) ManagerNodeOrder(project string, password string) error {
	managerClient := pc.NewManager(c.client)

	return managerClient.NodeOrder(project, password)
}

func (c *PCClient) ManagerNodeDelete(project string) error {
	managerClient := pc.NewManager(c.client)

	return managerClient.NodeDelete(project)
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

func (c *PCClient) ManagerNodeShow(project string) (*CustomerNode, error) {
	managerClient := pc.NewManager(c.client)

	node, err := managerClient.NodeGet(project)
	if err != nil {
		return nil, err
	}

	return &CustomerNode{
		ID:         node.ID,
		ApiAddress: node.ApiAddress,
		Plan:       node.Plan,
		Token:      node.Token,
		ValidUntil: node.ValidUntil,
		CreatedAt:  node.CreatedAt,
		Live:       node.Live,
	}, nil
}

// VerifyAuthorization verifies JWT token
func VerifyAuthorization(auth string) error {
	return pc.VerifyAuthorization(auth)
}

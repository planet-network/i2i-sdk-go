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

// VerifyAuthorization verifies JWT token
func VerifyAuthorization(auth string) error {
	return pc.VerifyAuthorization(auth)
}

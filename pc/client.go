package pc

import (
	"fmt"
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/domain"
)

var errInvalidKeyLength = fmt.Errorf("invalid key length")

type Client struct {
	// master key which is combined pin+secure_random after successful login
	masterKey domain.MasterKey
	// privateShareKey is key used for decrypting data share requests
	privateShareKey [32]byte
	publicShareKey  [32]byte
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetMasterKey(key domain.MasterKey) {
	c.masterKey = key
}

func (c *Client) EncryptDataKeyOrTable(data []byte) ([]byte, error) {
	secretBox := cryptography.NewSecretBox(c.masterKey)

	return secretBox.EncryptUnsafe(data)
}

func (c *Client) DecryptDataKeyOrTable(data []byte) ([]byte, error) {
	secretBox := cryptography.NewSecretBox(c.masterKey)

	return secretBox.Decrypt(data)
}

func (c *Client) EncryptValueKey(data []byte) ([]byte, error) {
	secretBox := cryptography.NewSecretBox(c.masterKey)

	return secretBox.Encrypt(data)
}

func (c *Client) DecryptValueKey(data []byte) ([32]byte, error) {
	var (
		secretBox = cryptography.NewSecretBox(c.masterKey)
		valueKey  [domain.ValueKeyLength]byte
	)

	key, err := secretBox.Decrypt(data)
	if err != nil {
		return valueKey, err
	}

	if len(key) != domain.ValueKeyLength {
		return valueKey, errInvalidKeyLength
	}

	copy(valueKey[:], key)

	return valueKey, nil
}

func (c *Client) EncryptPrivateExchangeKey(key [32]byte) ([]byte, error) {
	secretBox := cryptography.NewSecretBox(c.masterKey)

	return secretBox.Encrypt(key[:])
}

func (c *Client) CalculateMasterKey(derivedPassword, secureRandom [32]byte) error {
	var (
		preMasterKey = cryptography.CalculatePreMasterKey(derivedPassword)
		masterKey    = cryptography.CalculateMasterKey(preMasterKey, secureRandom)
	)

	c.masterKey = masterKey
	return nil
}

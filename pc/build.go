package client

import (
	"crypto/rand"

	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/models"
	"golang.org/x/crypto/nacl/box"
)

func (r *RestClient) BuildRegisterRequest(login, secret, method string) (*models.RegisterRequest, error) {
	sKey := cryptography.DerivedPassword([]byte(secret))

	secureRandom, err := cryptography.GenerateSecureRandom()
	if err != nil {
		return nil, err
	}

	if err := r.client.CalculateMasterKey(sKey, secureRandom); err != nil {
		return nil, err
	}

	publicExchangeKey, privateExchangeKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	encryptedPrivateExchangeKey, err := r.client.EncryptPrivateExchangeKey(*privateExchangeKey)
	if err != nil {
		return nil, err
	}

	authorizationKey := cryptography.CalculateAuthKey(sKey)

	return &models.RegisterRequest{
		Login:                    login,
		AuthorizationKey:         authorizationKey,
		VerificationMethod:       method,
		SecureRandom:             secureRandom,
		ExchangePublicKey:        *publicExchangeKey,
		EncryptedSharePrivateKey: encryptedPrivateExchangeKey,
	}, nil
}

func (r *RestClient) BuildLoginRequest(login string, secret string) *models.LoginRequest {
	var (
		sKey    = cryptography.DerivedPassword([]byte(secret))
		authKey = cryptography.CalculateAuthKey(sKey)
	)

	return &models.LoginRequest{
		Login:            login,
		AuthorizationKey: authKey,
	}
}

func (r *RestClient) BuildDataAddRequest(table, key, value string) (*models.DataAddRequest, error) {
	valueKey, err := cryptography.GenerateValueKey()
	if err != nil {
		return nil, err
	}

	encryptedValue, err := cryptography.NewSecretBox(valueKey).Encrypt([]byte(value))
	if err != nil {
		return nil, err
	}

	encryptedValueKey, err := r.client.EncryptValueKey(valueKey[:])
	if err != nil {
		return nil, err
	}

	encryptedTable, err := r.client.EncryptDataKeyOrTable([]byte(table))
	if err != nil {
		return nil, err
	}

	encryptedKey, err := r.client.EncryptDataKeyOrTable([]byte(key))
	if err != nil {
		return nil, err
	}

	return &models.DataAddRequest{
		Table:             encryptedTable,
		Key:               encryptedKey,
		Value:             encryptedValue,
		EncryptedValueKey: encryptedValueKey,
	}, nil
}

func (r *RestClient) BuildDataUpdateRequest(table, key, value string, valueKey [32]byte) (*models.DataUpdateRequest, error) {
	encryptedValue, err := cryptography.NewSecretBox(valueKey).Encrypt([]byte(value))
	if err != nil {
		return nil, err
	}

	encryptedTable, err := r.client.EncryptDataKeyOrTable([]byte(table))
	if err != nil {
		return nil, err
	}

	encryptedKey, err := r.client.EncryptDataKeyOrTable([]byte(key))
	if err != nil {
		return nil, err
	}

	return &models.DataUpdateRequest{
		Table: encryptedTable,
		Key:   encryptedKey,
		Value: encryptedValue,
	}, nil
}

func (r *RestClient) BuildDataDeleteRequest(table, key string) (*models.DataDeleteRequest, error) {
	encryptedTable, err := r.client.EncryptDataKeyOrTable([]byte(table))
	if err != nil {
		return nil, err
	}

	encryptedKey, err := r.client.EncryptDataKeyOrTable([]byte(key))
	if err != nil {
		return nil, err
	}

	return &models.DataDeleteRequest{
		Table: encryptedTable,
		Key:   encryptedKey,
	}, nil
}

func (r *RestClient) BuildDataGetRequest(table, key string) (*models.DataGetRequest, error) {
	encryptedTable, err := r.client.EncryptDataKeyOrTable([]byte(table))
	if err != nil {
		return nil, err
	}

	encryptedKey, err := r.client.EncryptDataKeyOrTable([]byte(key))
	if err != nil {
		return nil, err
	}

	return &models.DataGetRequest{
		Table: encryptedTable,
		Key:   encryptedKey,
	}, nil
}

func (r *RestClient) BuildDataListRequest(tables []string) (*models.DataListRequest, error) {
	list := make([][]byte, 0, len(tables))

	for i := range tables {
		encryptedTable, err := r.client.EncryptDataKeyOrTable([]byte(tables[i]))
		if err != nil {
			return nil, err
		}
		list = append(list, encryptedTable)
	}

	return &models.DataListRequest{Tables: list}, nil
}

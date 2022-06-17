package client

import (
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/models"
)

func (r *RestClient) ParseDataResponse(resp *models.DataResponse) (*models.DataResponse, error) {
	return r.parseDataResponse(resp)
}

func (r *RestClient) parseValueKeyFromDataResponse(resp *models.DataResponse) ([32]byte, error) {
	return r.client.DecryptValueKey(resp.EncryptedValueKey)
}

func (r *RestClient) parseDataResponse(resp *models.DataResponse) (*models.DataResponse, error) {
	decryptedTable, err := r.client.DecryptDataKeyOrTable(resp.Table)
	if err != nil {
		return nil, err
	}

	decryptedKey, err := r.client.DecryptDataKeyOrTable(resp.Key)
	if err != nil {
		return nil, err
	}

	decryptedValueKey, err := r.client.DecryptValueKey(resp.EncryptedValueKey)
	if err != nil {
		return nil, err
	}

	value, err := cryptography.NewSecretBox(decryptedValueKey).Decrypt(resp.Value)
	if err != nil {
		return nil, err
	}

	return &models.DataResponse{
		Table:             decryptedTable,
		Key:               decryptedKey,
		Value:             value,
		CreatedAt:         resp.CreatedAt,
		ModifiedAt:        resp.ModifiedAt,
		EncryptedValueKey: resp.EncryptedValueKey,
	}, nil
}

func (r *RestClient) ParseDataListResponse(list []*models.DataResponse) ([]*models.DataResponse, error) {
	parsedList := make([]*models.DataResponse, 0, len(list))

	for i := range list {
		data, err := r.parseDataResponse(list[i])
		if err != nil {
			return nil, err
		}
		parsedList = append(parsedList, data)
	}

	return parsedList, nil
}

func (r *RestClient) ParseTableListResponse(response *models.TableListResponse) (*models.TableListResponse, error) {
	parsedResponse := &models.TableListResponse{Tables: make([][]byte, 0, len(response.Tables))}

	for i := range response.Tables {
		decryptedTable, err := r.client.DecryptDataKeyOrTable(response.Tables[i])
		if err != nil {
			return nil, err
		}

		parsedResponse.Tables = append(parsedResponse.Tables, decryptedTable)
	}

	return parsedResponse, nil
}

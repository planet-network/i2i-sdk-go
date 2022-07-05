package pc

import (
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/models"
	"net/http"
)

func (r *RestClient) UserInfo() (*models.UserInfoResponse, error) {
	var (
		response = &models.UserInfoResponse{}
	)

	err := r.do(call{
		path:     models.PathUserInfo,
		method:   http.MethodGet,
		response: response,
	})

	if err != nil {
		return nil, err
	}

	return response, err
}

func (r *RestClient) UserExchangeKey(login string) (*models.GetPublicExchangeKeyResponse, error) {
	var (
		userID   = cryptography.LoginHash(login)
		request  = &models.GetPublicExchangeKeyRequest{User: userID}
		response = &models.GetPublicExchangeKeyResponse{}
	)

	err := r.do(call{
		path:     models.PathUserExchangeKey,
		method:   http.MethodGet,
		response: response,
		request:  request,
	})

	if err != nil {
		return nil, err
	}

	return response, err
}

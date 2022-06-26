package pc

import (
	"github.com/planet-network/i2i-sdk-go/pc/cryptography"
	"github.com/planet-network/i2i-sdk-go/pc/models"
	"net/http"
)

func (r *RestClient) Register(login, secret, method string) (*models.RegisterRequest, error) {
	request, err := r.BuildRegisterRequest(login, secret, method)
	if err != nil {
		return nil, err
	}

	err = r.do(call{
		path:    models.PathRegister,
		method:  http.MethodPost,
		request: request,
	})

	return request, err
}

func (r *RestClient) Login(login, secret string) (*models.LoginResponse, error) {
	var (
		response = &models.LoginResponse{}
	)

	err := r.do(call{
		path:     models.PathLogin,
		method:   http.MethodPost,
		request:  r.BuildLoginRequest(login, secret),
		response: response,
	})

	if err != nil {
		return nil, err
	}

	r.authorization = response.Authorization

	return response, err
}

func (r *RestClient) SecureRandom(secret string) (*models.SecureRandomResponse, error) {
	var (
		sKey     = cryptography.DerivedPassword([]byte(secret))
		authKey  = cryptography.CalculateAuthKey(sKey)
		response = &models.SecureRandomResponse{}
	)

	err := r.do(call{
		path:   models.PathLogin,
		method: http.MethodPost,
		request: &models.SecureRandomRequest{
			AuthorizationKey: authKey,
		},
		response: response,
	})

	if err != nil {
		return nil, err
	}

	return response, err
}

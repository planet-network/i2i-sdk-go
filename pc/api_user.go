package pc

import (
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

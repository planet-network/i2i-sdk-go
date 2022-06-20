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

func (r *RestClient) DataGet(table, key string) (*models.DataResponse, error) {
	return r.dataGet(table, key)
}

func (r *RestClient) dataGet(table, key string) (*models.DataResponse, error) {
	var dataGetResponse models.DataResponse

	request, err := r.BuildDataGetRequest(table, key)
	if err != nil {
		return nil, err
	}

	err = r.do(call{
		path:     models.PathDataGet,
		method:   http.MethodGet,
		request:  request,
		response: &dataGetResponse,
	})

	return &dataGetResponse, err
}

func (r *RestClient) DataAdd(table, key, value string) error {
	dataAddRequest, err := r.BuildDataAddRequest(table, key, value)
	if err != nil {
		return err
	}

	return r.do(call{
		path:    models.PathDataAdd,
		method:  http.MethodPost,
		request: dataAddRequest,
	})
}

func (r *RestClient) DataUpdate(table, key, value string) error {
	data, err := r.dataGet(table, key)
	if err != nil {
		return err
	}

	valueKey, err := r.parseValueKeyFromDataResponse(data)
	if err != nil {
		return err
	}

	request, err := r.BuildDataUpdateRequest(table, key, value, valueKey)
	if err != nil {
		return err
	}

	return r.do(call{
		path:    models.PathDataUpdate,
		method:  http.MethodPost,
		request: request,
	})
}

func (r *RestClient) DataDelete(table, key string) error {
	dataDeleteRequest, err := r.BuildDataDeleteRequest(table, key)
	if err != nil {
		return err
	}

	return r.do(call{
		path:    models.PathDataDelete,
		method:  http.MethodPost,
		request: dataDeleteRequest,
	})
}

func (r *RestClient) DataList(tables []string) ([]*models.DataResponse, error) {
	var response []*models.DataResponse

	dataListRequest, err := r.BuildDataListRequest(tables)
	if err != nil {
		return nil, err
	}

	err = r.do(call{
		path:     models.PathDataList,
		method:   http.MethodGet,
		request:  dataListRequest,
		response: &response,
	})

	return response, err
}

func (r *RestClient) TableList() (*models.TableListResponse, error) {
	var response models.TableListResponse

	err := r.do(call{
		path:     models.PathTableList,
		method:   http.MethodGet,
		response: &response,
	})

	return &response, err
}

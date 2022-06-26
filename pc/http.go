package pc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/planet-network/i2i-sdk-go/pc/domain"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type RestClient struct {
	client        *Client
	httpClient    http.Client
	serverURL     *url.URL
	authorization string
}

type call struct {
	path     string
	method   string
	request  interface{}
	response interface{}
}

func NewRestClient(addr string) (*RestClient, error) {
	serverURL, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}

	return &RestClient{
		client: NewClient(),
		httpClient: http.Client{
			Timeout: time.Second * 30,
		},
		serverURL: serverURL,
	}, nil
}

func (r *RestClient) SetAuthorization(authorization string) {
	r.authorization = authorization
}

func (r *RestClient) SetMasterKey(key domain.MasterKey) {
	r.client.SetMasterKey(key)
}

func (r *RestClient) do(req call) error {
	var (
		err     error
		request *http.Request
		callUrl = &url.URL{
			Scheme: r.serverURL.Scheme,
			Host:   r.serverURL.Host,
			Path:   req.path,
		}
	)

	if req.request != nil {
		data, err := json.Marshal(req.request)
		if err != nil {
			return err
		}
		request, err = http.NewRequest(req.method, callUrl.String(), bytes.NewBuffer(data))
	} else {
		request, err = http.NewRequest(req.method, callUrl.String(), nil)
	}
	if err != nil {
		return err
	}

	if err := verifyAuthorization(r.authorization); err != nil {
		return err
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.authorization))

	resp, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("call failed with code %d", resp.StatusCode)
	}

	if req.response == nil {
		return nil
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, req.response)
}

func verifyAuthorization(auth string) error {
	jwt.lol()
	return nil
}

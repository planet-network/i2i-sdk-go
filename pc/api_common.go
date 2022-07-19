package pc

import (
	"github.com/planet-network/i2i-sdk-go/pc/models"
	"net/http"
)

func (r *RestClient) Ping() error {
	err := r.do(call{
		path:   models.PathPing,
		method: http.MethodGet,
	})

	return err
}

func (r *RestClient) Capabilities() (*models.CapabilitiesResponse, error) {
	var capabilities models.CapabilitiesResponse

	err := r.do(call{
		path:     models.PathCapabilities,
		method:   http.MethodGet,
		response: &capabilities,
	})

	return &capabilities, err
}

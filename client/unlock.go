package client

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

// Unlock provides needed cryptographic keys to i2i to unlock it.
func (c *Client) Unlock() error {
	var (
		httpClient = http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		}
	)

	keyChain := &KeychainListenerProvider{
		NetworkPublicKey:   c.keychain.NetworkPublicKey.String(),
		NetworkPrivateKey:  c.keychain.NetworkPrivateKey.String(),
		StoragePublicKey:   c.keychain.StoragePublicKey.String(),
		StoragePrivateKey:  c.keychain.StoragePrivateKey.String(),
		SignaturePublicKey: c.keychain.SignaturePublicKey.String(),
	}

	data, err := json.Marshal(keyChain)
	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, c.nodeAddress(), bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	request.Header.Set("KEYCHAIN_TOKEN", c.token)

	resp, err := httpClient.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed with code %d", resp.StatusCode)
	}

	return nil
}

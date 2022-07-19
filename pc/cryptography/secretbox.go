package cryptography

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"
)

// SecretBox is structure performing cryptographic operations
// using symmetric cryptographic algorithm defined in:
// https://nacl.cr.yp.to
type SecretBox struct {
	key [32]byte
}

// NewSecretBox creates new instance of SecretBox
func NewSecretBox(key [32]byte) *SecretBox {
	return &SecretBox{key: key}
}

// Encrypt encrypts message using symmetric algorithm
func (s *SecretBox) Encrypt(message []byte) ([]byte, error) {
	var nonce [24]byte
	if _, err := rand.Read(nonce[:]); err != nil {
		return nil, err
	}

	return secretbox.Seal(nonce[:], message, &nonce, &s.key), nil
}

// EncryptUnsafe encrypts message using symmetric algorithm without nonce
func (s *SecretBox) EncryptUnsafe(message []byte) ([]byte, error) {
	var nonce [24]byte
	return secretbox.Seal(nonce[:], message, &nonce, &s.key), nil
}

// Decrypt decrypts message using symmetric algorithm
func (s *SecretBox) Decrypt(encrypted []byte) ([]byte, error) {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &s.key)
	if !ok {
		return nil, fmt.Errorf("failed to decrypt message")
	}

	return decrypted, nil
}

// DecryptToObject decrypts message using symmetric algorithm and deserializes to o
func (s *SecretBox) DecryptToObject(encrypted []byte, o interface{}) error {
	var decryptNonce [24]byte
	copy(decryptNonce[:], encrypted[:24])

	decrypted, ok := secretbox.Open(nil, encrypted[24:], &decryptNonce, &s.key)
	if !ok {
		return fmt.Errorf("failed to decrypt message")
	}

	return json.Unmarshal(decrypted, o)
}

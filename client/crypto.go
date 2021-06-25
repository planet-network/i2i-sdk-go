package client

import (
	"crypto/rand"

	"golang.org/x/crypto/nacl/box"

	"golang.org/x/crypto/nacl/sign"
)

// Signer signs and verifies small messages using public-key cryptography.
type Signer struct {
	privateKey Key64
	publicKey  Key32
}

// GenerateSigner generates Signer instance with set of cryptographic keys used for signing
func GenerateSigner() (*Signer, error) {
	publicKey, privateKey, err := sign.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	signer := &Signer{}
	copy(signer.publicKey[:], publicKey[:])
	copy(signer.privateKey[:], privateKey[:])
	return signer, nil
}

// Box is structure performing cryptographic operations
// using asymmetric cryptographic algorithm defined in:
// https://nacl.cr.yp.to
type Box struct {
	privateKey Key32
	publicKey  Key32
}

// NewOneShotBox creates Box instance with randomly generated keys
// It's supposed to be one time use.
func NewOneShotBox() (*Box, error) {
	publicKey, privateKey, err := box.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}
	return &Box{
		publicKey:  *publicKey,
		privateKey: *privateKey,
	}, nil
}

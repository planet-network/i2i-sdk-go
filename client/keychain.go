package client

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// FullKeychain contains set of cryptographic keys for i2i
type FullKeychain struct {
	// Network keys are used for encrypting network traffic
	// and presenting identity (public key)
	NetworkPublicKey  Key32
	NetworkPrivateKey Key32
	// Storage keys are used for encrypting storage
	StoragePublicKey  Key32
	StoragePrivateKey Key32
	// Signature keys are used for signing transactions
	SignaturePublicKey  Key32
	SignaturePrivateKey Key64
}

type KeychainListenerProvider struct {
	NetworkPublicKey   string `json:"network_public_key"`
	NetworkPrivateKey  string `json:"network_private_key"`
	StoragePublicKey   string `json:"storage_public_key"`
	StoragePrivateKey  string `json:"storage_private_key"`
	SignaturePublicKey string `json:"signature_public_key"`
}

// GenerateFullKeychain creates keychain with new keys
func GenerateFullKeychain() (*FullKeychain, error) {
	signer, err := GenerateSigner()
	if err != nil {
		return nil, err
	}
	mainBox, err := NewOneShotBox()
	if err != nil {
		return nil, err
	}
	storageBox, err := NewOneShotBox()
	if err != nil {
		return nil, err
	}
	return &FullKeychain{
		NetworkPublicKey: mainBox.publicKey, NetworkPrivateKey: mainBox.privateKey,
		StoragePublicKey: storageBox.publicKey, StoragePrivateKey: storageBox.privateKey,
		SignaturePublicKey: signer.publicKey, SignaturePrivateKey: signer.privateKey,
	}, nil
}

// SaveToFileSafe saves keychain to file.
// The file must not exist, otherwise error is returned.
func (k *FullKeychain) SaveToFileSafe(path string) error {
	if _, err := os.Stat(path); err == nil {
		return ErrAlreadyExist
	}

	raw, err := json.Marshal(k)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(path, raw, 0600)
}

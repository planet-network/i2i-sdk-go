package domain

import "encoding/hex"

const (
	authorizationKeyLength = 32
	secureRandomLength     = 32
	ValueKeyLength         = 32
	masterKeyLength        = 32
)

// SecureRandom is server generated random which becomes part of user master key
type SecureRandom [secureRandomLength]byte

// AuthorizationKey is used for authorizing user in the enclave
type AuthorizationKey [authorizationKeyLength]byte

// ValueKey is used for encrypting value of the data
type ValueKey [ValueKeyLength]byte

// MasterKey is main encryption key used only by client. This key must never be revealed to server
type MasterKey [masterKeyLength]byte

func (m *MasterKey) String() string {
	return hex.EncodeToString(m[:])
}

func (m *MasterKey) IsEmpty() bool {
	for i := range m {
		if m[i] != 0 {
			return false
		}
	}
	return true
}

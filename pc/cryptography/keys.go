package cryptography

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"github.com/planet-network/i2i-sdk-go/pc/domain"
	"golang.org/x/crypto/blake2b"
)

func CalculateAuthKey(sKey [32]byte) domain.AuthorizationKey {
	var (
		suffix           = []byte("auth key")
		buffer           = bytes.Join([][]byte{suffix, sKey[:]}, []byte{})
		sum              = blake2b.Sum256(buffer)
		authorizationKey domain.AuthorizationKey
	)

	copy(authorizationKey[:], sum[:])

	return authorizationKey
}

func CalculatePreMasterKey(sKey [32]byte) [32]byte {
	var (
		suffix       = []byte("master key")
		buffer       = bytes.Join([][]byte{suffix, sKey[:]}, []byte{})
		sum          = blake2b.Sum256(buffer)
		preMasterKey [32]byte
	)

	copy(preMasterKey[:], sum[:])

	return preMasterKey
}

func CalculateMasterKey(preMasterKey, sRand [32]byte) domain.MasterKey {
	var (
		buffer    = bytes.Join([][]byte{preMasterKey[:], sRand[:]}, []byte{})
		sum       = blake2b.Sum256(buffer)
		masterKey domain.MasterKey
	)

	copy(masterKey[:], sum[:])

	return masterKey
}

func GenerateSecureRandom() (domain.SecureRandom, error) {
	var secureRandom domain.SecureRandom
	_, err := rand.Read(secureRandom[:])
	return secureRandom, err
}

func GenerateValueKey() ([32]byte, error) {
	var valueKey domain.SecureRandom
	_, err := rand.Read(valueKey[:])
	return valueKey, err
}

func MasterKeyFromString(s string) (domain.MasterKey, error) {
	buffer, err := hex.DecodeString(s)
	if err != nil {
		return domain.MasterKey{}, err
	}

	if len(buffer) != 32 {
		return domain.MasterKey{}, errInvalidKeyLength
	}

	var masterKey domain.MasterKey
	copy(masterKey[:], buffer)
	return masterKey, nil
}

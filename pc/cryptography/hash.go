package cryptography

import "golang.org/x/crypto/blake2b"

// LoginHash creates user id from user login
func LoginHash(word string) [32]byte {
	var userID [32]byte

	hash := blake2b.Sum256([]byte(word))
	copy(userID[:], hash[:])

	return userID
}

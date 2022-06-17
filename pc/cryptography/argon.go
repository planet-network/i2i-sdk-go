package cryptography

import (
	"golang.org/x/crypto/argon2"
)

func DerivedPassword(password []byte) (ret [32]byte) {
	var salt [16]byte
	key := argon2.IDKey(password, salt[:], 1, 64*1024, 4, 32)
	copy(ret[:], key)
	return ret
}

package cryptography

import (
	"crypto/rand"
	"fmt"
)

// RandomString32 generates random 32 character string with 128 bit randomness
func RandomString32() (string, error) {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf[:]), nil
}

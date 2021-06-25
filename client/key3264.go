package client

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const (
	// KeySize defines size of cryptographic key in bytes
	KeySize = 32

	// SignKeySize defines size of cryptographic key used for signing in bytes
	SignKeySize = 64
)

// Key32 is 32 byte cryptographic key
type Key32 [KeySize]byte

// Key64 is 64 byte private key used for generating signatures
type Key64 [SignKeySize]byte

// Equal returns true if two keys are the same
func (k Key32) Equal(other Key32) bool {
	return bytes.Equal(k[:], other[:])
}

// Equal returns true if two keys are the same
func (k Key64) Equal(other Key64) bool {
	return bytes.Equal(k[:], other[:])
}

// String returns string representation of a key
func (k *Key32) String() string {
	return fmt.Sprintf("%x", k[:])
}

// String returns string representation of a key
func (k *Key64) String() string {
	return fmt.Sprintf("%x", k[:])
}

// Key32FromByte creates key from a slice of bytes
func Key32FromByte(raw []byte) (k Key32, err error) {
	if len(raw) != KeySize {
		return k, ErrKeyBadLength
	}

	copy(k[:], raw)
	return k, nil
}

// Key64FromByte creates key from a slice of bytes
func Key64FromByte(raw []byte) (k Key64, err error) {
	if len(raw) != SignKeySize {
		return k, ErrKeyBadLength
	}

	copy(k[:], raw)
	return k, nil
}

// Key32FromString creates key from a string
func Key32FromString(raw string) (k Key32, err error) {
	if len(raw) != 2*KeySize {
		return k, ErrKeyBadLength
	}

	decoded, err := hex.DecodeString(raw)
	if err != nil {
		return k, err
	}

	copy(k[:], decoded)
	return k, nil
}

// Key64FromString creates key from a string
func Key64FromString(raw string) (k Key64, err error) {
	if len(raw) != 4*KeySize {
		return k, ErrKeyBadLength
	}

	decoded, err := hex.DecodeString(raw)
	if err != nil {
		return k, err
	}

	copy(k[:], decoded)
	return k, nil
}

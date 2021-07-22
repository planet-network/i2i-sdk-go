package client

import (
	"errors"
	"fmt"
)

// ErrKeyBadLength is returned when length of the cryptographic key is invalid
var ErrKeyBadLength = errors.New("invalid key length")

// ErrAlreadyExist is returned when object already exist
var ErrAlreadyExist = errors.New("already exist")

// ErrHttpWithCode is returned when non 200 code is returned
func ErrHttpWithCode(code int) error {
	return fmt.Errorf("failed with http code: %d", code)
}

// ErrNotADir is returned when expected to have directory
func ErrNotADir() error {
	return errors.New("not a directory3")
}

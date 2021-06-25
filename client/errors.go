package client

import "errors"

// ErrKeyBadLength is returned when length of the cryptographic key is invalid
var ErrKeyBadLength = errors.New("invalid key length")

// ErrAlreadyExist is returned when object already exist
var ErrAlreadyExist = errors.New("already exist")

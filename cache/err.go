package cache

import (
	"errors"
)

var (
	// ErrEntryNotFound is an error type struct which is returned when entry was not found for provided key
	ErrEntryNotFound   = errors.New("Entry not found")
	ErrDataTypeInvalid = errors.New("Data type invalid")
)

var Nil = ErrEntryNotFound

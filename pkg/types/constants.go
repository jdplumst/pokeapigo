package types

import "errors"

var (
	ErrFetch     = errors.New("error fetching resource")
	ErrBody      = errors.New("error reading request body")
	ErrUnmarshal = errors.New("error unmarshalling request body")
)

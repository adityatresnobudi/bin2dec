package shared

import "errors"

var (
	ErrInvalidBinary = errors.New("input have number other than 0 or 1")
	ErrMoreThanEight = errors.New("input have more than 8 number")
)

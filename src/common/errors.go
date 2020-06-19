package common

import "errors"

var (
	// Errors - platform errors
	Errors = map[string]error{
		"ErrNotFound": errors.New("not found"),
	}
)

type TestLancaster struct {
	Foo string
}

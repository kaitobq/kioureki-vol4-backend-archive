package errors

import "errors"

var (
	ErrReadFile = errors.New("could not read SQL file")
	ErrExecFile = errors.New("could not execute SQL file")
)
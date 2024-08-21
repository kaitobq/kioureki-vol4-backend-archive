package errors

import "errors"

// repository

var (
	ErrPasswordEmpty = errors.New("password is empty")
)

// usecase

var (
	ErrEmailAlreadyInUse = errors.New("email is already in use")
	ErrInvalidPassword = errors.New("invalid password")
)

// controller
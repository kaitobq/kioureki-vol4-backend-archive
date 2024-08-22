package errors

import "errors"

// repository

// usecase

var (
	ErrUserAlreadyInOrganization = errors.New("user is already in organization")
	ErrInviteAlreadySent = errors.New("invite already sent to this user")
)

// controller
package error

import "errors"

var (
	ErrMissingAuthToken       = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat = errors.New("invalid auth token format")
	ErrGetUserProfile         = errors.New("error occured when getting user profile")
	ErrNoUserId               = errors.New("no user id")
	ErrInvalidUserId          = errors.New("invalid user id")
	ErrMissingRequestBody     = errors.New("Missing request body.")
	ErrInvalidRequestBody     = errors.New("Invalid request body.")
)

package error

import "errors"

var (
	ErrMissingAuthToken           = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat     = errors.New("invalid auth token format")
	ErrInvalidAuthToken           = errors.New("invalid auth token")
	ErrJWT                        = errors.New("jwt error")
	ErrNoUserId                   = errors.New("no user id")
	ErrInvalidUserId              = errors.New("invalid user id")
	ErrMissingRequestBody         = errors.New("Missing request body.")
	ErrInvalidRequestBody         = errors.New("Invalid request body.")
	ErrInvalidRegistrationRequest = errors.New("invalid registration info")
	ErrInvalidLoginRequest        = errors.New("invalid login info")
	ErrInvalidName                = errors.New("invalid name")
	ErrInvalidPassword            = errors.New("invalid password")
	ErrUserNotFound               = errors.New("user not found")
	ErrPhoneRegistered            = errors.New("phone is already registered")
	ErrEmailRegistered            = errors.New("email is already registered")
	ErrUserDeactivated            = errors.New("user is deactivated")
	ErrInvalidPhoneLogin          = errors.New("invalid phone or password")
	ErrInvalidEmailLogin          = errors.New("invalid email or password")

	// system error
	ErrGetUser    = errors.New("error occurred when getting user")
	ErrInsertUser = errors.New("error occurred when inserting user")
)

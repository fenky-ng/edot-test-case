package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:           http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat:     http.StatusUnauthorized,
	ErrInvalidAuthToken:           http.StatusUnauthorized,
	ErrNoUserId:                   http.StatusUnauthorized,
	ErrInvalidUserId:              http.StatusBadRequest,
	ErrMissingRequestBody:         http.StatusBadRequest,
	ErrInvalidRequestBody:         http.StatusBadRequest,
	ErrInvalidRegistrationRequest: http.StatusBadRequest,
	ErrInvalidLoginRequest:        http.StatusBadRequest,
	ErrInvalidName:                http.StatusBadRequest,
	ErrInvalidPassword:            http.StatusBadRequest,
	ErrUserNotFound:               http.StatusNotFound,
	ErrPhoneRegistered:            http.StatusConflict,
	ErrEmailRegistered:            http.StatusConflict,
	ErrUserDeactivated:            http.StatusForbidden,
	ErrInvalidPhoneLogin:          http.StatusUnauthorized,
	ErrInvalidEmailLogin:          http.StatusUnauthorized,
}

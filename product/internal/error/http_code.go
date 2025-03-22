package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:       http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat: http.StatusUnauthorized,
	ErrGetUserProfile:         http.StatusInternalServerError,
	ErrNoUserId:               http.StatusUnauthorized,
	ErrInvalidUserId:          http.StatusBadRequest,
	ErrMissingRequestBody:     http.StatusBadRequest,
	ErrInvalidRequestBody:     http.StatusBadRequest,
}

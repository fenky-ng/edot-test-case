package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:       http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat: http.StatusUnauthorized,
	ErrGetUserProfile:         http.StatusInternalServerError,
	ErrNoJWT:                  http.StatusUnauthorized,
	ErrNoUserId:               http.StatusUnauthorized,
	ErrInvalidUserId:          http.StatusBadRequest,
	ErrMissingRequestBody:     http.StatusBadRequest,
	ErrInvalidRequestBody:     http.StatusBadRequest,
	ErrInvalidProductId:       http.StatusBadRequest,
	ErrInvalidName:            http.StatusBadRequest,
	ErrInvalidPrice:           http.StatusBadRequest,
	ErrInvalidStatus:          http.StatusBadRequest,
	ErrProductNotFound:        http.StatusNotFound,
}

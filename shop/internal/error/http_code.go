package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:       http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat: http.StatusUnauthorized,
	ErrGetUserProfile:         http.StatusInternalServerError,
	ErrNoUserId:               http.StatusUnauthorized,
	ErrInvalidUserId:          http.StatusBadRequest,
	ErrInvalidShopId:          http.StatusBadRequest,
	ErrMissingRequestBody:     http.StatusBadRequest,
	ErrInvalidRequestBody:     http.StatusBadRequest,
	ErrInvalidName:            http.StatusBadRequest,
	ErrShopNotFound:           http.StatusNotFound,
	ErrShopDeactivated:        http.StatusForbidden,
	ErrAlreadyOwnAShop:        http.StatusConflict,
}

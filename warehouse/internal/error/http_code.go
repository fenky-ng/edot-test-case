package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:          http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat:    http.StatusUnauthorized,
	ErrGetUserProfile:            http.StatusInternalServerError,
	ErrNoJWT:                     http.StatusUnauthorized,
	ErrNoUserId:                  http.StatusUnauthorized,
	ErrInvalidUserId:             http.StatusBadRequest,
	ErrMissingRequestBody:        http.StatusBadRequest,
	ErrInvalidRequestBody:        http.StatusBadRequest,
	ErrInvalidName:               http.StatusBadRequest,
	ErrInvalidStatus:             http.StatusBadRequest,
	ErrInvalidUpdateRequest:      http.StatusBadRequest,
	ErrInvalidWarehouseId:        http.StatusBadRequest,
	ErrInvalidProductId:          http.StatusBadRequest,
	ErrMinOneProductIdQueryParam: http.StatusBadRequest,
	ErrMaxWarehousePerShop:       http.StatusBadRequest,
	ErrWarehouseNotFound:         http.StatusNotFound,
	ErrNotWarehouseOwner:         http.StatusForbidden,
}

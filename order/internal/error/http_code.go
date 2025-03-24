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
	ErrNoOrderItem:            http.StatusBadRequest,
	ErrInvalidProductId:       http.StatusBadRequest,
	ErrInvalidWarehouseId:     http.StatusBadRequest,
	ErrInvalidOrderQuantity:   http.StatusBadRequest,
	ErrProductNotFound:        http.StatusBadRequest,
	ErrProductNotActive:       http.StatusBadRequest,
	ErrUserOwnProduct:         http.StatusForbidden,
	ErrShopNotActive:          http.StatusBadRequest,
	ErrWarehouseNotActive:     http.StatusBadRequest,
	ErrInsufficientStock:      http.StatusBadRequest,
	ErrMissingOrderNo:         http.StatusBadRequest,
	ErrMissingPaymentRefNo:    http.StatusBadRequest,
	ErrOrderNotFound:          http.StatusNotFound,
}

package error

import "net/http"

var errorHttpCodeMap = map[error]int{
	ErrMissingAuthToken:                http.StatusUnauthorized,
	ErrInvalidAuthTokenFormat:          http.StatusUnauthorized,
	ErrGetUserProfile:                  http.StatusInternalServerError,
	ErrNoJWT:                           http.StatusUnauthorized,
	ErrNoUserId:                        http.StatusUnauthorized,
	ErrInvalidUserId:                   http.StatusBadRequest,
	ErrMissingApiKey:                   http.StatusBadRequest,
	ErrInvalidApiKey:                   http.StatusBadRequest,
	ErrMissingRequestBody:              http.StatusBadRequest,
	ErrInvalidRequestBody:              http.StatusBadRequest,
	ErrInvalidName:                     http.StatusBadRequest,
	ErrInvalidStatus:                   http.StatusBadRequest,
	ErrInvalidUpdateRequest:            http.StatusBadRequest,
	ErrInvalidWarehouseId:              http.StatusBadRequest,
	ErrInvalidProductId:                http.StatusBadRequest,
	ErrMinOneProductIdQueryParam:       http.StatusBadRequest,
	ErrMaxWarehousePerShop:             http.StatusBadRequest,
	ErrWarehouseNotFound:               http.StatusNotFound,
	ErrNotWarehouseOwner:               http.StatusForbidden,
	ErrInvalidStock:                    http.StatusBadRequest,
	ErrInvalidStockTransfer:            http.StatusBadRequest,
	ErrNotProductOwner:                 http.StatusForbidden,
	ErrInvalidStockTransferDestination: http.StatusBadRequest,
	ErrStockLocked:                     http.StatusConflict,
	ErrInsufficientStock:               http.StatusBadRequest,
	ErrWarehouseInactive:               http.StatusBadRequest,
	ErrInvalidQuantityDeduction:        http.StatusBadRequest,
}

package error

import "errors"

var (
	ErrMissingAuthToken          = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat    = errors.New("invalid auth token format")
	ErrGetUserProfile            = errors.New("error occured when getting user profile")
	ErrNoJWT                     = errors.New("no jwt")
	ErrNoUserId                  = errors.New("no user id")
	ErrInvalidUserId             = errors.New("invalid user id")
	ErrMissingRequestBody        = errors.New("Missing request body.")
	ErrInvalidRequestBody        = errors.New("Invalid request body.")
	ErrInvalidName               = errors.New("invalid name")
	ErrInvalidStatus             = errors.New("invalid status")
	ErrInvalidUpdateRequest      = errors.New("invalid update request")
	ErrInvalidWarehouseId        = errors.New("invalid warehouse id")
	ErrInvalidProductId          = errors.New("invalid product id")
	ErrMinOneProductIdQueryParam = errors.New("minimum one product id in query param")
	ErrMaxWarehousePerShop       = errors.New("reached max warehouse allowed per shop")
	ErrWarehouseNotFound         = errors.New("warehouse not found")
	ErrNotWarehouseOwner         = errors.New("not warehouse owner")

	// system error
	ErrGetMyShop       = errors.New("error occured when getting my shop")
	ErrInsertWarehouse = errors.New("error occurred when inserting warehouse")
	ErrUpdateWarehouse = errors.New("error occurred when updating warehouse")
	ErrGetWarehouses   = errors.New("error occurred when getting warehouses")
	ErrGetStocks       = errors.New("error occurred when getting stocks")
)

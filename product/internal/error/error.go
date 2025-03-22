package error

import "errors"

var (
	ErrMissingAuthToken       = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat = errors.New("invalid auth token format")
	ErrGetUserProfile         = errors.New("error occured when getting user profile")
	ErrNoJWT                  = errors.New("no jwt")
	ErrNoUserId               = errors.New("no user id")
	ErrInvalidUserId          = errors.New("invalid user id")
	ErrMissingRequestBody     = errors.New("Missing request body.")
	ErrInvalidRequestBody     = errors.New("Invalid request body.")
	ErrInvalidProductId       = errors.New("invalid product id")
	ErrInvalidName            = errors.New("invalid name")
	ErrInvalidPrice           = errors.New("invalid price")
	ErrInvalidStatus          = errors.New("invalid status")
	ErrProductNotFound        = errors.New("product not found")

	// system error
	ErrGetMyShop     = errors.New("error occured when getting my shop")
	ErrGetShops      = errors.New("error occured when getting shops")
	ErrGetStocks     = errors.New("error occurred when getting stocks")
	ErrGetProducts   = errors.New("error occurred when getting products")
	ErrGetProduct    = errors.New("error occurred when getting product")
	ErrInsertProduct = errors.New("error occurred when inserting product")
)

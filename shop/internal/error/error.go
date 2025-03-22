package error

import "errors"

var (
	ErrMissingAuthToken       = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat = errors.New("invalid auth token format")
	ErrGetUserProfile         = errors.New("error occured when getting user profile")
	ErrNoUserId               = errors.New("no user id")
	ErrInvalidUserId          = errors.New("invalid user id")
	ErrInvalidShopId          = errors.New("invalid shop id")
	ErrMissingRequestBody     = errors.New("Missing request body.")
	ErrInvalidRequestBody     = errors.New("Invalid request body.")
	ErrInvalidName            = errors.New("invalid name")
	ErrShopNotFound           = errors.New("shop not found")
	ErrShopDeactivated        = errors.New("shop is deactivated")
	ErrAlreadyOwnAShop        = errors.New("user already own a shop")

	// system error
	ErrGetShops   = errors.New("error occurred when getting shops")
	ErrGetShop    = errors.New("error occurred when getting shop")
	ErrInsertShop = errors.New("error occurred when inserting shop")
)

package error

import "errors"

var (
	ErrMissingAuthToken       = errors.New("missing auth token")
	ErrInvalidAuthTokenFormat = errors.New("invalid auth token format")
	ErrGetUserProfile         = errors.New("error occurred when getting user profile")
	ErrNoJWT                  = errors.New("no jwt")
	ErrNoUserId               = errors.New("no user id")
	ErrInvalidUserId          = errors.New("invalid user id")
	ErrMissingRequestBody     = errors.New("Missing request body.")
	ErrInvalidRequestBody     = errors.New("Invalid request body.")
	ErrNoOrderItem            = errors.New("no order item")
	ErrInvalidProductId       = errors.New("invalid product id")
	ErrInvalidWarehouseId     = errors.New("invalid warehouse id")
	ErrInvalidOrderQuantity   = errors.New("invalid order quantity")
	ErrProductNotFound        = errors.New("product not found")
	ErrProductNotActive       = errors.New("product not active")
	ErrUserOwnProduct         = errors.New("user own product")
	ErrShopNotActive          = errors.New("shop not active")
	ErrWarehouseNotActive     = errors.New("warehouse not active")
	ErrInsufficientStock      = errors.New("warehouse not active")
	ErrMissingOrderNo         = errors.New("missing order number")
	ErrMissingPaymentRefNo    = errors.New("missing payment reference number")
	ErrOrderNotFound          = errors.New("error not found")

	// system
	ErrDatabaseTransaction = errors.New("database transaction error")
	ErrGetMyShop           = errors.New("error occurred when getting my shop")
	ErrGetProducts         = errors.New("error occurred when getting products")
	ErrInsertOrder         = errors.New("error occurred when inserting order")
	ErrInsertOrderDetails  = errors.New("error occurred when inserting order details")
	ErrUpdateOrder         = errors.New("error occurred when updating order")
	ErrGetOrders           = errors.New("error occurred when getting orders")
	ErrGetOrderDetails     = errors.New("error occurred when getting order details")
	ErrDeductStocks        = errors.New("error occurred when deducting stocks")
)

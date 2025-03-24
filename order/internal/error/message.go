package error

var errorMessage = map[error]string{
	ErrMissingAuthToken:       "Missing authentication token.",
	ErrInvalidAuthTokenFormat: "Invalid authentication token format.",
	ErrGetUserProfile:         "Failed to fetch user profile. Please try again later.",
	ErrNoJWT:                  "Unauthorized: missing token.",
	ErrNoUserId:               "Unauthorized: missing user ID in token.",
	ErrInvalidUserId:          "Invalid user ID format. Please provide a valid UUID.",
	ErrMissingRequestBody:     "Missing request body.",
	ErrInvalidRequestBody:     "Invalid request body.",
	ErrNoOrderItem:            "At least one order item must be provided in the request to create order.",
	ErrInvalidProductId:       "Invalid product ID. Please provide a valid UUID.",
	ErrInvalidWarehouseId:     "Invalid warehouse ID. Please provide a valid UUID.",
	ErrInvalidOrderQuantity:   "Order quantity must be greater than 0.",
	ErrProductNotFound:        "Product not found.",
	ErrProductNotActive:       "Product is not active.",
	ErrUserOwnProduct:         "You cannot purchase your own product.",
	ErrShopNotActive:          "Shop is not active.",
	ErrWarehouseNotActive:     "Warehouse is not active.",
	ErrInsufficientStock:      "Insufficient stock available for this request.",
	ErrMissingOrderNo:         "Missing order number.",
	ErrMissingPaymentRefNo:    "Missing payment reference number.",
	ErrOrderNotFound:          "Order not found.",
}

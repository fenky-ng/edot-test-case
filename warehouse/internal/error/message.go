package error

var errorMessage = map[error]string{
	ErrMissingAuthToken:          "Missing authentication token.",
	ErrInvalidAuthTokenFormat:    "Invalid authentication token format.",
	ErrGetUserProfile:            "Failed to fetch user profile. Please try again later.",
	ErrNoJWT:                     "Unauthorized: missing token.",
	ErrNoUserId:                  "Unauthorized: missing user ID in token.",
	ErrInvalidUserId:             "Invalid user ID format. Please provide a valid UUID.",
	ErrMissingRequestBody:        "Missing request body.",
	ErrInvalidRequestBody:        "Invalid request body.",
	ErrInvalidName:               "The name must be at least 3 characters long. Please enter a valid name.",
	ErrInvalidStatus:             "The status must be either ACTIVE or INACTIVE. Please enter a valid status.",
	ErrInvalidUpdateRequest:      "No update information provided. Please include at least one field to update.",
	ErrInvalidWarehouseId:        "Invalid warehouse ID. Please provide a valid UUID.",
	ErrInvalidProductId:          "Invalid product ID. Please provide a valid UUID.",
	ErrMinOneProductIdQueryParam: "At least one product_id must be provided in the query parameters.",
	ErrMaxWarehousePerShop:       "Shop cannot have more warehouses. Maximum allowed limit reached.",
	ErrWarehouseNotFound:         "Warehouse not found.",
	ErrNotWarehouseOwner:         "You do not have permission to update this warehouse.",
}

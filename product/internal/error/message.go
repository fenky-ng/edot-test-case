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
	ErrInvalidProductId:       "Invalid product ID. Please provide a valid UUID.",
	ErrInvalidName:            "The name must be at least 3 characters long. Please enter a valid name.",
	ErrInvalidPrice:           "The price must be greater than 0. Please enter a valid price.",
	ErrInvalidStatus:          "The status must be either ACTIVE or INACTIVE. Please enter a valid status.",
	ErrProductNotFound:        "Product not found.",

	// system error
	ErrGetMyShop: "Failed to fetch my shop. Please try again later.",
}

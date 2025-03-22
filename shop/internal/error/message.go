package error

var errorMessage = map[error]string{
	ErrMissingAuthToken:       "Missing authentication token.",
	ErrInvalidAuthTokenFormat: "Invalid authentication token format.",
	ErrGetUserProfile:         "Failed to fetch user profile. Please try again later.",
	ErrNoUserId:               "Unauthorized: missing user ID in token.",
	ErrInvalidUserId:          "Invalid user ID format. Please provide a valid UUID.",
	ErrInvalidShopId:          "Invalid shop ID. Please provide a valid UUID.",
	ErrMissingRequestBody:     "Missing request body.",
	ErrInvalidRequestBody:     "Invalid request body.",
	ErrInvalidName:            "The name must be at least 3 characters long. Please enter a valid name.",
	ErrShopNotFound:           "Shop not found.",
	ErrShopDeactivated:        "Shop is deactivated. Please contact us for reactivation.",
	ErrAlreadyOwnAShop:        "You already own a shop. A user can only create one shop.",
}

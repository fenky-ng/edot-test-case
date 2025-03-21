package error

var errorMessage = map[error]string{
	ErrMissingAuthToken:       "Missing authentication token.",
	ErrInvalidAuthTokenFormat: "Invalid authentication token format.",
	ErrGetUserProfile:         "Failed to fetch user profile. Please try again later.",
	ErrNoUserId:               "Unauthorized: missing user ID in token.",
	ErrInvalidUserId:          "Invalid user ID format. Please provide a valid UUID.",
	ErrMissingRequestBody:     "Missing request body.",
	ErrInvalidRequestBody:     "Invalid request body.",
}

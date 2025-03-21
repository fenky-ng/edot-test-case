package error

var errorMessage = map[error]string{
	ErrMissingAuthToken:           "Missing authentication token.",
	ErrInvalidAuthTokenFormat:     "Invalid authentication token format.",
	ErrInvalidAuthToken:           "Invalid or expired authentication token.",
	ErrJWT:                        "Failed to process authentication token.",
	ErrNoUserId:                   "Unauthorized: missing user ID in token.",
	ErrInvalidUserId:              "Invalid user ID format. Please provide a valid UUID.",
	ErrMissingRequestBody:         "Missing request body.",
	ErrInvalidRequestBody:         "Invalid request body.",
	ErrInvalidRegistrationRequest: "Please provide either a phone number or an email address to register.",
	ErrInvalidLoginRequest:        "Please provide either a phone number or an email address to login.",
	ErrInvalidName:                "The name must be at least 3 characters long. Please enter a valid name.",
	ErrInvalidPassword:            "The password must be at least 6 characters long. Please enter a valid password.",
	ErrUserNotFound:               "User not found. Please check your phone number or email.",
	ErrPhoneRegistered:            "This phone is already associated with an account. Please log in or use a different phone.",
	ErrEmailRegistered:            "This email is already associated with an account. Please log in or use a different email.",
	ErrUserDeactivated:            "User is deactivated. Please contact us for account reactivation.",
	ErrInvalidPhoneLogin:          "Invalid phone or password.",
	ErrInvalidEmailLogin:          "Invalid email or password.",
}

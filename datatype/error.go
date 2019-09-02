package datatype

import "errors"

var (
	//ErrInvalidRequest invalid request
	ErrInvalidRequest = errors.New("Invalid request")

	// ErrServerError for any internal server error
	ErrServerError = errors.New("Server error")

	// ErrBadRequest send it when user tries to do something illegal
	ErrBadRequest = errors.New("Bad request")

	// ErrTryAgain for any try again error
	ErrTryAgain = errors.New("Can't complete this operation at the moment, please try again")

	ErrEmailInvalid           = errors.New("Email address is invalid")
	ErrEmailDoesntExist       = errors.New("Email doesn't exist")
	ErrEmailExists            = errors.New("Email is already registered")
	ErrUsernameExists         = errors.New("Username is already registered")
	ErrPasswordTooShort       = errors.New("Password must be at least 6 characters long")
	ErrPasswordSameAsEmail    = errors.New("Password can not be the same as email")
	ErrAgreeToTerms           = errors.New("You must agree to site terms & privacy")
	ErrInvalidCurrentPassword = errors.New("Invalid current password")
)

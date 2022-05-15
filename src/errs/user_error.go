package errs

import "errors"

// Define business error for user domain
var (
	ErrUserNotFound    = errors.New("user_not_found")
	ErrEmailExisted    = errors.New("email_existed")
	ErrEmailNotFound   = errors.New("email_not_found")
	ErrPasswordInvalid = errors.New("password_invalid")
)

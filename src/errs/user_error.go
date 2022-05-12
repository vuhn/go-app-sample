package errs

import "errors"

// Define business error for user domain
var (
	ErrEmailExisted    = errors.New("email_existed")
	ErrEmailNotFound   = errors.New("email_not_found")
	ErrPasswordInvalid = errors.New("password_invalid")
)

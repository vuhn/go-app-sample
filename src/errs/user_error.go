package errs

import "errors"

// Define business error for user domain
var (
	ErrEmailExisted = errors.New("email_existed")
)

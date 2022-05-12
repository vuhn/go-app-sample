package errs

import "errors"

// Define errors
var (
	ErrUnauthorized       = errors.New("unauthorized_error")
	ErrPermissionDenied   = errors.New("permission_denied_error")
	ErrInternalServer     = errors.New("internal_server_error")
	ErrResourceNotFound   = errors.New("resource_not_found")
	ErrInvalidRequestBody = errors.New("invalid_request_body")
)

package validator

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/dto"
)

// NewAPIValidator returns APIValidator instance
func NewAPIValidator() *APIValidator {
	return &APIValidator{
		validator: validator.New(),
	}
}

// APIValidator is validator for echo framework
type APIValidator struct {
	validator *validator.Validate
}

// Validate is method to validate http request
func (v *APIValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		errs := []string{}
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(errs, "invalid_"+strings.ToLower(err.Field()))
		}
		errorResp := &dto.ErrorResponse{
			Success: false,
			Errors:  errs,
		}
		return echo.NewHTTPError(http.StatusBadRequest, errorResp)
	}
	return nil
}

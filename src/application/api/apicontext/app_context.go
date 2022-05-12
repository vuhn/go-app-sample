package apicontext

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/dto"
)

// APIContextHanlder is middleware function to use APIContext
var APIContextHanlder = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		apiCtx := &APIContext{c}
		return next(apiCtx)
	}
}

// APIContext define custom context for API
type APIContext struct {
	echo.Context
}

// RenderJSON returns reponse for success cases
func (c *APIContext) RenderJSON(v interface{}) error {
	response := &dto.SuccessResponse{
		Success: true,
		Data:    v,
	}
	return c.JSON(http.StatusOK, response)
}

// RenderError returns response for error cases
func (c *APIContext) RenderError(code int, err interface{}) error {
	response := &dto.ErrorResponse{
		Success: false,
		Errors:  err,
	}
	return c.JSON(code, response)
}

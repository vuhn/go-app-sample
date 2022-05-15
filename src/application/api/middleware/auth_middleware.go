package middleware

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/api/appcontext"
	"github.com/vuhn/go-app-sample/application/dto"
	"github.com/vuhn/go-app-sample/pkg/token"
)

// AuthMiddleWare defines  user authentication middleware
type AuthMiddleWare struct {
	jwt token.Token
}

// NewAuthMiddleWare returns a AuthMiddleWare instance
func NewAuthMiddleWare(jwt token.Token) *AuthMiddleWare {
	return &AuthMiddleWare{
		jwt: jwt,
	}
}

// ValidateRequest is middleware to check if a request is authorized
func (a *AuthMiddleWare) ValidateRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if isSkipAuthMiddleWare(c.Request()) {
			return next(c)
		}

		unauthorizedResp := dto.NewErrorResponse("unauthorized")

		token := c.Request().Header.Get("Authorization")
		claims, err := a.jwt.Validate(token)
		if err != nil || claims == nil {
			return c.JSON(http.StatusUnauthorized, unauthorizedResp)
		}

		userID, ok := claims["user_id"]
		if !ok {
			return c.JSON(http.StatusUnauthorized, unauthorizedResp)
		}

		// Pass user id to context
		ctx := context.WithValue(c.Request().Context(), appcontext.UserAuthCtxKey, userID)
		c.Request().WithContext(ctx)

		return next(c)
	}
}

// isSkipAuthMiddleWare is to check if a request is ignored in auth middleware
func isSkipAuthMiddleWare(r *http.Request) bool {
	skipList := map[string]bool{
		"/users":       true,
		"/users/login": true,
	}
	if _, ok := skipList[r.URL.Path]; ok {
		return true
	}
	return false
}

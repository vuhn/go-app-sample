package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/dto"
	"github.com/vuhn/go-app-sample/errs"
	"github.com/vuhn/go-app-sample/pkg/idgenerator"
	"github.com/vuhn/go-app-sample/service"
)

type UserHandler struct {
	userService service.UserService
	idGenerator idgenerator.IDGenerator
}

// NewUserHandler setup rest api handlers for user
func NewUserHandler(echo *echo.Echo,
	userService service.UserService,
	idGenerator idgenerator.IDGenerator,
) {
	handler := &UserHandler{
		userService: userService,
		idGenerator: idGenerator,
	}
	echo.POST("/users", handler.CreateUser)
}

// CreateUser is method for create user endpoint
func (h *UserHandler) CreateUser(c echo.Context) error {
	user := dto.UserRequest{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(errs.ErrInvalidRequestBody.Error()))
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	user.ID = h.idGenerator.GenerateNewID()
	userEntity, err := h.userService.CreateUser(user.ToEntity())
	if err == errs.ErrEmailExisted {
		return c.JSON(http.StatusBadRequest, dto.NewErrorResponse(err.Error()))
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.NewErrorResponse(err.Error()))
	}

	userDto := dto.NewUserResponseFromEntity(userEntity)
	return c.JSON(http.StatusCreated, dto.NewSuccessResponse(userDto))
}

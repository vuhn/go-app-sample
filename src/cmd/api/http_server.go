package main

import (
	"github.com/labstack/echo/v4"
	"github.com/vuhn/go-app-sample/application/api/handler"
	"github.com/vuhn/go-app-sample/infrastructure/repository/postgresrepo"
	"github.com/vuhn/go-app-sample/pkg/idgenerator"
	"github.com/vuhn/go-app-sample/pkg/password"
	"github.com/vuhn/go-app-sample/pkg/token"
	"github.com/vuhn/go-app-sample/service/serviceimpl"
	"gorm.io/gorm"
)

// Dependencies contains dependencies
type Dependencies struct {
	DB          *gorm.DB
	Server      *echo.Echo
	IDGenerator idgenerator.IDGenerator
	Token       token.Token
	Password    password.Password
}

// InitAPIHandler is function to setup api handlers
func InitAPIHandler(deps *Dependencies) {
	db := deps.DB
	server := deps.Server
	idGenerator := deps.IDGenerator

	userRepository := postgresrepo.NewUserRepository(db)
	userService := serviceimpl.NewUserService(userRepository, deps.Token, deps.Password)
	handler.NewUserHandler(server, userService, idGenerator)
}

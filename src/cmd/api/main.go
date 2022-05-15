package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	apiMiddleware "github.com/vuhn/go-app-sample/application/api/middleware"

	"github.com/vuhn/go-app-sample/application/api/validator"
	"github.com/vuhn/go-app-sample/config"
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/infrastructure/db/postgres"
	"github.com/vuhn/go-app-sample/pkg/idgenerator"
	"github.com/vuhn/go-app-sample/pkg/password"
	"github.com/vuhn/go-app-sample/pkg/token"
)

func main() {
	e := echo.New()

	appConfig, err := config.LoadAppConfig()
	if err != nil {
		e.Logger.Error(err)
		panic("failed to load application configuration")
	}

	token := token.NewJWTToken(appConfig.Secret.JWTKey)
	authMiddleware := apiMiddleware.NewAuthMiddleWare(token)
	e.Use(middleware.Logger())
	e.Use(authMiddleware.ValidateRequest)
	e.Validator = validator.NewAPIValidator()

	db, err := postgres.NewDB(appConfig)
	if err != nil {
		e.Logger.Error(err)
		panic("failed to connect database")
	}
	// TODO: move to infrastructure db
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.Task{})

	apiDeps := &Dependencies{
		DB:          db,
		Server:      e,
		IDGenerator: idgenerator.NewIDGenerator(),
		Token:       token,
		Password:    password.NewBcryptPassword(),
	}

	InitAPIHandler(apiDeps)
	e.Logger.Info("Starting server...")
	serverAddress := fmt.Sprintf("0.0.0.0:%s", appConfig.Server.Port)
	e.Logger.Fatal(e.Start(serverAddress))
}

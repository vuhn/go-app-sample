package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/vuhn/go-app-sample/application/api/validator"
	"github.com/vuhn/go-app-sample/config"
	"github.com/vuhn/go-app-sample/entity"
	"github.com/vuhn/go-app-sample/infrastructure/db/postgres"
	"github.com/vuhn/go-app-sample/pkg/idgenerator"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Validator = validator.NewAPIValidator()

	appConfig, err := config.LoadAppConfig()
	if err != nil {
		e.Logger.Error(err)
		panic("failed to load application configuration")
	}

	db, err := postgres.NewDB(appConfig)
	if err != nil {
		e.Logger.Error(err)
		panic("failed to connect database")
	}
	// TODO: move to infrastructure db
	db.AutoMigrate(&entity.User{})

	apiDeps := &Dependencies{
		DB:          db,
		Server:      e,
		IDGenerator: idgenerator.NewIDGenerator(),
	}

	InitAPIHandler(apiDeps)
	e.Logger.Fatal(e.Start(":8080"))
}

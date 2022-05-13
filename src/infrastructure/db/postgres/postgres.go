package postgres

import (
	"fmt"

	"github.com/vuhn/go-app-sample/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(config *config.AppConfig) (*gorm.DB, error) {
	dbConfig := config.Database
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=%s"
	dsn = fmt.Sprintf(dsn, dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

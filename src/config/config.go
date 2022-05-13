package config

import "github.com/kelseyhightower/envconfig"

type AppConfig struct {
	Database Database
	Server   Server
	Secret   Secret
}

// Database contains postgres db config
type Database struct {
	Host       string `envconfig:"DB_HOST" default:"localhost"`
	User       string `envconfig:"DB_USER" default:"postgres"`
	Password   string `envconfig:"DB_PASSWORD" default:"postgres"`
	Port       string `envconfig:"DB_PORT" default:"5432"`
	DBName     string `envconfig:"DB_NAME" default:"go-app-sample"`
	SSLMode    string `envconfig:"DB_SSL_MODE" default:"disable"`
	IsCloudSQL bool   `envconfig:"IS_CLOUD_SQL" default:"false"`
}

type Server struct {
	Port string `envconfig:"PORT" default:"8080"`
	Host string `envconfig:"SERVER_HOST" default:"0.0.0.0"`
}

type Secret struct {
	JWTKey string `envconfig:"JWT_KEY" default:"112233445566"`
}

func LoadAppConfig() (*AppConfig, error) {
	cfg := AppConfig{}
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

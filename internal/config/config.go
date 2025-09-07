package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database Database
}

type Database struct {
	Host     string `envconfig:"DATABASE_HOST" default:"localhost"`
	Port     string `envconfig:"DATABASE_PORT" default:"3306"`
	Username string `envconfig:"DATABASE_USER" default:"postgres"`
	Password string `envconfig:"DATABASE_PASSWORD" default:"postgres"`
	Name     string `envconfig:"DATABASE_NAME" default:"postgresdb"`
}

func Environ() (*Config, error) {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}
	return &config, nil
}

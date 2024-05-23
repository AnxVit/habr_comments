package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Debug struct {
	Logger   bool `env:"DEBUG_LOGGER"`
	Database bool `env:"DEBUG_DATABASE"`
}

type Server struct {
	MainHost string `env:"HOST"         env-default:"0.0.0.0"`
	MainPort int    `env:"PORT"         env-default:"8080"`
}

type Config struct {
	Debug      Debug
	PostgreSQL Postgres
	Server     Server
}

func NewConfig() (*Config, error) {
	cfg := new(Config)

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

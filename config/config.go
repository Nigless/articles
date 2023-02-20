package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App      App
		HTTP     HTTP
		Log      Log
		Postgres Postgres
	}

	App struct {
		Name    string `env:"APP_NAME" env-required:"true"`
		Version string `env:"APP_VERSION" env-required:"true"`
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" env-required:"true"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL" env-required:"true"`
	}

	Postgres struct {
		Username string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Host     string `env:"POSTGRES_HOST" 	 env-required:"true"`
		Port     int    `env:"POSTGRES_PORT" 	 env-required:"true"`
		Database string `env:"POSTGRES_DB" env-required:"true"`
	}
)

// NewConfig returns app config.
func New() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

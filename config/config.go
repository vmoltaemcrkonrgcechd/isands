package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPAddr    string `yaml:"http_addr"`
	PostgresURL string `env:"POSTGRES_URL"`
}

func New() (*Config, error) {
	cfg := new(Config)

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	if err = godotenv.Load(); err != nil {
		return nil, err
	}

	if err = cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

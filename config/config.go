package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Http HTTP    `yaml:"http"`
		Pg   PSQL    `yaml:"postgres"`
		Auth JwtAuth `yaml:"jwt_auth"`
	}

	HTTP struct {
		Port string `env-required:"true" yaml:"port"`
		Host string `env-required:"true" yaml:"host"`
	}

	PSQL struct {
		Host     string `env-required:"true" yaml:"host"`
		Port     string `env-required:"true" yaml:"port"`
		User     string `env-required:"true" yaml:"user"`
		Password string `env-required:"true" yaml:"password"`
		DbName   string `env-required:"true" yaml:"db_name"`
		SslMode  string `env-required:"true" yaml:"ssl_mode"`
	}

	JwtAuth struct {
		Key string `env-required:"true" yaml:"key"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}

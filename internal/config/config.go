package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"PORT" default:"8081"`
	Env  string `envconfig:"ENV" default:"development"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := envconfig.Process("", cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

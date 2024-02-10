package config

import (
	"github.com/epes/econfig"
)

type Config struct {
	Port int `yaml:"port"`
}

func New(environment econfig.Environment) (Config, error) {
	var c Config
	if err := econfig.Populate(&c, "config/", environment); err != nil {
		return Config{}, err
	}

	return c, nil
}

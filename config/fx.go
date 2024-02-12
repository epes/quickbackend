package config

import (
	"github.com/epes/econfig"
	"time"
)

type Config struct {
	Port                      int      `yaml:"port"`
	AccessControlAllowOrigin  []string `yaml:"access_control_allow_origin"`
	AccessControlAllowMethods []string `yaml:"access_control_allow_methods"`
	AccessControlAllowHeaders []string `yaml:"access_control_allow_headers"`
	ServerStartTime           time.Time
}

func New(environment econfig.Environment) (Config, error) {
	var c Config
	if err := econfig.Populate(&c, "config/", environment); err != nil {
		return Config{}, err
	}

	c.ServerStartTime = time.Now()

	return c, nil
}

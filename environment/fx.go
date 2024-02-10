package environment

import (
	"github.com/epes/econfig"
	"os"
)

const key = "QUICK_BACKEND_ENV"

func New() econfig.Environment {
	os.Setenv("TZ", "UTC")

	switch os.Getenv(key) {
	case "base":
		return econfig.EnvironmentBase
	case "dev":
		return econfig.EnvironmentDevelopment
	case "prod":
		return econfig.EnvironmentProduction
	case "test":
		return econfig.EnvironmentTest
	default:
		return econfig.EnvironmentBase
	}
}

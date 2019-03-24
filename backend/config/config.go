package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var Backend = Specification{}

type Specification struct {
	Address          string `default:":8080"`
	DatabaseURL      string `required:"true"`
	DatabaseName     string `required:"true"`
	JWTSecret        string `required:"true"`
	LocationIQAPIKey string `required:"true"`
}

func Load() error {
	err := envconfig.Process("supereasy", &Backend)

	if err != nil {
		return fmt.Errorf("config.LoadConfig: %s", err)
	}

	return nil
}

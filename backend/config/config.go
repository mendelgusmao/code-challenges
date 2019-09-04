package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var Backend Specification

type Specification struct {
	Address  string `default:":9091"`
	Database string `default:"taxchallenge.boltdb"`
}

var hooks []func(Specification) error

func Load() error {
	err := envconfig.Process("taxchallenge", &Backend)

	if err != nil {
		return fmt.Errorf("config.LoadConfig: %s", err)
	}

	for _, hook := range hooks {
		if err := hook(Backend); err != nil {
			return err
		}
	}

	return nil
}

func AfterLoad(hook func(Specification) error) {
	hooks = append(hooks, hook)
}

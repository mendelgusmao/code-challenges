package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

var Backend Specification

type Specification struct {
	Address string `default:":9091"`
	Source  string
}

var hooks []func(Specification) error

func Load() error {
	if err := envconfig.Process("zapchallenge", &Backend); err != nil {
		return errors.Wrap(err, "config.LoadConfig")
	}

	for _, hook := range hooks {
		if err := hook(Backend); err != nil {
			return errors.Wrap(err, "config.LoadConfig (hooks)")
		}
	}

	return nil
}

func Hook(hook func(Specification) error) {
	hooks = append(hooks, hook)
}

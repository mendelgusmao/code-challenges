package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type afterLoadFunc func(*Specification) error

var afterLoad = []afterLoadFunc{}

var Frontend = Specification{}

type Specification struct {
	Address        string `default:":8000"`
	BackendAddress string
	SessionKey     string
}

func Load() error {
	err := envconfig.Process("me_gu_frontend", &Frontend)

	if err != nil {
		return fmt.Errorf("config.LoadConfig: %s", err)
	}

	for _, f := range afterLoad {
		if e := f(&Frontend); e != nil {
			return fmt.Errorf("config.LoadConfig (after load): %s", e)
		}
	}

	return nil
}

func AfterLoad(after afterLoadFunc) {
	afterLoad = append(afterLoad, after)
}

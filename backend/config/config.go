package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

var Backend = Specification{}

type Specification struct {
	Address  string `default:":8080"`
	Database string `required:"true"`
}

func Load() error {
	err := envconfig.Process("supereasy", &Backend)

	if err != nil {
		return fmt.Errorf("config.LoadConfig: %s", err)
	}

	for _, f := range afterLoad {
		if e := f(&Backend); e != nil {
			return fmt.Errorf("config.LoadConfig (after load): %s", e)
		}
	}

	return nil
}

func AfterLoad(after afterLoadFunc) {
	afterLoad = append(afterLoad, after)
}

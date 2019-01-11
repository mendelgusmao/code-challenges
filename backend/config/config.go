package config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type afterLoadFunc func(*Specification) error

var afterLoad = []afterLoadFunc{}

var Backend = Specification{}

type Specification struct {
	Address                  string        `default:":8001"`
	Database                 string        `required:"true"`
	PasswordResetExpiration  time.Duration `default:"12h"`
	PasswordResetFromAddress string
	SMTPAddress              string
	SMTPPort                 int
	SMTPUser                 string
	SMTPPassword             string
}

func Load() error {
	err := envconfig.Process("me_gu_backend", &Backend)

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

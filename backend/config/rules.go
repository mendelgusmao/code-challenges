package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

var TaxRules map[string]string

func loadRules(filename string) error {
	if filename == "" {
		return nil
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.Wrap(err, "config.loadRules")
	}

	if err := json.Unmarshal(content, &TaxRules); err != nil {
		return errors.Wrap(err, "config.loadRules")
	}

	return nil
}

func init() {
	AfterLoad(func(backend Specification) error {
		return loadRules(backend.TaxRules)
	})
}

package config

import (
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
	"github.com/pkg/errors"
)

type PortalsSpecification struct {
	BoundingBox model.BoundingBox             `yaml:"bounding_box"`
	PortalRules map[string]*model.PortalRules `yaml:"portal_rules"`
}

var Portals = PortalsSpecification{}

func loadPortals(filename string) error {
	if filename == "" {
		return nil
	}

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.Wrap(err, "config.loadPortals")
	}

	if err := yaml.Unmarshal(content, &Portals); err != nil {
		return errors.Wrap(err, "config.loadPortals")
	}

	return buildPortal()
}

func buildPortal() error {
	for portal, rules := range Portals.PortalRules {
		if err := rules.BuildExpression(Portals.BoundingBox); err != nil {
			return errors.Wrapf(err, "config.buildPortal (%s)", portal)
		}
	}

	return nil
}

func init() {
	Hook(func(backend Specification) error {
		return loadPortals(backend.Portals)
	})
}

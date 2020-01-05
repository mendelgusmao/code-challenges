package model

import (
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/pkg/errors"
)

var errExpressionShouldReturnBoolean = errors.New("expression should return a boolean value")

type PortalRules struct {
	Rules               []string                       `yaml:"rules"`
	EvaluableExpression *govaluate.EvaluableExpression `yaml:"-"`
}

func (r *PortalRules) BuildExpression(boundingBox BoundingBox) error {
	var err error

	combinedRules := strings.Join(r.Rules, " && ")
	functions := map[string]govaluate.ExpressionFunction{
		"insideBoundingBox": func(args ...interface{}) (interface{}, error) {
			values := make([]float64, len(args))

			for i, arg := range args {
				values[i] = arg.(float64)
			}

			lat := args[0].(float64)
			lon := args[1].(float64)

			return lat >= boundingBox.MinLat && lat <= boundingBox.MinLat &&
				lon >= boundingBox.MinLon && lon <= boundingBox.MinLon, nil
		},
	}

	r.EvaluableExpression, err =
		govaluate.NewEvaluableExpressionWithFunctions(combinedRules, functions)

	if err != nil {
		return errors.Wrap(err, "PortalRules.BuildExpression")
	}

	return nil
}

func (r *PortalRules) Test(listing Listing) (bool, error) {
	params := listing.ToMap()
	result, err := r.EvaluableExpression.Evaluate(params)

	if err != nil {
		return false, errors.Wrap(err, "PortalRules.Test")
	}

	truthy, ok := result.(bool)

	if !ok {
		return false,
			errors.Wrap(errExpressionShouldReturnBoolean, "PortalRules.Test")
	}

	return truthy, nil
}

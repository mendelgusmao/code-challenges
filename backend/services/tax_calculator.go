package services

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/pkg/errors"
)

type Rules map[string]string

type taxCalculator struct {
	rules Rules
}

func NewTaxCalculator(rules Rules) taxCalculator {
	return taxCalculator{
		rules: rules,
	}
}

func (c *taxCalculator) Calculate(product Product) (float64, error) {
	for rule, formula := range c.rules {
		result, err := c.evaluate(rule, product)

		if err != nil {
			return 0.0, errors.Wrap(err, "taxCalculator.calculate")
		}

		if v, ok := result.(bool); !ok {
			return 0.0, fmt.Errorf("rule should be a bool expression")
		} else if ok && !v {
			continue
		}

		result, formulaErr := c.evaluate(formula, product)

		if formulaErr != nil {
			return 0.0, errors.Wrap(formulaErr, "taxCalculator.calculate")
		}

		return result.(float64), nil
	}

	return 0.0, nil
}

func (c *taxCalculator) evaluate(expression string, product Product) (interface{}, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)

	if err != nil {
		return nil, errors.Wrap(err, "taxCalculator.evaluate")
	}

	result, err := expr.Evaluate(product.toMap())

	if err != nil {
		return nil, errors.Wrap(err, "taxCalculator.evaluate")
	}

	return result, nil
}

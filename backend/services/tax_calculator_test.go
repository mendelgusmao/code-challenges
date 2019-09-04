package services

import (
	"testing"
)

func TestTaxCalculator(t *testing.T) {
	rules := Rules{
		"price > 6750": "price / 100 * 15.5",
	}

	product := Product{
		Price: 10000,
	}

	expectedResult := product.Price / 100 * 15.5
	calculator := NewTaxCalculator(rules)
	result, err := calculator.Calculate(product)

	if err != nil {
		t.Error(err)
	}

	if result != expectedResult {
		t.Fatalf("expecting result to be %0.2f, got %0.2f", expectedResult, result)
	}
}

func TestTaxCalculatorEmptyRules(t *testing.T) {
	rules := Rules{}

	product := Product{}
	calculator := NewTaxCalculator(rules)

	value, err := calculator.Calculate(product)

	if value != 0.0 || err != nil {
		t.Fatal("expecting returned value to be 0.0 with no error")
	}
}

func TestTaxCalculatorRuleError(t *testing.T) {
	rules := Rules{
		"1 + 1%": "price / 100 * 15.5",
	}

	product := Product{}
	calculator := NewTaxCalculator(rules)

	_, err := calculator.Calculate(product)

	if err == nil {
		t.Fatal("expecting error")
	}
}

func TestTaxCalculatorFormulaError(t *testing.T) {
	rules := Rules{
		"2 > 1": "invalid",
	}

	product := Product{}
	calculator := NewTaxCalculator(rules)

	_, err := calculator.Calculate(product)

	if err == nil {
		t.Fatal("expecting error")
	}
}

func TestTaxCalculatorRuleReturnType(t *testing.T) {
	rules := Rules{
		"1 + 1": "",
	}

	product := Product{}
	calculator := NewTaxCalculator(rules)

	value, err := calculator.Calculate(product)

	if value != 0.0 || err == nil {
		t.Fatal("expecting returned value to be 0.0 with no error")
	}
}

func TestTaxCalculatorRuleFalseReturn(t *testing.T) {
	rules := Rules{
		"1 > 2": "",
	}

	product := Product{}
	calculator := NewTaxCalculator(rules)

	value, err := calculator.Calculate(product)

	if value != 0.0 || err != nil {
		t.Fatal("expecting returned value to be 0.0 with no error")
	}
}

package model

import (
	"testing"
)

func TestPortalRulesBuildExpression(t *testing.T) {
	scenarios := []struct {
		rules          []string
		expectingError bool
	}{
		{
			rules: []string{"1 > 0"},
		},
		{
			rules:          []string{"..."},
			expectingError: true,
		},
	}

	for i, scenario := range scenarios {
		portalRule := PortalRules{
			Rules: scenario.rules,
		}

		err := portalRule.BuildExpression(BoundingBox{})

		if scenario.expectingError && err == nil {
			t.Fatalf("scenario %d failed: expecting error", i)
		}
	}
}

func TestPortalRulesTest(t *testing.T) {
	scenarios := []struct {
		rules          []string
		listing        Listing
		expectedResult bool
		expectingError bool
	}{
		{
			rules: []string{"parkingSpaces > 0"},
			listing: Listing{
				ParkingSpaces: 1,
			},
			expectedResult: true,
		},
		{
			rules: []string{"bedrooms == 1"},
			listing: Listing{
				Bedrooms: 3,
			},
			expectedResult: false,
		},
		{
			rules:          []string{"non_existent_field == 1"},
			listing:        Listing{},
			expectingError: true,
		},
		{
			rules:          []string{"1 + 1"},
			listing:        Listing{},
			expectingError: true,
		},
	}

	for i, scenario := range scenarios {
		portalRule := &PortalRules{
			Rules: scenario.rules,
		}

		if err := portalRule.BuildExpression(BoundingBox{}); err != nil {
			t.Fatal(err)
		}

		truthy, err := portalRule.Test(scenario.listing)

		if scenario.expectingError && err == nil {
			t.Fatalf("scenario %d failed: expecting error", i)
		}

		if !scenario.expectingError && err != nil {
			t.Fatalf("scenario %d failed: not expecting error `%v`", i, err)
		}

		if scenario.expectedResult != truthy {
			t.Fatalf("scenario %d failed: expecting result to be `%v`, got '%v'", i, scenario.expectedResult, truthy)
		}
	}
}

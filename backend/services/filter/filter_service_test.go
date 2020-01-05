package filter

import (
	"reflect"
	"testing"

	"github.com/mendelgusmao/zap-challenge/backend/services/model"
)

func TestFilterServiceApply(t *testing.T) {
	scenarios := []struct {
		portalRules      *model.PortalRules
		listings         []model.Listing
		expectedListings []model.Listing
		expectingError   bool
	}{
		{
			portalRules: &model.PortalRules{
				Rules: []string{"bedrooms > 1"},
			},
			listings: []model.Listing{
				model.Listing{
					Bedrooms: 3,
				},
			},
			expectedListings: []model.Listing{
				model.Listing{
					Bedrooms: 3,
				},
			},
		},
		{
			portalRules: &model.PortalRules{
				Rules: []string{"bedrooms > 1"},
			},
			listings: []model.Listing{
				model.Listing{
					Bedrooms: 1,
				},
			},
			expectedListings: []model.Listing{},
		},
		{
			portalRules: &model.PortalRules{
				Rules: []string{"non_existent_field == 1"},
			},
			listings: []model.Listing{
				model.Listing{
					Bedrooms: 1,
				},
			},
			expectingError:   true,
			expectedListings: []model.Listing{},
		},
	}

	for i, scenario := range scenarios {
		if err := scenario.portalRules.BuildExpression(); err != nil {
			t.Error(err)
		}

		filterService := NewFilterService(scenario.portalRules)
		filteredListings, err := filterService.Apply(scenario.listings)

		if scenario.expectingError && err == nil {
			t.Fatalf("scenario %d failed: expecting error", i)
		}

		if !scenario.expectingError && err != nil {
			t.Fatalf("scenario %d failed: not expecting error `%v`", i, err)
		}

		if !reflect.DeepEqual(scenario.expectedListings, filteredListings) {
			t.Fatalf("scenario %d failed: expecting result to be `%v`, got '%v'", i, scenario.expectedListings, filteredListings)
		}
	}
}

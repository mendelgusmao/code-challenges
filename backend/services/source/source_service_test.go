package source

import (
	"errors"
	"reflect"
	"testing"

	"github.com/mendelgusmao/zap-challenge/backend/services/model"
)

type DummyFetcher struct{}

func (f DummyFetcher) Fetch(string) ([]model.Listing, error) {
	return []model.Listing{
		model.Listing{
			ID: "dummy id'",
		},
	}, nil
}

type DummyErrorFetcher struct{}

func (f DummyErrorFetcher) Fetch(string) ([]model.Listing, error) {
	return []model.Listing{}, errors.New("dummy error")
}

func TestFetchSourceService(t *testing.T) {
	scenarios := []struct {
		fetcher         Fetcher
		expectedListing []model.Listing
		expectedError   error
	}{
		{
			fetcher: DummyFetcher{},
			expectedListing: []model.Listing{
				model.Listing{
					ID: "dummy id'",
				},
			},
		},
		{
			fetcher:         DummyErrorFetcher{},
			expectedError:   errors.New("dummy error"),
			expectedListing: []model.Listing{},
		},
	}

	for i, scenario := range scenarios {
		sourceService := NewSourceService("", scenario.fetcher)
		listing, err := sourceService.Fetch()

		if scenario.expectedError != nil && err == nil {
			t.Fatalf("scenario %d expecting error `%v`, got nil", i, scenario.expectedError)
		}

		if !reflect.DeepEqual(scenario.expectedListing, listing) {
			t.Fatalf("scenario %d expecting listing `%#+v`, got `%#+v`", i, scenario.expectedListing, listing)
		}
	}
}

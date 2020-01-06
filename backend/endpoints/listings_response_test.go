package endpoints

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/davecgh/go-spew/spew"
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
)

type listingsResponseTestScenario struct {
	listings         []model.Listing
	page, size       int64
	expectedResponse *ListingsResponse
}

func TestListingsResponsePaginate(t *testing.T) {
	scenarios := []listingsResponseTestScenario{
		{
			listings: generateListings(1, 30),
			expectedResponse: &ListingsResponse{
				PageSize:   10,
				PageNumber: 1,
				TotalCount: 30,
				Listings:   generateListings(1, 10),
			},
			page: 1,
			size: 10,
		},
	}

	for i, scenario := range scenarios {
		response := NewListingsResponse(scenario.listings).Paginate(scenario.page, scenario.size)

		if !reflect.DeepEqual(response, scenario.expectedResponse) {
			responseText := spew.Sdump(response)
			expectedResponseText := spew.Sdump(scenario.expectedResponse)

			t.Log(diff.LineDiff(expectedResponseText, responseText))
			t.Fatalf("scenario %d failed", i)
		}
	}
}

func generateListings(start, end int) []model.Listing {
	listings := make([]model.Listing, 0)

	for i := start; i <= end; i++ {
		listings = append(listings, model.Listing{
			ID: fmt.Sprintf("%d", i+1),
		})
	}

	return listings
}

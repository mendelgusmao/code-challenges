package source

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/mendelgusmao/zap-challenge/backend/services/model"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}

type ContentGetter struct {
	content string
}

func NewContentGetter(content string) *ContentGetter {
	return &ContentGetter{
		content: content,
	}
}

func (g *ContentGetter) Get(string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body: nopCloser{
			bytes.NewBufferString(g.content),
		},
	}, nil
}

type ErrorGetter struct{}

func (g *ErrorGetter) Get(string) (*http.Response, error) {
	return nil, errors.New("dummy error")
}

func TestHTTPFetcherFetch(t *testing.T) {
	scenarios := []struct {
		fetcher          *HTTPFetcher
		expectedListings []model.Listing
		expectingError   bool
	}{
		{
			fetcher: &HTTPFetcher{
				client: NewContentGetter(`[
					{
						"id": "some id"
					}
				]`),
			},
			expectedListings: []model.Listing{
				model.Listing{
					ID: "some id",
				},
			},
		},
		{
			fetcher: &HTTPFetcher{
				client: NewContentGetter(`invalid`),
			},
			expectingError:   true,
			expectedListings: []model.Listing{},
		},
		{
			fetcher: &HTTPFetcher{
				client: &ErrorGetter{},
			},
			expectingError:   true,
			expectedListings: []model.Listing{},
		},
	}

	for i, scenario := range scenarios {
		httpFetcher := scenario.fetcher
		listings, err := httpFetcher.Fetch("")

		if scenario.expectingError && err == nil {
			t.Fatalf("scenario %d failed: expecting error", i)
		}

		if !scenario.expectingError && err != nil {
			t.Fatalf("scenario %d failed: not expecting error `%v`", i, err)
		}

		if !reflect.DeepEqual(scenario.expectedListings, listings) {
			t.Fatalf("scenario %d failed: expecting `%v`, got `%v`", i, scenario.expectedListings, listings)
		}
	}
}

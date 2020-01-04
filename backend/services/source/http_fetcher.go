package source

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gregjones/httpcache"
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
	"github.com/pkg/errors"
)

type Getter interface {
	Get(string) (*http.Response, error)
}

type HTTPFetcher struct {
	client Getter
}

func NewHTTPFetcher() *HTTPFetcher {
	return &HTTPFetcher{
		client: &http.Client{
			Timeout: 60 * time.Second,
			Transport: httpcache.NewTransport(
				httpcache.NewMemoryCache(),
			),
		},
	}
}

func (f *HTTPFetcher) Fetch(url string) ([]model.Listing, error) {
	listing := make([]model.Listing, 0)
	response, err := f.client.Get(url)

	if err != nil {
		return listing, errors.Wrap(err, "HTTPFetcher.Fetch")
	}

	if err := json.NewDecoder(response.Body).Decode(&listing); err != nil {
		return listing, errors.Wrap(err, "HTTPFetcher.Fetch")
	}

	return listing, nil
}

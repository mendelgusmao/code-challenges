package source

import (
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
	"github.com/pkg/errors"
)

type SourceService struct {
	fetcher Fetcher
	url     string
}

func NewSourceService(url string, fetcher Fetcher) *SourceService {
	return &SourceService{
		url:     url,
		fetcher: fetcher,
	}
}

func (s *SourceService) Fetch() ([]model.Listing, error) {
	listing, err := s.fetcher.Fetch(s.url)

	if err != nil {
		return listing, errors.Wrap(err, "SourceService.Fetch")
	}

	return listing, nil
}

package source

import "github.com/mendelgusmao/zap-challenge/backend/services/model"

type Fetcher interface {
	Fetch(string) ([]model.Listing, error)
}

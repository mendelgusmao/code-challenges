package filter

import (
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
	"github.com/pkg/errors"
)

type FilterService struct {
	portalRules *model.PortalRules
}

func NewFilterService(portalRules *model.PortalRules) *FilterService {
	return &FilterService{
		portalRules: portalRules,
	}
}

func (s *FilterService) Apply(listings []model.Listing) ([]model.Listing, error) {
	filteredListings := make([]model.Listing, 0)

	for _, listing := range listings {
		truthy, err := s.portalRules.Test(listing)

		if err != nil {
			return []model.Listing{}, errors.Wrap(err, "FilterService.Apply")
		}

		if truthy {
			filteredListings = append(filteredListings, listing)
		}
	}

	return filteredListings, nil
}

package endpoints

import (
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
)

type ListingsResponse struct {
	PageNumber int32           `json:"pageNumber,omitempty"`
	PageSize   int32           `json:"pageSize,omitempty"`
	TotalCount int32           `json:"totalCount,omitempty"`
	Listings   []model.Listing `json:"listings"`
}

func NewListingsResponse(listings []model.Listing) *ListingsResponse {
	return &ListingsResponse{
		Listings: listings,
	}
}

func (r *ListingsResponse) Paginate(page, size int64) *ListingsResponse {
	start := (page - 1) * size
	total := int64(len(r.Listings))

	if start > total {
		start = total
	}

	end := start + size

	if end > total {
		end = total
	}

	return &ListingsResponse{
		PageNumber: int32(page),
		PageSize:   int32(size),
		TotalCount: int32(total),
		Listings:   r.Listings[start:end],
	}
}

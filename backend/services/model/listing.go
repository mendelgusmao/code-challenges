package model

import (
	"strconv"
	"time"
)

type Listing struct {
	UsableAreas   int       `json:"usableAreas"`
	ListingType   string    `json:"listingType"`
	CreatedAt     time.Time `json:"createdAt"`
	ListingStatus string    `json:"listingStatus"`
	ID            string    `json:"id"`
	ParkingSpaces int       `json:"parkingSpaces"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Owner         bool      `json:"owner"`
	Images        []string  `json:"images"`
	Address       struct {
		City         string `json:"city"`
		Neighborhood string `json:"neighborhood"`
		GeoLocation  struct {
			Precision string `json:"precision"`
			Location  struct {
				Lon float64 `json:"lon"`
				Lat float64 `json:"lat"`
			} `json:"location"`
		} `json:"geoLocation"`
	} `json:"address"`
	Bathrooms    int `json:"bathrooms"`
	Bedrooms     int `json:"bedrooms"`
	PricingInfos struct {
		YearlyIptu      string `json:"yearlyIptu"`
		Price           string `json:"price"`
		BusinessType    string `json:"businessType"`
		MonthlyCondoFee string `json:"monthlyCondoFee"`
	} `json:"pricingInfos"`
}

func (l Listing) ToMap() map[string]interface{} {
	yearlyIPTU, _ := strconv.ParseFloat(l.PricingInfos.YearlyIptu, 64)
	price, _ := strconv.ParseFloat(l.PricingInfos.Price, 64)
	monthlyCondoFee, _ := strconv.ParseFloat(l.PricingInfos.MonthlyCondoFee, 64)

	return map[string]interface{}{
		"usableAreas":                 l.UsableAreas,
		"listingType":                 l.ListingType,
		"createdAt":                   l.CreatedAt,
		"listingStatus":               l.ListingStatus,
		"id":                          l.ID,
		"parkingSpaces":               l.ParkingSpaces,
		"updatedAt":                   l.UpdatedAt,
		"owner":                       l.Owner,
		"addressCity":                 l.Address.City,
		"addressNeighborhood":         l.Address.Neighborhood,
		"addressGeoPrecision":         l.Address.GeoLocation.Precision,
		"addressGeoLon":               l.Address.GeoLocation.Location.Lon,
		"addressGeoLat":               l.Address.GeoLocation.Location.Lat,
		"bathrooms":                   l.Bathrooms,
		"bedrooms":                    l.Bedrooms,
		"pricingInfosYearlyIptu":      yearlyIPTU,
		"pricingInfosPrice":           price,
		"pricingInfosBusinessType":    l.PricingInfos.BusinessType,
		"pricingInfosMonthlyCondoFee": monthlyCondoFee,
	}
}

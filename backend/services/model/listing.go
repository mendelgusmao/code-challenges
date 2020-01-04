package model

import "time"

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

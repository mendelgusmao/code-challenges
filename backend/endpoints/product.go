package endpoints

import (
	"strconv"
	"strings"

	"github.com/mendelgusmao/tax-challenge/backend/services"
)

type product struct {
	Name   string `json:"name"`
	Height string `json:"height"`
	Length string `json:"length"`
	Width  string `json:"width"`
	Weight string `json:"weight"`
	Price  string `json:"price"`
	Tax    string `json:"tax"`
}

func (p product) toServicesProduct() services.Product {
	height, _ := strconv.ParseFloat(strings.Replace(p.Height, ",", ".", -1), 64)
	length, _ := strconv.ParseFloat(strings.Replace(p.Length, ",", ".", -1), 64)
	width, _ := strconv.ParseFloat(strings.Replace(p.Width, ",", ".", -1), 64)
	weight, _ := strconv.ParseFloat(strings.Replace(p.Weight, ",", ".", -1), 64)
	price, _ := strconv.ParseFloat(p.Price, 64)

	return services.Product{
		Name:   p.Name,
		Height: height,
		Length: length,
		Width:  width,
		Weight: weight,
		Price:  price,
	}
}

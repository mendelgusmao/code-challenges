package endpoints

import (
	"fmt"
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

func (p *product) fromTaxedProduct(servicesProduct services.TaxedProduct) {
	*p = product{
		Name:   servicesProduct.Name,
		Height: strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Height), ".", ",", -1),
		Length: strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Length), ".", ",", -1),
		Width:  strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Width), ".", ",", -1),
		Weight: strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Weight), ".", ",", -1),
		Price:  strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Price), ".", ",", -1),
		Tax:    strings.Replace(fmt.Sprintf("%.2f", servicesProduct.Tax), ".", ",", -1),
	}
}

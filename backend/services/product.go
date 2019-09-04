package services

type Product struct {
	Name   string  `json:"name"`
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Width  float64 `json:"width"`
	Weight float64 `json:"weight"`
	Price  float64 `json:"price"`
}

type TaxedProduct struct {
	Product
	ID  uint64  `json:"id,omitempty"`
	Tax float64 `json:"tax,omitempty"`
}

func (p Product) toMap() map[string]interface{} {
	return map[string]interface{}{
		"name":   p.Name,
		"height": p.Height,
		"length": p.Length,
		"width":  p.Width,
		"weight": p.Weight,
		"price":  p.Price,
	}
}

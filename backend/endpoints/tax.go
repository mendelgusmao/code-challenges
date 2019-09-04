package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/tax-challenge/backend/config"
	"github.com/mendelgusmao/tax-challenge/backend/middleware"
	"github.com/mendelgusmao/tax-challenge/backend/router"
	"github.com/mendelgusmao/tax-challenge/backend/services"
)

func init() {
	router.Router.HandleFunc("/tax", postTax).Methods("POST")
}

func postTax(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := context.Get(r, "jsonDecoder").(middleware.JSONDecoderFunc)
	product := product{}

	if !jsonDecoder(&product) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taxCalculator := services.NewTaxCalculator(config.TaxRules)
	tax, err := taxCalculator.Calculate(product.toServicesProduct())

	if err != nil {
		log.Println("postTax:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"tax": strings.Replace(fmt.Sprintf("%0.2f", tax), ".", ",", -1),
	}

	json.NewEncoder(w).Encode(response)
}

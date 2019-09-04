package endpoints

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/mendelgusmao/tax-challenge/backend/config"
	"github.com/mendelgusmao/tax-challenge/backend/middleware"
	"github.com/mendelgusmao/tax-challenge/backend/router"
	"github.com/mendelgusmao/tax-challenge/backend/services"
	"go.etcd.io/bbolt"
)

func init() {
	router.Router.HandleFunc("/track", postTrack).Methods("POST")
	router.Router.HandleFunc("/track/{id}", getTrack).Methods("GET")
}

func postTrack(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := context.Get(r, "jsonDecoder").(middleware.JSONDecoderFunc)
	db := context.Get(r, "db").(*bbolt.DB)
	product := product{}

	if !jsonDecoder(&product) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	servicesProduct := product.toServicesProduct()
	taxCalculator := services.NewTaxCalculator(config.TaxRules)
	tax, err := taxCalculator.Calculate(servicesProduct)

	if err != nil {
		log.Println("postTrack:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productStorage := services.NewProductStorage(db)
	id, err := productStorage.Store(services.TaxedProduct{
		Product: servicesProduct,
		Tax:     tax,
	})

	if err != nil {
		log.Println("postTrack:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"id": fmt.Sprintf("%d", id),
	})
}

func getTrack(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*bbolt.DB)
	productStorage := services.NewProductStorage(db)

	id, _ := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)

	taxedProduct, err := productStorage.Retrieve(id)
	product := product{}
	product.fromTaxedProduct(taxedProduct)

	if err != nil {
		if strings.Contains(err.Error(), "EOF") {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		log.Println("postTrack:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

package endpoints

import (
	"net/http"

	"github.com/gorilla/context"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/middleware"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/router"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"
	"go.etcd.io/bbolt"
)

func init() {
	router.Router.HandleFunc("/cars", putCars).Methods("PUT")
}

func putCars(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := context.Get(r, "jsonDecoder").(middleware.JSONDecoderFunc)
	db := context.Get(r, "db").(*bbolt.DB)
	service := services.NewCarsService(db)

	cars := []services.Car{}

	if jsonDecoder(&cars) {
		if err := service.Clear(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := service.Put(cars); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

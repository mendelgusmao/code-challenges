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
	carsService := services.NewCarsService(db)
	journeysService := services.NewJourneysService(db)

	cars := []services.Car{}

	if jsonDecoder(&cars) {
		if err := journeysService.Clear(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := carsService.Clear(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := carsService.Put(cars); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

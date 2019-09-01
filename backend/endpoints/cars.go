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

	cars := []services.Car{}

	if jsonDecoder(&cars) {
		carsService := services.NewCarsService(db)

		if err := clearPool(db); err != nil {
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

func clearPool(db *bbolt.DB) error {
	journeysService := services.NewJourneysService(db)
	carsService := services.NewCarsService(db)
	tripsService := services.NewTripsService(db)

	if err := journeysService.Clear(); err != nil {
		return err
	}

	if err := carsService.Clear(); err != nil {
		return err
	}

	if err := tripsService.Clear(); err != nil {
		return err
	}

	return nil
}

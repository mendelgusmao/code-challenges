package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/middleware"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/router"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"
	"go.etcd.io/bbolt"
)

func init() {
	router.Router.HandleFunc("/locate", postLocate).Methods("POST").
		MatcherFunc(middleware.IsURLEncodedForm)
}

func postLocate(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*bbolt.DB)
	journeysService := services.NewJourneysService(db)
	tripsService := services.NewTripsService(db)
	carsService := services.NewCarsService(db)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	journeyID, err := strconv.Atoi(r.FormValue("ID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if _, err := journeysService.Find(journeyID); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	trip, tripErr := tripsService.FindByJourneyID(journeyID)

	if tripErr != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	car, carErr := carsService.Find(trip.CarID)

	if carErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}

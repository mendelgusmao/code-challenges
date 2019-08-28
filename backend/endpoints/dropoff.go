package endpoints

import (
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/middleware"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/router"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"
	"go.etcd.io/bbolt"
)

func init() {
	router.Router.HandleFunc("/dropoff", postDropoff).Methods("POST").
		MatcherFunc(middleware.IsURLEncodedForm)
}

func postDropoff(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*bbolt.DB)
	journeysService := services.NewJourneysService(db)

	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	journeyID, err := strconv.Atoi(r.FormValue("ID"))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, findErr := journeysService.Find(journeyID)

	if findErr != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

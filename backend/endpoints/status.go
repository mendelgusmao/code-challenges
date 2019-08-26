package endpoints

import (
	"net/http"

	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/router"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/services"
)

func init() {
	router.Router.HandleFunc("/status", getStatus).Methods("GET")
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	if services.Running() {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusServiceUnavailable)
}

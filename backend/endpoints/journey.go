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
	router.Router.HandleFunc("/journey", postJourney).Methods("POST")
}

func postJourney(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := context.Get(r, "jsonDecoder").(middleware.JSONDecoderFunc)
	db := context.Get(r, "db").(*bbolt.DB)
	journeysService := services.NewJourneysService(db)

	journey := services.Journey{}

	if jsonDecoder(&journey) {
		if err := journeysService.Insert(journey); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
}

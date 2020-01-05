package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mendelgusmao/zap-challenge/backend/config"
	"github.com/mendelgusmao/zap-challenge/backend/router"
	"github.com/mendelgusmao/zap-challenge/backend/services/filter"
	"github.com/mendelgusmao/zap-challenge/backend/services/source"
)

func init() {
	router.Router.HandleFunc("/listings/{portal}", getListings).Methods("GET")
}

func getListings(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	portal := vars["portal"]
	portalRules, ok := config.Portals.PortalRules[portal]

	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	httpFetcher := source.NewHTTPFetcher()
	sourceService := source.NewSourceService(config.Backend.Source, httpFetcher)
	listings, err := sourceService.Fetch()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	filterService := filter.NewFilterService(portalRules)
	filteredListings, err := filterService.Apply(listings)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filteredListings)
}

package endpoints

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mendelgusmao/zap-challenge/backend/config"
	"github.com/mendelgusmao/zap-challenge/backend/router"
	"github.com/mendelgusmao/zap-challenge/backend/services/filter"
	"github.com/mendelgusmao/zap-challenge/backend/services/model"
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

	paginatedResponse := paginateResponse(r, filteredListings)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paginatedResponse)
}

func paginateResponse(r *http.Request, listings []model.Listing) *ListingsResponse {
	var (
		page int64 = 1
		size int64 = 10
	)

	r.ParseForm()

	sizeFormValue, err := strconv.ParseInt(r.FormValue("size"), 10, 32)

	if err == nil && sizeFormValue >= 10 && sizeFormValue <= 100 {
		size = sizeFormValue
	}

	pageFormValue, err := strconv.ParseInt(r.FormValue("page"), 10, 32)

	if err == nil && pageFormValue > 0 {
		page = pageFormValue
	}

	return NewListingsResponse(listings).Paginate(page, size)
}

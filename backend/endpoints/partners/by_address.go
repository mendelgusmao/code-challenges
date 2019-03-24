package partners

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	subrouter.HandleFunc("/by_address", findByAddress).Methods("GET")
	subrouter.Use(middleware.Geocoder)
}

func findByAddress(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mongo.Database)
	error := context.Get(r, "error").(middleware.ErrorFunc)
	geocoder := context.Get(r, "geocoder").(middleware.GeocoderFunc)
	dao := newDAO(db)

	lat, long, err := geocoder(r.FormValue("address"))

	if err != nil {
		log.Printf("partners.findByAddress: %s", err)
		error(http.StatusInternalServerError, "unexpected error")

		return
	}

	partners, err := dao.findByLocation(lat, long, 5.0)

	if err != nil {
		log.Printf("partners.findByAddress: %s", err)
		error(http.StatusInternalServerError, "unexpected error")

		return
	}

	if len(partners) == 0 {
		error(http.StatusNotFound, "no partners found")

		return
	}

	simplifiedPartners := make([]SimplifiedPartner, len(partners))

	for i, partner := range partners {
		simplifiedPartners[i] = partner.Simplified()
	}

	json.NewEncoder(w).Encode(simplifiedPartners)
}

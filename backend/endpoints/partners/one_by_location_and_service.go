package partners

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

const distance = 10.0 // km

func init() {
	subrouter.HandleFunc("/one_by_location_and_service", findOneByLocationAndService).Methods("GET")
}

func findOneByLocationAndService(w http.ResponseWriter, r *http.Request) {
	db := context.Get(r, "db").(*mongo.Database)
	error := context.Get(r, "error").(middleware.ErrorFunc)
	dao := newDAO(db)

	lat, long, err := parseCoordinates(r.FormValue("coordinates"))

	if err != nil {
		log.Printf("partners.findOneByLocationAndService: %s", err)
		error(http.StatusBadRequest, "invalid coordinates")

		return
	}

	service := r.FormValue("service")
	partners, err := dao.findByLocationAndService(service, lat, long, distance)

	if err != nil {
		log.Printf("partners.findOneByLocationAndService: %s", err)
		error(http.StatusInternalServerError, "unexpected error")

		return
	}

	if len(partners) == 0 {
		error(http.StatusNotFound, "no partners found")

		return
	}

	json.NewEncoder(w).Encode(partners[0])
}

func parseCoordinates(coords string) (float64, float64, error) {
	coordinates := strings.Split(coords, ",")

	if len(coordinates) != 2 {
		return 0, 0, errors.New("invalid coordinates")
	}

	lat, err := strconv.ParseFloat(coordinates[0], 64)

	if err != nil {
		return 0, 0, fmt.Errorf("parseCoordinates:, %s", err)
	}

	long, err := strconv.ParseFloat(coordinates[1], 64)

	if err != nil {
		return 0, 0, fmt.Errorf("parseCoordinates: %s", err)
	}

	return lat, long, nil
}

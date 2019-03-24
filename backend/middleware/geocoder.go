package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/config"
	"github.com/mendelgusmao/supereasy/backend/lib/locationiq"
)

type GeocoderFunc func(string) (float64, float64, error)

func Geocoder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "geocoder", GeocoderFunc(func(address string) (float64, float64, error) {
			return locationiq.NewClient(config.Backend.LocationIQAPIKey).Geocode(address)
		}))

		next.ServeHTTP(w, r)
	})
}

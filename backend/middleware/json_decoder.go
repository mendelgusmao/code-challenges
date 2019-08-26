package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type JSONDecoderFunc func(interface{}) bool

func JSONDecoder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "jsonDecoder", JSONDecoderFunc(func(target interface{}) bool {
			if err := json.NewDecoder(r.Body).Decode(target); err != nil {
				log.Printf("middleware.JSONDecoder: %s", err)
				w.WriteHeader(http.StatusBadRequest)

				return false
			}

			return true
		}))

		next.ServeHTTP(w, r)
	})
}

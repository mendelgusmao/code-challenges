package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/context"
)

type JSONDecoderFunc func(target interface{}) bool

func JSONDecoder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "jsonDecoder", JSONDecoderFunc(func(target interface{}) bool {
			if err := json.NewDecoder(r.Body).Decode(target); err != nil {
				log.Printf("middleware.JSONDecoder: %s", err)
				http.Error(w, "", http.StatusBadRequest)
				json.NewEncoder(w).Encode(errorMessage{Error: "invalid payload"})

				return false
			}

			return true
		}))

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

type errorMessage struct {
	Error string `json:"error"`
}

type ErrorFunc func(int, string)

func Error(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "error", ErrorFunc(func(status int, message string) {
			http.Error(w, "", status)
			json.NewEncoder(w).Encode(errorMessage{Error: message})
		}))

		next.ServeHTTP(w, r)
	})
}

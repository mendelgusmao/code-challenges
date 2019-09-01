package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/database"
)

func Database(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := database.Instance()

		if db == nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		context.Set(r, "db", db)
		next.ServeHTTP(w, r)
	})
}

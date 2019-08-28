package middleware

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/context"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/config"
	"go.etcd.io/bbolt"
)

var (
	db   *bbolt.DB
	once sync.Once
)

func Database(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		once.Do(func() {
			var err error
			db, err = bbolt.Open(config.Backend.Database, 0600, nil)

			if err != nil {
				log.Printf("opening database: %v", err)
			}
		})

		if db == nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		context.Set(r, "db", db)
		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Database(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient(options.Client().ApplyURI(config.Backend.DatabaseURL))

		if err != nil {
			log.Printf("middleware.Database: %s\n", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		db := client.Database(config.Backend.DatabaseName)
		context.Set(r, "db", &db)

		next.ServeHTTP(w, r)
	})
}

package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	gorillaContext "github.com/gorilla/context"
	"github.com/mendelgusmao/supereasy/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Database(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := mongo.NewClient(options.Client().ApplyURI(config.Backend.DatabaseURL))

		if err != nil {
			log.Printf("middleware.Database: %s\n", err)
			http.Error(w, "", http.StatusInternalServerError)

			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		db := client.Database(config.Backend.DatabaseName)

		if err := client.Connect(ctx); err != nil {
			log.Printf("middleware.Database: %s\n", err)
			http.Error(w, "", http.StatusInternalServerError)

			return
		}

		ctxPing, cancelPing := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancelPing()

		if err := client.Ping(ctxPing, readpref.Primary()); err != nil {
			log.Printf("middleware.Database: %s\n", err)
			http.Error(w, "", http.StatusInternalServerError)

			return
		}

		gorillaContext.Set(r, "db", db)
		next.ServeHTTP(w, r)
		client.Disconnect(ctx)
	})
}

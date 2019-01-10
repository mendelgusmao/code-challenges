package middleware

import (
	"database/sql"
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/backend/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
)

func init() {
	config.AfterLoad(checkDatabaseConnection)
}

func Database(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", config.Backend.Database)

		if err != nil {
			log.Printf("middleware.Database: %s\n", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		context.Set(r, "db", db)

		next.ServeHTTP(w, r)
	})
}

func checkDatabaseConnection(c *config.Specification) error {
	log.Println("checking database connection")
	db, err := sql.Open("mysql", config.Backend.Database)
	defer db.Close()

	if err == nil {
		log.Println("successful database connection")
	}

	return err
}

package middleware

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/config"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

const sessionName = "session_id"

var store sessions.Store

func init() {
	config.AfterLoad(createSessionStore)
}

func createSessionStore(c *config.Specification) error {
	log.Println("creating session store with specified session key")
	store = sessions.NewCookieStore([]byte(c.SessionKey))

	return nil
}

func Session(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := store.Get(r, sessionName)

		if err != nil {
			log.Printf("middleware.Session: %s", err)
			http.Error(w, "", http.StatusInternalServerError)
		} else {
			context.Set(r, "session", session)
		}

		next.ServeHTTP(w, r)
	})
}

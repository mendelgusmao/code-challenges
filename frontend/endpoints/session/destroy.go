package session

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/middleware"
	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

func init() {
	subrouter := router.Router.PathPrefix("/session/destroy").Subrouter()
	subrouter.Use(middleware.RequireSession)

	subrouter.HandleFunc("", showDestroySession).Methods("GET")
	subrouter.HandleFunc("", destroySession).Methods("POST")
}

func showDestroySession(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "session/destroy", nil)
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	session.Values["valid"] = false

	if err := session.Save(r, w); err != nil {
		log.Printf("sessions.destroySession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	session.AddFlash("Successfully logged out!")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

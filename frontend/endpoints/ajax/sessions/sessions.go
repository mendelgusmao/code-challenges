package sessions

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"github.com/gorilla/sessions"
)

const sessionName = "session_id"

func init() {
	subrouter := router.AJAX.PathPrefix("/sessions").Subrouter()

	subrouter.HandleFunc("", getSession).Methods("GET")
	subrouter.HandleFunc("", createSession).Methods("POST")
	subrouter.HandleFunc("", destroySession).Methods("DELETE")
}

func getSession(w http.ResponseWriter, r *http.Request) {
	session := retrieveSession(w, r)

	if session == nil {
		return
	}

	if valid, ok := session.Values["valid"]; !ok || (ok && !valid.(bool)) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	session := retrieveSession(w, r)

	if session == nil {
		return
	}

	session.Values["valid"] = true

	if err := session.Save(r, w); err != nil {
		log.Printf("sessions.createSession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	session := retrieveSession(w, r)

	if session == nil {
		return
	}

	session.Values["valid"] = false

	if err := session.Save(r, w); err != nil {
		log.Printf("sessions.destroySession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}
}

func retrieveSession(w http.ResponseWriter, r *http.Request) *sessions.Session {
	session, err := store.Get(r, sessionName)

	if err != nil {
		log.Printf("sessions.retrieveSession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return nil
	}

	return session
}

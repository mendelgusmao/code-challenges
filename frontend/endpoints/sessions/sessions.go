package sessions

import (
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
)

func init() {
	subrouter := router.Router.PathPrefix("/sessions").Subrouter()

	subrouter.HandleFunc("", getSession).Methods("GET")
	subrouter.HandleFunc("", createSession).Methods("POST")
	subrouter.HandleFunc("", destroySession).Methods("DELETE")
}

func getSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func createSession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

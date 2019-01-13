package sessions

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func init() {
	subrouter := router.AJAX.PathPrefix("/sessions").Subrouter()

	subrouter.HandleFunc("", getSession).Methods("GET")
	subrouter.HandleFunc("", createSession).Methods("POST")
	subrouter.HandleFunc("", destroySession).Methods("DELETE")
}

func getSession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	if valid, ok := session.Values["valid"]; !ok || (ok && !valid.(bool)) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	credentials := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	response, err := resty.R().
		SetBody(credentials).
		SetResult(map[string]interface{}{}).
		Post("/users/authenticate")

	if err != nil {
		log.Printf("sessions.createSession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if response.StatusCode() == http.StatusOK {
		session.Values["valid"] = true
		userData := response.Result().(*map[string]interface{})

		for key, value := range *userData {
			session.Values[key] = value
		}

		if sessionErr := session.Save(r, w); sessionErr != nil {
			log.Printf("sessions.createSession: %s", sessionErr)
			http.Error(w, "", http.StatusInternalServerError)
		}
	} else {
		log.Printf("sessions.createSession: backend returned %d", response.StatusCode())
		http.Error(w, "", response.StatusCode())
	}
}

func destroySession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	session.Values["valid"] = false

	if err := session.Save(r, w); err != nil {
		log.Printf("sessions.destroySession: %s", err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

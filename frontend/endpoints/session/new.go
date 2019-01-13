package session

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func init() {
	subrouter := router.Router.PathPrefix("/session").Subrouter()

	subrouter.HandleFunc("/new", newSession).Methods("GET")
	subrouter.HandleFunc("/new", createSession).Methods("POST")
}

func newSession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	valid, ok := session.Values["valid"]

	if ok && valid.(bool) {
		session.AddFlash("Already logged in!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templates.Render(w, r, "session/new", nil)
}

func createSession(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)

	valid, ok := session.Values["valid"]

	if ok && valid.(bool) {
		session.AddFlash("Already logged in!")
		http.Redirect(w, r, "/profile", http.StatusSeeOther)
		return
	}

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

	switch response.StatusCode() {
	case http.StatusOK:
		session.Values["valid"] = true
		userData := response.Result().(*map[string]interface{})

		for key, value := range *userData {
			session.Values[key] = value
		}

		if sessionErr := session.Save(r, w); sessionErr != nil {
			log.Printf("sessions.createSession: %s", sessionErr)
			session.AddFlash("Error logging in!")
			http.Redirect(w, r, "/session/new", http.StatusSeeOther)
			return
		}

		session.AddFlash("Successfully logged in!")
		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.StatusForbidden:
		session.AddFlash("Couldn't log in. Check your credentials.")
		http.Redirect(w, r, "/session/new", http.StatusSeeOther)

	default:
		log.Printf("sessions.createSession: backend returned %d", response.StatusCode())
		session.AddFlash("Error logging in!")
		http.Redirect(w, r, "/session/new", http.StatusSeeOther)
	}
}

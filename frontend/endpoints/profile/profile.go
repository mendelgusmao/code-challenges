package profile

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
	subrouter := router.Router.PathPrefix("/profile").Subrouter()
	subrouter.Use(middleware.RequireSession)

	subrouter.HandleFunc("", profile).Methods("GET")
}

func profile(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	id := session.Values["id"].(int)

	user, err := retrieveUser(id)

	if err != nil {
		log.Printf("profile: %s", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	templateData := struct {
		User map[string]interface{}
	}{
		User: *user,
	}

	templates.Render(w, r, "profile/profile", templateData)
}

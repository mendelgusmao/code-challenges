package profile

import (
	"fmt"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/middleware"
	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	resty "gopkg.in/resty.v1"
)

func init() {
	subrouter := router.Router.PathPrefix("/profile").Subrouter()
	subrouter.Use(middleware.RequireSession)

	subrouter.HandleFunc("/edit", edit).Methods("GET")
	subrouter.HandleFunc("/edit", update).Methods("POST")
}

func edit(w http.ResponseWriter, r *http.Request) {
	templates.
		NewRenderer("profile/edit", "profile/form").
		Do(w, r, buildTemplateData(w, r, "/profile/edit"))
}

func update(w http.ResponseWriter, r *http.Request) {
	session := context.Get(r, "session").(*sessions.Session)
	id := int(session.Values["id"].(float64))
	uri := fmt.Sprintf("/users/%d", id)

	userData := struct {
		Email     string `json:"email"`
		FullName  string `json:"full_name"`
		Telephone string `json:"telephone"`
	}{
		Email:     r.FormValue("email"),
		FullName:  r.FormValue("full_name"),
		Telephone: r.FormValue("telephone"),
	}

	response, err := resty.R().
		SetBody(userData).
		Patch(uri)

	if err != nil {
		session.AddFlash("There was an error updating your profile.")
	} else {
		if response.StatusCode() == http.StatusNoContent {
			session.AddFlash("Profile updated successfully.")
		} else {
			session.AddFlash("There was an error updating your profile.")
		}
	}

	session.Save(r, w)

	templates.
		NewRenderer("profile/edit", "profile/form").
		Do(w, r, buildTemplateData(w, r, "/profile/edit"))
}

package profile

import (
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/middleware"
	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
)

func init() {
	subrouter := router.Router.PathPrefix("/profile").Subrouter()
	subrouter.Use(middleware.RequireSession)

	subrouter.HandleFunc("/edit", edit).Methods("GET")
}

func edit(w http.ResponseWriter, r *http.Request) {
	templateData := struct {
		User map[string]string
	}{
		User: map[string]string{
			"full_name": "John Doe",
			"email":     "john.doe@mail.org",
			"telephone": "+1 (949) 555-1234",
		},
	}

	templates.Render(w, r, "profile/edit", templateData)
}

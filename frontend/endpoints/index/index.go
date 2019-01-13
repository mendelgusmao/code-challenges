package index

import (
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
	"bitbucket.org/mendelgusmao/me_gu/frontend/templates"
)

func init() {
	subrouter := router.Router.PathPrefix("/").Subrouter()

	subrouter.HandleFunc("/", index).Methods("GET")
}

func index(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, r, "index", nil)
}

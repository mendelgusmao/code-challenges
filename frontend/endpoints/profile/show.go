package profile

import (
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/middleware"
	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
)

func init() {
	subrouter := router.Router.PathPrefix("/profile").Subrouter()
	subrouter.Use(middleware.RequireSession)

	subrouter.HandleFunc("", show).Methods("GET")
}

func show(w http.ResponseWriter, r *http.Request) {
	handler(w, r, "profile/show")
}

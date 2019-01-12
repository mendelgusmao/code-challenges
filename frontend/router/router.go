package router

import "github.com/gorilla/mux"

var (
	Router = mux.NewRouter()
	AJAX   = Router.PathPrefix("/ajax").Subrouter()
)

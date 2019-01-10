package router

import (
	"bitbucket.org/mendelgusmao/me_gu/backend/middleware"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.Database)
}

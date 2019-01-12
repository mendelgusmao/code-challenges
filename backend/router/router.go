package router

import (
	"bitbucket.org/mendelgusmao/me_gu/backend/middleware"
	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.ContentType("application/json; charset=utf-8"))
	Router.Use(middleware.Database)
	Router.Use(middleware.Logging)
}

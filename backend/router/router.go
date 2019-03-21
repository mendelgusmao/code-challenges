package router

import (
	"github.com/gorilla/mux"
	"github.com/mendelgusmao/supereasy/backend/middleware"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.ContentType("application/json; charset=utf-8"))
	Router.Use(middleware.Database)
	Router.Use(middleware.Logging)
	Router.Use(middleware.JSONDecoder)
	Router.Use(middleware.Error)
}

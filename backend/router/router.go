package router

import (
	"github.com/gorilla/mux"
	"github.com/mendelgusmao/zap-challenge/backend/middleware"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.ContentType("application/json; charset=utf-8"))
	Router.Use(middleware.Logging)
}

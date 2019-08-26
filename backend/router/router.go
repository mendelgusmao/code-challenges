package router

import (
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/middleware"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(middleware.ContentType("application/json; charset=utf-8"))
}

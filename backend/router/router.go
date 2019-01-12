package router

import (
	"bitbucket.org/mendelgusmao/me_gu/backend/middleware"
	common "bitbucket.org/mendelgusmao/me_gu/common/middleware"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.Use(common.ContentType("application/json; charset=utf-8"))
	Router.Use(middleware.Database)
	Router.Use(common.Logging)
}

package router

import (
	common "bitbucket.org/mendelgusmao/me_gu/common/middleware"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
	AJAX   = Router.PathPrefix("/ajax").Subrouter()
)

func init() {
	Router.Use(common.Logging)
	AJAX.Use(common.ContentType("application/json; charset=utf-8"))
}

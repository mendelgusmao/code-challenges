package router

import (
	common "bitbucket.org/mendelgusmao/me_gu/common/middleware"
	"bitbucket.org/mendelgusmao/me_gu/frontend/middleware"
	"github.com/gorilla/mux"
)

var (
	Router = mux.NewRouter()
)

func init() {
	Router.Use(common.Logging)
	Router.Use(middleware.Session)
}

package partners

import (
	"github.com/gorilla/mux"
	"github.com/mendelgusmao/supereasy/backend/middleware"
	"github.com/mendelgusmao/supereasy/backend/router"
)

var subrouter *mux.Router = router.Router.PathPrefix("/partners").Subrouter()

func init() {
	subrouter.Use(middleware.RequireToken)
}

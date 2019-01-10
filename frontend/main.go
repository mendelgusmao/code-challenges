package main

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/frontend/config"
	_ "bitbucket.org/mendelgusmao/me_gu/frontend/endpoints"
	"bitbucket.org/mendelgusmao/me_gu/frontend/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("starting frontend server")
	log.Fatal(http.ListenAndServe(config.Frontend.Address, router.Router))
}

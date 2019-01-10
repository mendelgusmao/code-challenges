package main

import (
	"log"
	"net/http"

	"bitbucket.org/mendelgusmao/me_gu/backend/config"
	_ "bitbucket.org/mendelgusmao/me_gu/backend/endpoints"
	"bitbucket.org/mendelgusmao/me_gu/backend/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("starting backend server at", config.Backend.Address)
	log.Fatal(http.ListenAndServe(config.Backend.Address, router.Router))
}

package main

import (
	"log"
	"net/http"

	"github.com/mendelgusmao/tax-challenge/backend/config"
	_ "github.com/mendelgusmao/tax-challenge/backend/endpoints"
	"github.com/mendelgusmao/tax-challenge/backend/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("starting backend server at", config.Backend.Address)
	log.Fatal(http.ListenAndServe(config.Backend.Address, router.Router))
}

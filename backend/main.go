package main

import (
	"log"
	"net/http"

	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/config"
	_ "gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/endpoints"
	"gitlab.com/cabify-challenge/car-pooling-challenge-mendelgusmao/backend/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("starting backend server at", config.Backend.Address)
	log.Fatal(http.ListenAndServe(config.Backend.Address, router.Router))
}

package main

import (
	"log"
	"net/http"

	"github.com/mendelgusmao/zap-challenge/backend/config"
	_ "github.com/mendelgusmao/zap-challenge/backend/endpoints"
	"github.com/mendelgusmao/zap-challenge/backend/router"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatal(err)
	}

	log.Println("starting backend server at", config.Backend.Address)
	log.Fatal(http.ListenAndServe(config.Backend.Address, router.Router))
}

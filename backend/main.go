package main

import (
	"log"
	"net/http"

	_ "bitbucket.org/mendelgusmao/me_gu/backend/endpoints"
	"bitbucket.org/mendelgusmao/me_gu/backend/router"
)

func main() {
	log.Fatal(
		http.ListenAndServe(":8000", router.Router),
	)
}

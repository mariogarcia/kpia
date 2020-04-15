package main

import (
	"github.com/mariogarcia/kpia/api"
	"log"
	"net/http"
)

func main() {
	service := api.New()
	log.Fatal(http.ListenAndServe(":8000", service.Router()))
}

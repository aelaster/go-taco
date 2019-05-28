package main

import (
	"log"
	"net/http"

	"github.com/aelaster/go-taco/api"
)

func main() {
	router := api.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"
	"os"

	"github.com/dowdeswells/testapi/repository"
	"github.com/dowdeswells/testapi/router"
)

func main() {
	address := getAddress()
	r := router.NewRouter(repository.NewRepository)
	log.Println("Starting Server on Address " + address)
	http.ListenAndServe(address, r)
}

func getAddress() string {
	port := os.Getenv("PORT")
	return ":" + port
}

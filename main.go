package main

import (
    "net/http"
    "github.com/dowdeswells/testapi/router"
    "log"
    "os"
)

func main() {
    address := getAddress()
    r := router.NewRouter()
    log.Println("Starting Server on Address " + address)
    http.ListenAndServe(address, r)
}


func getAddress() string {
    port := os.Getenv("PORT");
    return ":" + port;
}

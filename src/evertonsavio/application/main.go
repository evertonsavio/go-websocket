package main

import (
	"github.com/havyx/golang-websocket/src/evertonsavio/application/services"
	"log"
	"net/http"
)

func main()  {
	
	mux := routes()

	log.Println("Starting channel listener")
	go services.ListenToWsChannel()

	log.Println("Starting on port 8080")

	_ = http.ListenAndServe(":8080", mux)


}
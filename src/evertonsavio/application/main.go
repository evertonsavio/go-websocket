package main

import (
	"github.com/havyx/golang-websocket/src/evertonsavio/application/services"
	"log"
	"net/http"
)

func main()  {

	ticket := "eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJmNDgxZjc4ZS1kYzkyLTRlZTctOTgxNC1jNWI5OTE3MDNmMzAifQ.VqAUqSbuw0sinyQt8dRCx1IUX9I1KjgSv6h14YUhCYuUORABSOCvq0h4IZL4aB9jwUBY0XxiPIoO4U3VgYHw-Q"
	ExtractClaims(ticket)
	
	mux := routes()

	log.Println("Starting channel listener")
	go services.ListenToWsChannel()

	log.Println("Starting on port 8080")

	_ = http.ListenAndServe(":8080", mux)


}
package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/havyx/golang-websocket/src/evertonsavio/application/handlers"
	"github.com/havyx/golang-websocket/src/evertonsavio/application/services"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(services.WsEndpoint))

	return mux
}

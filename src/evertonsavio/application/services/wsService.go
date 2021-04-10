package services

import (
	"log"
	"net/http"

	"github.com/havyx/golang-websocket/src/evertonsavio/application/models"
	"github.com/gorilla/websocket"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}


/*FUNC WS ENDPOINT UPGRADE CONNECTION
=======================================================================================*/
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil { log.Println(err) }

	log.Println("CLIENT CONNECTED TO ENDPOINT")

	var response models.WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(response)

	if err != nil { log.Println(err) }
}

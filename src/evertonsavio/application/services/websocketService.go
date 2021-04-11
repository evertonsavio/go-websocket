package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/havyx/golang-websocket/src/evertonsavio/application/types"
	"github.com/gorilla/websocket"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var wsChan = make(chan types.WsJsonPayload)

var clients = make(map[types.WebSocketConnection]string)

/*FUNC WEBSOCKET ENDPOINT UPGRADE CONNECTION
=================================================================================================*/
func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil { log.Println(err) }

	log.Println("CLIENT CONNECTED TO ENDPOINT")

	var response types.WsJsonResponse
	response.Message = `Connected to server`

	conn := types.WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)

	if err != nil { log.Println(err) }

	go ListenForWs(&conn)
}

/*GO ROUTINE LISTEN TO WEBSOCKET CONNECTIONS
=================================================================================================*/
func ListenForWs(conn *types.WebSocketConnection){
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload types.WsJsonPayload

	for{
		err := conn.ReadJSON(&payload)
		if err != nil {
			//do nothing
		}else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

/*LISTEN TO CHANNEL PAYLOAD AND BROADCAST
=================================================================================================*/
func ListenToWsChannel(){
	var response types.WsJsonResponse

	for{
		e := <-wsChan

		response.Action = "Got here"
		response.Message = fmt.Sprintf("Some message %s", e.Action)
		broadcastToAll(response)
	}
}

/*BROADCAST TO ALL OR REMOVE CLIENT CONNECTION
=================================================================================================*/
func broadcastToAll(response types.WsJsonResponse){
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}
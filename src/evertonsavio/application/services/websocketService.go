package services

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gorilla/websocket"
	"github.com/havyx/golang-websocket/src/evertonsavio/application/structs"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var wsChan = make(chan structs.WsJsonPayload)

var clients = make(map[structs.WebSocketConnection]string)

/*FUNC WEBSOCKET ENDPOINT UPGRADE CONNECTION
=================================================================================================*/
func WsEndpoint(w http.ResponseWriter, r *http.Request) {

	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("CLIENT CONNECTED TO ENDPOINT")

	var response structs.WsJsonResponse
	response.Message = `Connected to server`

	conn := structs.WebSocketConnection{Conn: ws}
	clients[conn] = ""

	err = ws.WriteJSON(response)

	if err != nil {
		log.Println(err)
	}

	go ListenForWs(&conn)
}

/*GO ROUTINE LISTEN TO WEBSOCKET CONNECTIONS
=================================================================================================*/
func ListenForWs(conn *structs.WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload structs.WsJsonPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			//do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

/*LISTEN TO CHANNEL PAYLOAD AND BROADCAST
=================================================================================================*/
func ListenToWsChannel() {
	var response structs.WsJsonResponse

	for {
		e := <-wsChan

		switch e.Action {
		case "username":
			//get a list of all users and broadcast
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "List_users"
			response.Message = fmt.Sprintf("%s: %s", e.Username, e.Message)
			response.ConnectedUsers = users
			broadcastToAll(response)

		case "left":
			response.Action = "List_users"
			delete(clients, e.Conn)
			users := getUserList()
			response.ConnectedUsers = users
			broadcastToAll(response)
		case "broadcast":
			response.Action = "Got here"
			response.Message = fmt.Sprintf("%s: %s", e.Username, e.Message)
			broadcastToAll(response)
		}
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
}

/*BROADCAST TO ALL OR REMOVE CLIENT CONNECTION
=================================================================================================*/
func broadcastToAll(response structs.WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

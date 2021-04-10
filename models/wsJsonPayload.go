package models

import (
	"github.com/gorilla/websocket"
)

type webSocketConnection struct {
	*websocket.Conn
}

type WsJsonPayload struct {
	Action string `json:"action"`
	Username string `json:"username"`
	Message string `json:"message"`
	Conn webSocketConnection `json:"-"`
}
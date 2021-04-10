package handlers

import (
	"log"
	"net/http"

	"github.com/havyx/golang-websocket/models"
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

/*FUNC - HOME HANDLER: Render home page
=======================================================================================*/
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Panicln(err)
	}
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

func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

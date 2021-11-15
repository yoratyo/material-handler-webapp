package shared

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// We set our Read and Write buffer sizes
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// The Upgrade function will take in an incoming request and upgrade the request
// into a websocket connection
func NewWebsocketConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// this line allows other origin hosts to connect to our
	// websocket server
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// creates our websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	// returns our new websocket connection
	return ws, nil
}

package client

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin:      origin,
		HandshakeTimeout: 1 * time.Second,
	}
)

func origin(r *http.Request) bool {
	return true
}

//NewWebSocketUpgrader - return new upgrade for websocket connect
func NewWebSocketUpgrader(w http.ResponseWriter, r *http.Request) (res *websocket.Conn, err error) {
	res, err = upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	return res, err
}

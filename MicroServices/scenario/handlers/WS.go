package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
)

//Ws - handler for websocket, send status of scenario to client
func Ws(w http.ResponseWriter, r *http.Request) {
	var res bool
	c, err := client.NewWebSocketUpgrader(w, r)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Print("upgrade:", err)
	}
	check := jwttoken.Parse(string(message))
	if check {
		for {
			res = scn.СheckRun()
			if res {
				err := websocket.WriteJSON(c, scn.GetState)
				if err != nil {
					log.Println("Error: ", err.Error())
					return
				}
			} else {
				err := websocket.WriteJSON(c, nil)
				if err != nil {
					log.Println("Error: ", err.Error())
					return
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

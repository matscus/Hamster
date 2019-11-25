package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/matscus/Hamster/MicroServices/check/check"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Services/service"
)

//ResultWs - results slice WS for websocket
type ResultWs struct {
	Result []WS `json:"Result"`
}

//WS  - struct for websocket
type WS struct {
	ServiceName string `json:"servicename"`
	Host        string `json:"host"`
	State       bool   `json:"state"`
}

//Ws - handler for websocket, send status of scenario to client
func Ws(w http.ResponseWriter, r *http.Request) {
	c, err := client.NewWebSocketUpgrader(w, r)
	if err != nil {
		log.Printf("[ERROR] Not created new WS upgrader %s", err.Error())
	}
	defer c.Close()
	_, message, err := c.ReadMessage()
	if err != nil {
		log.Printf("[ERROR] Token error %s", err.Error())
	}
	ok := jwttoken.Parse(string(message))
	if ok {
		datachan := make(chan *[]service.Service, 10)
		_, message, err = c.ReadMessage()
		if err != nil {
			log.Printf("[ERROR] Not Read message: %s", err.Error())
		}
		alldata, err := check.InitGetResponseAllData(string(message))
		if err != nil {
			log.Printf("[ERROR] Get service error: %s", err.Error())
		}
		datachan <- alldata
		go func(datachan chan *[]service.Service) {
		outer:
			for {
				_, message, err = c.ReadMessage()
				if err != nil {
					errClose := c.Close()
					if errClose != nil {
						log.Printf("[ERROR] WS Close: %s", err.Error())
					}
					log.Printf("[ERROR] upgrade : %s", err.Error())
					break outer
				}
				alldata, err := check.InitGetResponseAllData(string(message))
				if err != nil {
					log.Print("GetService:", err)
				}
				datachan <- alldata
			}
		}(datachan)
		data := <-datachan
	outer:
		for {
			select {
			case data = <-datachan:
				res, err := check.CheckStend(data)
				err = websocket.WriteJSON(c, res)
				if err != nil {
					errClose := c.Close()
					if errClose != nil {
						log.Printf("[ERROR] WS Close: %s", err.Error())
					}
					log.Printf("[ERROR] Chan : %s", err.Error())
					break outer
				}
			default:
				res, err := check.CheckStend(data)
				err = websocket.WriteJSON(c, res)
				if err != nil {
					errClose := c.Close()
					if errClose != nil {
						log.Printf("[ERROR] WS Close: %s", err.Error())
					}
					log.Printf("[ERROR] WS Default: %s", err.Error())
					break outer
				}
			}
			time.Sleep(1 * time.Second)
		}
	}
}

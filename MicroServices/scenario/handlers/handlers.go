package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
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

//GetData - handle return state all scenario and generators
func GetData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["project"]
	if ok {
		if len(scn.GetResponseAllData.Scenarios) == 0 || len(scn.GetResponseAllData.Generators) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"len GetResponceAllData slice equally 0\"}"))
		} else {
			res := scn.GetResponse{}
			l := len(scn.GetResponseAllData.Scenarios)
			iter := 0
			for i := 0; i < l; i++ {
				projects := scn.GetResponseAllData.Scenarios[i].Projects
				for i := 0; i < len(projects); i++ {
					if projects[i] == page {
						res.Scenarios = append(res.Scenarios, scn.GetResponseAllData.Scenarios[iter])
						break
					}
				}
				iter++
			}
			scn.CheckRunsGen()
			res.Generators = scn.GetResponseAllData.Generators
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\" params hot found \"}"))
	}
}

//UpdateData - handle for update scenario values to table
func UpdateData(w http.ResponseWriter, r *http.Request) {
	var s scenario.Scenario
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err = s.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"update done\"}"))
			scn.InitData()
		}
	}

}

//NewScenario - handle to insert new scenario to table
func NewScenario(w http.ResponseWriter, r *http.Request) {
	var s scenario.Scenario
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err = s.InsertToDB()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"update done\"}"))
			scn.InitData()
		}
	}
}

//StartScenario - handle to start scenario
func StartScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StartRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	ok := s.CheckParams()
	if ok {
		runsgen, err := scn.CheckGen(s.Generators)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			err := json.NewEncoder(w).Encode(runsgen)
			if err != nil {
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		} else {
			if len(runsgen) == 0 {
				scn.LastRunsParams.Store(s.Name, s)
				scn.RunsGenerators.Store(s.Name, s)
				err = s.Start()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Message\":\"the run\"}"))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				err := json.NewEncoder(w).Encode(runsgen)
				if err != nil {
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				}
			}
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"Startup options cannot be empty or equal to 0\"}"))
	}
}

//StopScenario - handle to stop scenario
func StopScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StopRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	s.Stop()
}

//GetLastParams - init slace for response last scenario params
func GetLastParams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["name"]
	if ok {
		res, ок := scn.LastRunsParams.Load(page)
		if ок {
			params := res.(scn.StartRequest)
			err := json.NewEncoder(w).Encode(params)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		} else {
			w.WriteHeader(http.StatusOK)
			nilres := scn.StartRequest{}
			err := json.NewEncoder(w).Encode(nilres)
			if err != nil {
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\" params hot found \"}"))
	}
}

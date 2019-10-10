package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
)

func MiddlewareFiles(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Keep-Alive", "300")
		w.Header().Add("Content-Disposition", "attachment")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Methods", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		w.Header().Add("Access-Control-Max-Age", "600")
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Strict-Transport-Security", "max-age=31536000")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		h.ServeHTTP(w, r)
	}
}

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
	switch r.Method {
	case "PUT":
		var s scenario.Scenario
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			s.ID = id
			s.Name = r.FormValue("scenarioName")
			s.Type = r.FormValue("scenarioType")
			s.Gun = r.FormValue("gun")
			s.Projects = []string{r.FormValue("project")}
			ifExist, _ := s.CheckScenario()
			if ifExist {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"dublicate scenario name in the project\"}"))
			} else {
				oldname, err := s.GetNameForID()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					os.Rename(os.Getenv("DIRPROJECTS")+"/"+s.Projects[0]+"/"+s.Gun+"/"+oldname+".zip", os.Getenv("DIRPROJECTS")+"/"+s.Projects[0]+"/"+s.Gun+"/"+s.Name+".zip")
					r.ParseMultipartForm(32 << 20)
					file, _, _ := r.FormFile("uploadFile")
					if file == nil {
						s.Update()
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("{\"Message\":\"update done\"}"))
						scn.InitData()
					} else {
						defer file.Close()
						f, err := os.OpenFile(os.Getenv("DIRPROJECTS")+"/"+s.Projects[0]+"/"+s.Gun+"/"+s.Name+".zip", os.O_CREATE|os.O_RDWR, os.FileMode(0755))
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
							w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
						}
						defer f.Close()
						_, err = io.Copy(f, file)
						if err != nil {
							err = s.Update()
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
				}
			}
		}
	case "DELETE":
		var s scenario.Scenario
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			s.ID = id
			s.Name = r.FormValue("scenarioName")
			s.Type = r.FormValue("scenarioType")
			s.Gun = r.FormValue("gun")
			s.Projects = []string{r.FormValue("project")}
			os.Remove(os.Getenv("DIRPROJECTS") + "/" + s.Projects[0] + "/" + s.Gun + "/" + s.Name + ".zip")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				s.DeleteScenario()
				scn.InitData()
			}
		}
	}
}

//NewScenario - handle to insert new scenario to table
func NewScenario(w http.ResponseWriter, r *http.Request) {
	var s scenario.Scenario
	s.Name = r.FormValue("scenarioName")
	s.Type = r.FormValue("scenarioType")
	s.Gun = r.FormValue("gun")
	s.Projects = []string{r.FormValue("project")}
	r.ParseMultipartForm(32 << 20)
	ifExist, _ := s.CheckScenario()
	if ifExist {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"dublicate scenario name in the project\"}"))
	} else {
		file, _, err := r.FormFile("uploadFile")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			f, err := os.OpenFile(os.Getenv("DIRPROJECTS")+"/"+s.Projects[0]+"/"+s.Gun+"/"+s.Name+".zip", os.O_CREATE|os.O_RDWR, os.FileMode(0755))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
			defer f.Close()
			_, err = io.Copy(f, file)
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

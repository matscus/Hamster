package handlers

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/JMXParser/jmxparser"
	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
	"github.com/matscus/Hamster/Package/ScriptCache/scriptcache"
)

var (
	cache = scriptcache.New(1*time.Minute, 5*time.Minute)
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
		//|| len(scn.GetResponseAllData.Generators) == 0
		if len(scn.GetResponseAllData.Scenarios) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"len GetResponceAllData slice equally 0\"}"))
		} else {
			res := scn.GetResponse{}
			l := len(scn.GetResponseAllData.Scenarios)
			for i := 0; i < l; i++ {
				if scn.GetResponseAllData.Scenarios[i].Projects == page {
					res.Scenarios = append(res.Scenarios, scn.GetResponseAllData.Scenarios[i])
				}
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
			s.Projects = r.FormValue("project")
			oldname, err := s.GetNameForID()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				os.Rename(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+oldname+".zip", os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip")
				r.ParseMultipartForm(32 << 20)
				file, _, _ := r.FormFile("uploadFile")
				if file == nil {
					s.Update()
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Message\":\"update done\"}"))
					scn.InitData()
				} else {
					defer file.Close()
					f, err := os.OpenFile(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip", os.O_CREATE|os.O_RDWR, os.FileMode(0755))
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
			s.Projects = r.FormValue("project")
			os.Remove(os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip")
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
	s.Projects = r.FormValue("project")
	r.ParseMultipartForm(32 << 20)
	ifExist, _ := s.CheckScenario()
	if ifExist {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"dublicate scenario name in the project\"}"))
	} else {
		file, header, err := r.FormFile("uploadFile")
		defer file.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			authHeader := r.Header.Get("Authorization")
			splitToken := strings.Split(authHeader, "Bearer ")
			authHeader = strings.TrimSpace(splitToken[1])
			fileName := authHeader[0:19] + header.Filename
			cacheScripts, ok := cache.Get(fileName)
			if ok {
				scripts := cacheScripts.(scn.ScriptCache)
				newFile := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip"
				err := ioutil.WriteFile(newFile, scripts.ScriptFile, os.FileMode(0755))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					l := len(scripts.ParseParams)
					for i := 0; i < l; i++ {
						var tg scenario.ThreadGroup
						tg.ThreadGroupName = scripts.ParseParams[i].ThreadGroupName
						tg.ThreadGroupType = scripts.ParseParams[i].ThreadGroupType
						for _, v := range scripts.ParseParams[i].ThreadGroupParams {
							params := scenario.ThreadGroupParams{Type: v.Type, Name: v.Name, Value: v.Value}
							tg.ThreadGroupParams = append(tg.ThreadGroupParams, params)
						}
						s.ThreadGroups = append(s.ThreadGroups, tg)
					}
					err = s.InsertToDB()
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
					} else {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("{\"Message\":\"Create done\"}"))
						scn.InitData()
					}

				}
			} else {
				newFile := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip"
				f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
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
					tempDir := os.Getenv("DIRPROJECTS") + "/temp/"
					os.Mkdir(tempDir, os.FileMode(0755))
					cmd := exec.Command("unzip", newFile, "-d", tempDir)
					cmd.Run()
					filesInfo, _ := ioutil.ReadDir(tempDir)
					fileIfNotExist := true
					for i := 0; i < len(filesInfo); i++ {
						name := filesInfo[i].Name()
						if strings.Contains(name, ".jmx") {
							fileIfNotExist = false
							os.Mkdir(tempDir, os.FileMode(0755))
							file, err := os.Open(tempDir + name)
							defer file.Close()
							byteValue, _ := ioutil.ReadAll(file)
							var testplan jmxparser.JmeterTestPlan
							xml.Unmarshal(byteValue, &testplan)
							tgParams, err := testplan.GetTreadGroupsParams(byteValue)
							if err != nil {
								err = os.RemoveAll(tempDir)
								w.WriteHeader(http.StatusInternalServerError)
								w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
							} else {
								l := len(tgParams)
								for i := 0; i < l; i++ {
									var tg scenario.ThreadGroup
									tg.ThreadGroupName = tgParams[i].ThreadGroupName
									tg.ThreadGroupType = tgParams[i].ThreadGroupType
									for _, v := range tgParams[i].ThreadGroupParams {
										params := scenario.ThreadGroupParams{Type: v.Type, Name: v.Name, Value: v.Value}
										tg.ThreadGroupParams = append(tg.ThreadGroupParams, params)
									}
									s.ThreadGroups = append(s.ThreadGroups, tg)
								}
								err = os.RemoveAll(tempDir)
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
										w.Write([]byte("{\"Message\":\"Create done\"}"))
										scn.InitData()
										break
									}
								}
							}
						}
					}
					if fileIfNotExist {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("{\"Message\": not found jmx file}"))
					}
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
	runsgen, err := scn.CheckGen(s.Generators)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(runsgen)
		if err != nil {
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	} else {
		if len(runsgen) == 0 {
			scn.LastRunsParams.Store(s.Name+s.Projects, s)
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
		lastParamName, ok := params["project"]
		if ok {
			res, ок := scn.LastRunsParams.Load(page + lastParamName)
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
		}

	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\" params hot found \"}"))
	}
}

//PreCheckScenario - handle to rpe check scenario file, if mandatory thread groups params is nil, return fasle.
func PreCheckScenario(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadFile")
	bytesFile := make([]byte, 0, 0)
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		file.Read(bytesFile)
		authHeader := r.Header.Get("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")
		authHeader = strings.TrimSpace(splitToken[1])
		tempParseDir := os.Getenv("DIRPROJECTS") + "/tempParseDir/"
		err = os.Mkdir(tempParseDir, os.FileMode(0755))
		if err != nil {
			if os.IsExist(err) {
				tempParseDir = tempParseDir + authHeader[0:19] + "/"
				os.Mkdir(tempParseDir, os.FileMode(0755))
				fileName := authHeader[0:19] + header.Filename
				newFile := tempParseDir + fileName
				f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
				if err != nil {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				}
				defer f.Close()
				_, err = io.Copy(f, file)
				if err != nil {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				}
				cmd := exec.Command("unzip", newFile, "-d", tempParseDir)
				cmd.Run()
				preparseResponce := make([]scn.PreParseResponce, 0, 0)
				filesInfo, _ := ioutil.ReadDir(tempParseDir)
				fileIfNotExist := true
				for i := 0; i < len(filesInfo); i++ {
					name := filesInfo[i].Name()
					if strings.Contains(name, ".jmx") {
						fileIfNotExist = false
						tempFile, err := os.Open(tempParseDir + name)
						defer file.Close()
						byteValue, _ := ioutil.ReadAll(tempFile)
						var testplan jmxparser.JmeterTestPlan
						xml.Unmarshal(byteValue, &testplan)
						tgParams, err := testplan.GetTreadGroupsParams(byteValue)
						if err != nil {
							err = os.RemoveAll(tempParseDir)
							w.WriteHeader(http.StatusInternalServerError)
							w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
						} else {
							l := len(tgParams)
							for i := 0; i < l; i++ {
								res := make([]string, 0, 4)
								l := len(tgParams[i].ThreadGroupParams)
								for ii := 0; ii < l; ii++ {
									if tgParams[i].ThreadGroupParams[ii].Value == "" {
										res = append(res, tgParams[i].ThreadGroupParams[ii].Type)
									}
								}
								if len(res) > 0 {
									preparseResponce = append(preparseResponce, scn.PreParseResponce{ThreadGroupName: tgParams[i].ThreadGroupName, FailedParams: res})
								}
							}
							if len(preparseResponce) == 0 {
								cache.Set(fileName, scn.ScriptCache{ScriptFile: bytesFile, ParseParams: tgParams}, 1*time.Minute)
								os.RemoveAll(tempParseDir)
								w.WriteHeader(http.StatusOK)
								w.Write([]byte("{\"Message\":\"script structure complies with the standard \"}"))
							} else {
								os.RemoveAll(tempParseDir)
								w.WriteHeader(http.StatusOK)
								err := json.NewEncoder(w).Encode(preparseResponce)
								if err != nil {
									w.WriteHeader(http.StatusOK)
									w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
								}
							}
						}
					}
				}
				if fileIfNotExist {
					w.WriteHeader(http.StatusNoContent)
					w.Write([]byte("{\"Message\": not found jmx file}"))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		}
	}
}

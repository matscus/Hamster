package handlers

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/JMXParser/jmxparser"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
)

//NewScenario - handle to insert new scenario to table
func NewScenario(w http.ResponseWriter, r *http.Request) {
	var s scenario.Scenario
	s.Name = r.FormValue("scenarioName")
	s.Type = r.FormValue("scenarioType")
	s.Gun = r.FormValue("gun")
	s.Projects = r.FormValue("project")
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Parse Multi part Form error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario Parse Multi part Form error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	ifExist, _ := s.CheckScenario()
	if ifExist {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Dublicate scenario name in the project\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Dublicate scenario name in the projec, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
		return
	}
	file, header, err := r.FormFile("uploadFile")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Upload scenario error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Upload scenario error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
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
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO Write file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario IO Write file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
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
			errRemove := os.Remove(newFile)
			if errRemove != nil {
				_, errWrite := w.Write([]byte("{\"Message\":\" Scenatio insert to DB error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenatio insert to DB error and remove file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			} else {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenatio insert to DB error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenatio insert to DB error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
			return
		}
		err = scn.InitData()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Create scenario done, but  scenario init data error" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Create scenario done, but  scenario init data error and but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
		} else {
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"Create scenario done\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Create scenario done, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
		}
	} else {
		newFile := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip"
		f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario open new file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario open new file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errRemove := os.Remove(newFile)
			if errRemove != nil {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO Copy error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario IO Copy and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			} else {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO Copy error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario IO Copy error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
			return
		}
		tempDir := os.Getenv("DIRPROJECTS") + "/temp/"
		err = os.Mkdir(tempDir, os.FileMode(0755))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create temp dir error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario create temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		cmd := exec.Command("unzip", newFile, "-d", tempDir)
		err = cmd.Run()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errRemove := os.RemoveAll(tempDir)
			if errRemove != nil {
				errRemoveNewFile := os.Remove(newFile)
				if errRemoveNewFile != nil {
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario unzip file error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario unzip file and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				} else {
					_, errWrite := w.Write([]byte("{\"Message\":\" Unzip  " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario unzip file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				}
			} else {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario unzip file error " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario unzip file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
			return
		}
		filesInfo, err := ioutil.ReadDir(tempDir)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errRemove := os.RemoveAll(tempDir)
			if errRemove != nil {
				errRemoveNewFile := os.Remove(newFile)
				if errRemoveNewFile != nil {
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario read temp dir error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario read temp dir and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				} else {
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario read temp dir error:" + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario read temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				}
			} else {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario read temp dir error:" + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario read temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
			return
		}
		fileIfNotExist := true
		for i := 0; i < len(filesInfo); i++ {
			name := filesInfo[i].Name()
			if strings.Contains(name, ".jmx") {
				fileIfNotExist = false
				file, err := os.Open(tempDir + name)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errRemove := os.RemoveAll(tempDir)
					if errRemove != nil {
						errRemoveNewFile := os.Remove(newFile)
						if errRemoveNewFile != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario open jmx file error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario open jmx file and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario open jmx file error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario open jmx file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario open jmx file error: " + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario open jmx file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
				defer file.Close()
				byteValue, err := ioutil.ReadAll(file)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errRemove := os.RemoveAll(tempDir)
					if errRemove != nil {
						errRemoveNewFile := os.Remove(newFile)
						if errRemoveNewFile != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO ReadAll file error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR]Scenario IO ReadAll file and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO ReadAll file error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario IO ReadAll file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO ReadAll file error: " + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario IO ReadAll file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
				var testplan jmxparser.JmeterTestPlan
				err = xml.Unmarshal(byteValue, &testplan)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errRemove := os.RemoveAll(tempDir)
					if errRemove != nil {
						errRemoveNewFile := os.Remove(newFile)
						if errRemoveNewFile != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Unmarshal file error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario Unmarshal file and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Unmarshal file error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario Unmarshal file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Unmarshal file error: " + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario Unmarshal file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
				tgParams, err := testplan.GetTreadGroupsParams(byteValue)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errRemove := os.RemoveAll(tempDir)
					if errRemove != nil {
						errRemoveNewFile := os.Remove(newFile)
						if errRemoveNewFile != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario get tread groups params error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario get tread groups params and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario get tread groups params error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario get tread groups params error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario get tread groups params error: " + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR]Scenario get tread groups params error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
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
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario remove tempdir: " + err.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario remove tempdir and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario remove tempdir error: " + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario remove tempdir and remove file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
				err = s.InsertToDB()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					errRemove := os.RemoveAll(tempDir)
					if errRemove != nil {
						errRemoveNewFile := os.Remove(newFile)
						if errRemoveNewFile != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario insert to DB error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + " and error remove file " + errRemoveNewFile.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario insert to DB and remove temp dir and file errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario insert to DB error:" + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario insert to DB error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					} else {
						_, errWrite := w.Write([]byte("{\"Message\":\"InsertToDB" + err.Error() + "\"}"))
						if errWrite != nil {
							log.Printf("[ERROR]Scenario insert to DB error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
						}
					}
					return
				}
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create complited\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario create complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
				}
				err = scn.InitData()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create complited, but not update init data " + err.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario create complited, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				} else {
					w.WriteHeader(http.StatusOK)
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create complited\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario create complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
					}
				}
				break
			}
			if fileIfNotExist {
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario not found jmx file in zip\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario not found jmx file in zip, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
		}
	}
}

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
	"strings"
	"time"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/JMXParser/jmxparser"
)

//PreCheckScenario - handle to rpe check scenario file, if mandatory thread groups params is nil, return fasle.
func PreCheckScenario(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadFile")
	if err != nil {
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario upload file error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario upload file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	defer file.Close()
	bytesFile := make([]byte, 0, 0)
	_, err = file.Read(bytesFile)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario read file error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario read file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = strings.TrimSpace(splitToken[1])
	tempParseDir := os.Getenv("DIRPROJECTS") + "/tempParseDir/"
	err = os.Mkdir(tempParseDir, os.FileMode(0755))
	if err != nil {
		if os.IsExist(err) {
			tempParseDir = tempParseDir + authHeader[0:19] + "/"
			err = os.Mkdir(tempParseDir, os.FileMode(0755))
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create temp dir error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario create temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			fileName := authHeader[0:19] + header.Filename
			newFile := tempParseDir + fileName
			f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario open file error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario open file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			defer f.Close()
			_, err = io.Copy(f, file)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario IO Copy file error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario IO Copy file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			cmd := exec.Command("unzip", newFile, "-d", tempParseDir)
			err = cmd.Run()
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario unzip file error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario unzip file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			preparseResponce := make([]scn.PreParseResponce, 0, 0)
			filesInfo, err := ioutil.ReadDir(tempParseDir)
			if err != nil {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario read temp dir error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario read temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			fileIfNotExist := true
			for i := 0; i < len(filesInfo); i++ {
				name := filesInfo[i].Name()
				if strings.Contains(name, ".jmx") {
					fileIfNotExist = false
					tempFile, err := os.Open(tempParseDir + name)
					defer file.Close()
					byteValue, _ := ioutil.ReadAll(tempFile)
					var testplan jmxparser.JmeterTestPlan
					err = xml.Unmarshal(byteValue, &testplan)
					if err != nil {
						errRemove := os.RemoveAll(tempParseDir)
						w.WriteHeader(http.StatusInternalServerError)
						if errRemove != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Unmarshal file error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario Unmarshal file and remove temp dir errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Unmarshal file error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario Unmarshal file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
						return
					}
					tgParams, err := testplan.GetTreadGroupsParams(byteValue)
					if err != nil {
						errRemove := os.RemoveAll(tempParseDir)
						w.WriteHeader(http.StatusInternalServerError)
						if errRemove != nil {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario get tread groups params error: " + err.Error() + " and error remove tempDir " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario get tread groups params  and remove temp dir errors, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						} else {
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario get tread groups params error: " + err.Error() + " and error remove file " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario get tread groups params error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
						return
					}
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
						errRemove := os.RemoveAll(tempParseDir)
						if errRemove != nil {
							w.WriteHeader(http.StatusInternalServerError)
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario remove temp dir error: " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario remove temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
							return
						}
						w.WriteHeader(http.StatusOK)
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario structure complies with the standard\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario structure complies with the standard, but Not Writing to ResponseWriter due: %s", errWrite.Error())
						}
						return
					} else {
						errRemove := os.RemoveAll(tempParseDir)
						if errRemove != nil {
							w.WriteHeader(http.StatusInternalServerError)
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario remove temp dir error: " + errRemove.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR] Scenario remove temp dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
							return
						}
						err = json.NewEncoder(w).Encode(preparseResponce)
						if err != nil {
							w.WriteHeader(http.StatusOK)
							_, errWrite := w.Write([]byte("{\"Message\":\"Scenario encode tgParams error: " + err.Error() + "\"}"))
							if errWrite != nil {
								log.Printf("[ERROR]Scenario encode tgParams, but  Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
							}
						}
					}
				}
			}
			if fileIfNotExist {
				w.WriteHeader(http.StatusNoContent)
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario - not found jmx file\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario - not found jmx file, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create dir error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario create dir error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
		}
	}
}

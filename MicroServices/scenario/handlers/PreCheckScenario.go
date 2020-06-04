package handlers

import (
	"encoding/json"
	"encoding/xml"
	"errors"
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
	"github.com/matscus/Hamster/Package/errorImpl"
)

//PreCheckScenario - handle to rpe check scenario file, if mandatory thread groups params is nil, return fasle.
func PreCheckScenario(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("uploadFile")
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Get form uploadFile error", err))
		return
	}
	defer file.Close()
	bytesFile := make([]byte, 0, 0)
	_, err = file.Read(bytesFile)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Read file  error", err))
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
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Mkdir tempPasredir error", err))
				return
			}
			fileName := authHeader[0:19] + header.Filename
			newFile := tempParseDir + fileName
			f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Open file error", err))
				return
			}
			defer f.Close()
			_, err = io.Copy(f, file)
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("IO copy file error", err))
				return
			}
			cmd := exec.Command("unzip", newFile, "-d", tempParseDir)
			err = cmd.Run()
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("unzip file error", err))
				return
			}
			preparseResponce := make([]scn.PreParseResponce, 0, 0)
			filesInfo, err := ioutil.ReadDir(tempParseDir)
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Read tempParseDir error", err))
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
						strError := errorImpl.ScenarioError("Unmarshal testplan error", err)
						errRemove := os.RemoveAll(tempParseDir)
						if errRemove != nil {
							strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
							errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemove))
							return
						} else {
							errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Unmarshal testplan error", err))
							return
						}
					}
					tgParams, err := testplan.GetTreadGroupsParams(byteValue)
					if err != nil {
						strError := errorImpl.ScenarioError("Get treadGroup error", err)
						errRemove := os.RemoveAll(tempParseDir)
						if errRemove != nil {
							strError = errors.New(strError.Error() + ",RemoveAll temp dir error" + err.Error())
							errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Get treadGroup error", errRemove))
							return
						} else {
							errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Get treadGroup error", err))
							return
						}
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
						err := os.RemoveAll(tempParseDir)
						if err != nil {
							errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Remove testParseDir error", err))
							return
						}
						w.WriteHeader(http.StatusOK)
						_, errWrite := w.Write([]byte("{\"Message\":\"Scenario structure complies with the standard\"}"))
						if errWrite != nil {
							log.Printf("[ERROR] Scenario structure complies with the standard, but Not Writing to ResponseWriter due: %s", errWrite.Error())
						}
						return
					} else {
						err := os.RemoveAll(tempParseDir)
						if err != nil {
							errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Remove testParseDir error", err))
							return
						}
						err = json.NewEncoder(w).Encode(preparseResponce)
						if err != nil {
							errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Encode preparse error", err))
							return
						}
					}
				}
			}
			if fileIfNotExist {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("File not exist error", nil))
				return
			}
		}
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Create dir error", err))
		return
	}
}

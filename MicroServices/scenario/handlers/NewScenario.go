package handlers

import (
	"encoding/xml"
	"errors"
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
	"github.com/matscus/Hamster/Package/errorImpl"
)

//NewScenario - handle to insert new scenario to table
func NewScenario(w http.ResponseWriter, r *http.Request) {
	s := scenario.Scenario{
		Name:     r.FormValue("scenarioName"),
		Type:     r.FormValue("scenarioType"),
		Gun:      r.FormValue("gun"),
		Projects: r.FormValue("project"),
		DBClient: PgClient,
	}
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Parse multipartform error", err))
		return
	}
	ifExist, _ := s.CheckScenario()
	if ifExist {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Dublicate name in the project", nil))
		return
	}
	file, header, err := r.FormFile("uploadFile")
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Get form uploadFile error", err))
		return
	}
	defer file.Close()
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
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Write fole error", err))
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
			strError := errorImpl.ScenarioError("Insert to dabasase error", err)
			errRemove := os.Remove(newFile)
			if errRemove != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", Remove fole error", errRemove))
				return
			} else {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Insert to dabasase error", err))
				return
			}
		}
		err = scn.InitData()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Init data error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Create scenario done\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Create scenario done, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
		return
	}
	newFile := os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip"
	f, err := os.OpenFile(newFile, os.O_CREATE|os.O_RDWR, os.FileMode(0755))
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Open file error", err))
		return
	}
	defer f.Close()
	_, err = io.Copy(f, file)
	if err != nil {
		strError := errorImpl.ScenarioError("IO copy error", err)
		errRemove := os.Remove(newFile)
		if errRemove != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", Remove file error", errRemove))
			return
		} else {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("IO copy error", err))
			return
		}
	}
	tempDir := os.Getenv("DIRPROJECTS") + "/temp/"
	err = os.Mkdir(tempDir, os.FileMode(0755))
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Open file error", err))
		return
	}
	cmd := exec.Command("unzip", newFile, "-d", tempDir)
	err = cmd.Run()
	if err != nil {
		strError := errorImpl.ScenarioError("Unzip file error", err)
		errRemove := os.RemoveAll(tempDir)
		if errRemove != nil {
			strError = errors.New(strError.Error() + ", RemoveAll temp dir error" + err.Error())
			errRemoveNewFile := os.Remove(newFile)
			if errRemoveNewFile != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
				return
			} else {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
				return
			}
		} else {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Unzip file error", err))
			return
		}
	}
	filesInfo, err := ioutil.ReadDir(tempDir)
	if err != nil {
		strError := errorImpl.ScenarioError("Read temp dir error", err)
		errRemove := os.RemoveAll(tempDir)
		if errRemove != nil {
			strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
			errRemoveNewFile := os.Remove(newFile)
			if errRemoveNewFile != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
				return
			} else {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
				return
			}
		} else {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Read temp dir error", err))
			return
		}
	}
	fileIfNotExist := true
	for i := 0; i < len(filesInfo); i++ {
		name := filesInfo[i].Name()
		if strings.Contains(name, ".jmx") {
			fileIfNotExist = false
			file, err := os.Open(tempDir + name)
			if err != nil {
				strError := errorImpl.ScenarioError("Open jmx file error", err)
				errRemove := os.RemoveAll(tempDir)
				if errRemove != nil {
					strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
						return
					} else {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
						return
					}
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Open jmx file error", err))
					return
				}
			}
			defer file.Close()
			byteValue, err := ioutil.ReadAll(file)
			if err != nil {
				strError := errorImpl.ScenarioError("Read file error", err)
				errRemove := os.RemoveAll(tempDir)
				if errRemove != nil {
					strError = errorImpl.ScenarioError(strError.Error()+",RemoveAll temp dir error", errRemove)
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
						return
					} else {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
						return
					}
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Read file error", err))
					return
				}
			}
			var testplan jmxparser.JmeterTestPlan
			err = xml.Unmarshal(byteValue, &testplan)
			if err != nil {
				strError := errorImpl.ScenarioError("Unmarshal testplan error", err)
				errRemove := os.RemoveAll(tempDir)
				if errRemove != nil {
					strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
						return
					} else {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
						return
					}
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Unmarshal testplan error", err))
					return
				}
			}
			tgParams, err := testplan.GetTreadGroupsParams(byteValue)
			if err != nil {
				strError := errorImpl.ScenarioError("Get treadgroup error", err)
				errRemove := os.RemoveAll(tempDir)
				if errRemove != nil {
					strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
						return
					} else {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
						return
					}
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Unmarshal testplan error", err))
					return
				}
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
				strError := errorImpl.ScenarioError("Remove test dir error", err)
				errRemoveNewFile := os.Remove(newFile)
				if errRemoveNewFile != nil {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
					return
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Remove test dir error", err))
					return
				}
			}
			err = s.InsertToDB()
			if err != nil {
				strError := errorImpl.ScenarioError("Insert database error", err)
				errRemove := os.RemoveAll(tempDir)
				if errRemove != nil {
					strError = errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove)
					errRemoveNewFile := os.Remove(newFile)
					if errRemoveNewFile != nil {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error(), errRemoveNewFile))
						return
					} else {
						errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError(strError.Error()+", RemoveAll temp dir error", errRemove))
						return
					}
				} else {
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Insert database error", err))
					return
				}
			}
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario create complited\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario create complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
			err = scn.InitData()
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Init data error", err))
				return
			}
			break
		}
		if fileIfNotExist {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Scenario not found jmx file in zip", nil))
			return
		}
	}
}

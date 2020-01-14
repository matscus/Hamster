package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
)

//UpdateOrDeleteScenario - handle for update scenario values to table
func UpdateOrDeleteScenario(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		var s scenario.Scenario
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Parse scenario ID error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Parse scenario ID error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		s.ID = id
		s.Name = r.FormValue("scenarioName")
		s.Type = r.FormValue("scenarioType")
		s.Gun = r.FormValue("gun")
		s.Projects = r.FormValue("project")
		oldname, err := s.GetNameForID()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Get scenario name error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Get scenario name error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = os.Rename(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+oldname+".zip", os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Rename scenario zip file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Rename scenario zip file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Parse Multi part Form error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario Parse Multi part Form error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Upload scenario file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Upload scenario file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		if file == nil {
			err = s.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Update scenario error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			err = scn.InitData()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"Init scenarios data error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Init scenarios data error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario complited\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Update scenario complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
			return
		}
		defer file.Close()
		f, err := os.OpenFile(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip", os.O_CREATE|os.O_RDWR, os.FileMode(0755))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Open scenario file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Open scenario file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"IO Copy scenario file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] IO Copy scenario file error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = s.Update()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Update scenario error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = scn.InitData()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Update done, but init data error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Update done, but init data error and error Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Update done, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		var s scenario.Scenario
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Parse scenario ID error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Parse scenario ID error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		s.ID = id
		s.Name = r.FormValue("scenarioName")
		s.Type = r.FormValue("scenarioType")
		s.Gun = r.FormValue("gun")
		s.Projects = r.FormValue("project")
		err = s.DeleteScenario()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Delete scenario error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Delete scenario error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = os.Remove(os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Delete from DB is complited, but remove scenario file error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Delete from DB is complited, but remove scenario file error and Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		err = scn.InitData()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Delete scenario complited, but but scenario init data error: " + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Delete scenario complited, but but scenario init data error and Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Delete scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Delete scenario complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

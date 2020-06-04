package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//UpdateOrDeleteScenario - handle for update scenario values to table
func UpdateOrDeleteScenario(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Parse form scenaroiID error", err))
			return
		}
		s := scenario.Scenario{
			ID:       id,
			Name:     r.FormValue("scenarioName"),
			Type:     r.FormValue("scenarioType"),
			Gun:      r.FormValue("gun"),
			Projects: r.FormValue("project"),
			DBClient: PgClient,
		}
		oldname, err := s.GetNameForID()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Get name error", err))
			return
		}
		err = os.Rename(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+oldname+".zip", os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip")
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Rename error", err))
			return
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Parse multipartform error", err))
			return
		}
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Get form uploadFiles error", err))
			return
		}
		if file == nil {
			err = s.Update()
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Update error", err))
				return
			}
			err = scn.InitData()
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Init data error", err))
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
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Open file error", err))
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("IO copy file error", err))
			return
		}
		err = s.Update()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Update error", err))
			return
		}
		err = scn.InitData()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Init data error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Update done, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Parse form scenarioID error", err))
			return
		}
		s := scenario.Scenario{
			ID:       id,
			Name:     r.FormValue("scenarioName"),
			Type:     r.FormValue("scenarioType"),
			Gun:      r.FormValue("gun"),
			Projects: r.FormValue("project"),
			DBClient: PgClient,
		}
		err = s.DeleteScenario()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Delete error", err))
			return
		}
		err = os.Remove(os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip")
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Remove dir error", err))
			return
		}
		err = scn.InitData()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Init data error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Delete scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Delete scenario complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

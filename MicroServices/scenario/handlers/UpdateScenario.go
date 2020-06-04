package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/Scenario/scenario"
	"github.com/matscus/Hamster/Package/httperror"
)

//UpdateOrDeleteScenario - handle for update scenario values to table
func UpdateOrDeleteScenario(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
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
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = os.Rename(os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+oldname+".zip", os.Getenv("DIRPROJECTS")+"/"+s.Projects+"/"+s.Gun+"/"+s.Name+".zip")
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = r.ParseMultipartForm(32 << 20)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		file, _, err := r.FormFile("uploadFile")
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		if file == nil {
			err = s.Update()
			if err != nil {
				httperror.WriteError(w, http.StatusInternalServerError, err)
				return
			}
			err = scn.InitData()
			if err != nil {
				httperror.WriteError(w, http.StatusInternalServerError, err)
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
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = s.Update()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = scn.InitData()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Update scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Update done, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		id, err := strconv.ParseInt(r.FormValue("scenarioID"), 10, 64)
		log.Println(id)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
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
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = os.Remove(os.Getenv("DIRPROJECTS") + "/" + s.Projects + "/" + s.Gun + "/" + s.Name + ".zip")
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		err = scn.InitData()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Delete scenario complited\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Delete scenario complited, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

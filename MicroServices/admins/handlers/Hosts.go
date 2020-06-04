package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/Package/Hosts/hosts"
	"github.com/matscus/Hamster/Package/httperror"
)

//GetAllHosts -  handle function, for get all hosts
func GetAllHosts(w http.ResponseWriter, r *http.Request) {
	allhosts, err := pgClient.GetAllHosts()
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(allhosts)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

//GetAllHostsWithProject - return all host with users projects
func GetAllHostsWithProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	project, _ := params["project"]
	allhosts, err := pgClient.GetAllHostsWithProject(project)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = json.NewEncoder(w).Encode(allhosts)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

//Hosts -  handle function, for new,update and delete host
func Hosts(w http.ResponseWriter, r *http.Request) {
	host := hosts.Host{}
	err := json.NewDecoder(r.Body).Decode(&host)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	host.DBClient = pgClient
	switch r.Method {
	case "POST":
		err = host.Create()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Host created \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Host created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "PUT":
		err = host.Update()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Host updated \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Host updated, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		err = host.Delete()
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Host deleted \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Host deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

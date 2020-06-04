package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/matscus/Hamster/Package/Projects/projects"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//GetAllProjects -  return all projects
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	allproject, err := pgClient.GetAllProjects()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Get all projects error", err))
		return
	}
	err = json.NewEncoder(w).Encode(allproject)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Encode get all projects error", err))
		return
	}
}

//Projects -  create new Project, update Project and delete Project
func Projects(w http.ResponseWriter, r *http.Request) {
	project := projects.Project{}
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Decode project error", err))
		return
	}
	project.DBClient = pgClient
	switch r.Method {
	case "POST":
		err = project.Create()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Create  error", err))
			return
		}
		os.MkdirAll(os.Getenv("DIRPROJECTS")+"/"+project.Name+"/jmeter", 0777)
		if os.IsExist(err) {
			//todo
		} else {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Mkdir project error", err))
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Project created \"}"))
		if errWrite != nil {
			log.Printf("[ERROR]User created, but  Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "PUT":
		name, err := project.DBClient.GetProjectName(project.ID)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Get project name error", err))
			return
		}
		os.Rename(os.Getenv("DIRPROJECTS")+"/"+name, os.Getenv("DIRPROJECTS")+"/"+project.Name)
		err = project.Update()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Rename project dir error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Project updated \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Project updated, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		os.Remove(os.Getenv("DIRPROJECTS") + "/" + project.Name)
		err = project.Delete()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Remove project dir error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Project deleted \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Project deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

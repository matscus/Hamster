package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Projects/projects"
)

//GetAllProjects -  return all projects
func GetAllProjects(w http.ResponseWriter, r *http.Request) {
	allproject, err := client.PGClient{}.New().GetAllProjects()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err := json.NewEncoder(w).Encode(allproject)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

//Projects -  create new Project, update Project and delete Project
func Projects(w http.ResponseWriter, r *http.Request) {
	project := projects.Project{}
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		switch r.Method {
		case "POST":
			err = project.Create()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"User created \"}"))
			}
		case "DELETE":
			err = project.Delete()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"Host deleted \"}"))
			}
		}
	}
}

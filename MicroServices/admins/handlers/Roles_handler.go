package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Roles/roles"
)

//GetAllRoles -  return all projects
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	allroles, err := client.PGClient{}.New().GetAllRoles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err := json.NewEncoder(w).Encode(allroles)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

//Roles -  create new Roles, update Roles and delete Roles
func Roles(w http.ResponseWriter, r *http.Request) {
	role := roles.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		switch r.Method {
		case "POST":
			err = role.Create()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"User created \"}"))
			}
		case "PUT":
			err = role.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"Host updated \"}"))
			}
		case "DELETE":
			err = role.Delete()
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

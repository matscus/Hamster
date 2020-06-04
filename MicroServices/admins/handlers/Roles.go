package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Roles/roles"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//GetAllRoles -  return all projects
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	allroles, err := pgClient.GetAllRoles()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Get all roles error", err))
		return
	}
	err = json.NewEncoder(w).Encode(allroles)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Encode get all roles error", err))
		return
	}
}

//Roles -  create new Roles, update Roles and delete Roles
func Roles(w http.ResponseWriter, r *http.Request) {
	role := roles.Role{}
	err := json.NewDecoder(r.Body).Decode(&role)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Decode roles error", err))
		return
	}
	role.DBClient = pgClient
	switch r.Method {
	case "POST":
		err = role.Create()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Create error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Role created \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Role created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "PUT":
		err = role.Update()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Update error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Role updated \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Role updated, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	case "DELETE":
		err = role.Delete()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Delete error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Role deleted \"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Role deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
	}
}

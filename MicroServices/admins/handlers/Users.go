package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Users/users"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//GetAllUsers -  handle function, for get new token
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allusers, err := pgClient.GetAllUsers()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Get all users error", err))
		return
	}
	err = json.NewEncoder(w).Encode(allusers)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Encode get all users error", err))
		return
	}
}

//Users -  create new users, update users and delete users
func Users(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Decode users error", err))
		return
	}
	user.DBClient = pgClient
	switch r.Method {
	case "POST":
		err = user.Create()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Create user error", err))
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"User created \"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Users created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
		}
	case "PUT":
		err := user.Update()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Update user error", err))
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, errWrite := w.Write([]byte("{\"Message\":\"User updated \"}"))
			if errWrite != nil {
				log.Printf("[ERROR] User updated, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
		}
	case "DELETE":
		if user.ID != 1 {
			err := user.Delete()
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Delete user error", err))
				return
			} else {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"User deleted \"}"))
				if errWrite != nil {
					log.Printf("[ERROR] User deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
				}
			}
		} else {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("You cannot remove God", nil))
			return
		}
	}
}

//ChangePassword -  change user password
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Decode user error", err))
		return
	}
	user.DBClient = pgClient
	err = user.ChangePassword()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AdminsError("Change password error", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte("{\"Message\":\"Password changed \"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Password changed, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

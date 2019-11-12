package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Users/users"
)

//GetAllUsers -  handle function, for get new token
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allusers, err := client.PGClient{}.New().GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err := json.NewEncoder(w).Encode(allusers)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

//Users -  create new users, update users and delete users
func Users(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		switch r.Method {
		case "POST":
			err = user.Create()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"User created \"}"))
			}
		case "PUT":
			err := user.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"User updated \"}"))
			}
		case "DELETE":
			if user.ID != 1 {
				err := user.Delete()
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Message\":\"User deleted \"}"))
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"You cannot remove God \"}"))
			}
		}
	}
}

//ChangePassword -  change user password
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err = user.ChangePassword()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"Password changed \"}"))
		}
	}
}

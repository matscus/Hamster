package handlers

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
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
		client := client.PGClient{}.New()
		switch r.Method {
		case "POST":
			ok, err := client.UserNameIfExist(user.User)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				if ok {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\": User name is exist }"))
				} else {
					err = client.NewUser(user.User, user.Password, user.Role, user.Projects)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
					} else {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("{\"Message\":\"User created \"}"))
					}
				}
			}
		case "PUT":
			err := client.UpdateUser(user.ID, user.Password, user.Role, user.Projects)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"User updated \"}"))
			}
		case "DELETE":
			if user.ID != 0 {
				err := client.DeleteUser(user.ID)
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
		client := client.PGClient{}.New()
		h := sha256.New()
		pass, err := b64.StdEncoding.DecodeString(user.Password)
		h.Write([]byte(pass))
		err = client.ChangeUserPassword(user.ID, fmt.Sprintf("%x", h.Sum(nil)))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"Password changed \"}"))
		}
	}
}

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

//NewUser -  create new users
func NewUser(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		client := client.PGClient{}.New()
		ok, err := client.UserNameIfExist(user.User)
		if ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\": User name is exist }"))
		} else {
			w.WriteHeader(http.StatusOK)
			client.NewUser(user.User, user.Password, user.Role, user.Projects)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

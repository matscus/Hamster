package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Users/users"
	"github.com/matscus/Hamster/Package/httperror"
)

//GetToken -  handle function, for get new token
func GetToken(w http.ResponseWriter, r *http.Request) {
	var user users.User
	var err error
	var token string
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	user.DBClient = pgClient
	res, err := user.CheckUser()
	if !res {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	ok, err := user.CheckPasswordExp()
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if ok {
		token, err = user.NewTokenString(false)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Token\":\"" + token + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Token created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
		return
	}
	token, err = user.NewTokenString(true)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte("{\"Token\":\"" + token + "\"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Token created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Users/users"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//GetToken -  handle function, for get new token
func GetToken(w http.ResponseWriter, r *http.Request) {
	var user users.User
	var err error
	var token string
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AuthError("Decode user error", err))
		return
	}
	user.DBClient = pgClient
	res, err := user.CheckUser()
	if !res {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AuthError("Check users error", err))
		return
	}
	ok, err := user.CheckPasswordExp()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AuthError("Check password exp error", err))
		return
	}
	if ok {
		token, err = user.NewTokenString(false)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AuthError("Get new token error", err))
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
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.AuthError("Get new token error", err))
		return
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte("{\"Token\":\"" + token + "\"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Token created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

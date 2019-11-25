package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Users/users"
)

//Middleware - middleware handle function, for check auth and set default headers
func Middleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Max-Age", "600")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
		} else {
			f(w, r)
		}
	}
}

//GetToken -  handle function, for get new token
func GetToken(w http.ResponseWriter, r *http.Request) {
	var user users.User
	var err error
	var token string
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	u := user.New()
	res, err := u.CheckUser()
	if !res {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	ok, err := u.CheckPasswordExp()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	if ok {
		token, err = u.NewTokenString(false)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Token\":\"" + token + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Token created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
		}
		return
	}
	token, err = u.NewTokenString(true)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte("{\"Token\":\"" + token + "\"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Token created, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

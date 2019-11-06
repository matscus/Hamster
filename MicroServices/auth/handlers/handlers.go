package handlers

import (
	"encoding/json"
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
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	}
	u := user.New()
	res, err := u.CheckUser()
	if !res {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		ok, err := u.CheckPasswordExp()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			if ok {
				token, err = u.NewTokenString(false)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Token\":\"" + token + "\"}"))
				}
			} else {
				token, err = u.NewTokenString(true)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("{\"Token\":\"" + token + "\"}"))
				}
			}
		}
	}
}

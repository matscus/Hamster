package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Users/users"
)

//GetAllUsers -  handle function, for get new token
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allusers, err := client.PGClient{}.New().GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	err = json.NewEncoder(w).Encode(allusers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

//Users -  create new users, update users and delete users
func Users(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	switch r.Method {
	case "POST":
		err = user.Create()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
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
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
			}
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
				w.WriteHeader(http.StatusInternalServerError)
				_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			} else {
				w.WriteHeader(http.StatusOK)
				_, errWrite := w.Write([]byte("{\"Message\":\"User deleted \"}"))
				if errWrite != nil {
					log.Printf("[ERROR] User deleted, but Not Writing to ResponseWriter due: %s", errWrite.Error())
				}
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"You cannot remove God \"}"))
			if errWrite != nil {
				log.Printf("[ERROR] ALARMA!!!, someone tried deleted the god, was send in the bus, but Not Writing to ResponseWriter due: %s", errWrite.Error())
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
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	err = user.ChangePassword()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	_, errWrite := w.Write([]byte("{\"Message\":\"Password changed \"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Password changed, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

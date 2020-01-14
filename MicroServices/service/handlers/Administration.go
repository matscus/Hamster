package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/matscus/Hamster/Package/JWTToken/jwttoken"
	"github.com/matscus/Hamster/Package/Services/service"
)

//Administration hadlrer for install service to host. Nor auto runnable
func Administration(w http.ResponseWriter, r *http.Request) {
	var err error
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	own := jwttoken.GetUser(strings.TrimSpace(splitToken[1]))
	s := service.Service{}
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	switch r.Method {
	case "POST":
		err = s.InstallServiceToRemoteHost(own)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Message\":\"Service install}"))
		return
	case "DELETE":
		if own == s.Own || own == "admin" {
			err = s.DeleteServiceToRemoteHost(own)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"Service install}"))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\": You are not a owner for this service}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
}

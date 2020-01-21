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
	s := service.Service{}
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	sshUser, err := pgClient.GetUserToHost(s.Host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	s.DBClient = pgClient
	authHeader := r.Header.Get("Authorization")
	splitToken := strings.Split(authHeader, "Bearer ")
	own := jwttoken.GetUser(strings.TrimSpace(splitToken[1]))
	switch r.Method {
	case "POST":
		serviceBin, err := s.DBClient.GetServiceBin(s.BinsID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
		s.Name = serviceBin.Name
		if s.RunSTR == "" {
			s.RunSTR = serviceBin.RunSTR
		}
		s.Type = serviceBin.Type
		err = s.Create(sshUser, own)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Message\":\"Service install}"))
		OnceInitAllData()
		return
	case "PUT":
		if own == s.Owner || own == "admin" {
			err = s.Update()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"Service install}"))
			InitGetResponseAllData()
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\": You are not a owner for this service}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		OnceInitAllData()
		return
	case "DELETE":
		if own == s.Owner || own == "admin" {
			service, err := pgClient.GetService(s.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
				return
			}
			s.Name = service.Name
			s.Type = service.Type
			s.Host = service.Host
			err = s.Delete(own)
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
		OnceInitAllData()
		return
	}
}

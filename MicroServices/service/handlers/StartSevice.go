package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Package/Services/service"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//StartSevice - handle for start services and update him status
func StartSevice(w http.ResponseWriter, r *http.Request) {
	var s service.Service
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ServiceError("Decode service error", err))
		return
	}
	tempService, _ := AllService[s.ID]
	s.Name = tempService.Name
	s.Type = tempService.Type
	s.Host = tempService.Host
	s.Mutex = tempService.Mutex
	if s.RunSTR == "" {
		s.RunSTR = tempService.RunSTR
		log.Print(s.RunSTR)
		user, _ := HostsAndUsers.Load(s.Host)
		err = s.Run(user.(string))
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ServiceError("Run error", err))
			return
		}
		s.Status = "run"
		//s.Mutex.Lock()
		UpdateState(&s)
		//s.Mutex.Unlock()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"Message\":\"service start\"}"))
		return
	}
	user, _ := HostsAndUsers.Load(s.Host)
	err = s.Run(user.(string))
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ServiceError("Run error", err))
		return
	}
	s.Status = "run"
	s.Mutex.Lock()
	UpdateState(&s)
	s.Mutex.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Message\":\"service start\"}"))
}

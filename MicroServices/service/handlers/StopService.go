package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Services/service"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//StopService - handle for stop services and update him status
func StopService(w http.ResponseWriter, r *http.Request) {
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
	user, _ := HostsAndUsers.Load(s.Host)
	err = s.Stop(user.(string))
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ServiceError("Stop error", err))
		return
	}
	s.Status = "stop"
	s.Mutex.Lock()
	UpdateState(&s)
	s.Mutex.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Message\":\"service stop\"}"))
}

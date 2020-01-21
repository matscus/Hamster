package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Services/service"
)

//StopService - handle for stop services and update him status
func StopService(w http.ResponseWriter, r *http.Request) {
	var s service.Service
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	tempService, _ := AllService[s.ID]
	s.Name = tempService.Name
	s.Type = tempService.Type
	user, _ := HostsAndUsers.Load(s.Host)
	err = s.Stop(user.(string))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		return
	}
	s.Status = "stop"
	s.Mutex.Lock()
	UpdateState(&s)
	s.Mutex.Unlock()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Message\":\"service stop\"}"))
}

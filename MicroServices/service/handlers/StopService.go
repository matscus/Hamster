package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Services/service"
)

//StopService - handle for stop services and update him status
func StopService(w http.ResponseWriter, r *http.Request) {
	var service service.Service
	var err error
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		user, _ := serv.HostsAndUsers.Load(service.Host)
		err = service.Stop(user.(string))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		} else {
			service.Status = "stop"
			service.Mutex.Lock()
			serv.UpdateState(&service)
			service.Mutex.Unlock()
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("{\"Message\":\"service stop\"}"))
		}
	}
}

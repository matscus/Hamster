package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Services/service"
)

//StartSevice - handle for start services and update him status
func StartSevice(w http.ResponseWriter, r *http.Request) {
	var service service.Service
	var err error
	err = json.NewDecoder(r.Body).Decode(&service)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		if service.RunSTR == "" {
			str, err := client.PGClient{}.GetServiceRunSTR(service.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
			service.RunSTR = str
			user, _ := serv.HostsAndUsers.Load(service.Host)
			err = service.Run(user.(string))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				service.Status = "run"
				service.Mutex.Lock()
				serv.UpdateState(&service)
				service.Mutex.Unlock()
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"service start\"}"))
			}
		} else {
			user, _ := serv.HostsAndUsers.Load(service.Host)
			err = service.Run(user.(string))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				service.Status = "run"
				service.Mutex.Lock()
				serv.UpdateState(&service)
				service.Mutex.Unlock()
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"service start\"}"))
			}
		}
	}
}

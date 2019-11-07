package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/Package/Clients/client"
	"github.com/matscus/Hamster/Package/Hosts/hosts"
)

//GetAllHosts -  handle function, for get new token
func GetAllHosts(w http.ResponseWriter, r *http.Request) {
	allhosts, err := client.PGClient{}.New().GetAllHosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err := json.NewEncoder(w).Encode(allhosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

//Hosts -  handle function, for new,update and delete host
func Hosts(w http.ResponseWriter, r *http.Request) {
	host := hosts.Host{}
	err := json.NewDecoder(r.Body).Decode(&host)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		client := client.PGClient{}.New()
		switch r.Method {
		case "POST":
			ok, err := client.HostIfExist(host.Host)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				if ok {
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("{\"Message\": Host is exist }"))
				} else {
					err = client.NewHost(host.Host, host.User, host.Type, host.Projects)
					if err != nil {
						w.WriteHeader(http.StatusInternalServerError)
						w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
					} else {
						w.WriteHeader(http.StatusOK)
						w.Write([]byte("{\"Message\":\"User created \"}"))
					}
				}
			}
		case "PUT":
			err = client.UpdateHost(host.ID, host.Host, host.User, host.Type, host.Projects)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"Host updated \"}"))
			}
		case "DELETE":
			err = client.DeleteHost(host.ID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("{\"Message\":\"Host deleted \"}"))
			}
		}
	}

	allhosts, err := client.PGClient{}.New().GetAllHosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
	} else {
		err := json.NewEncoder(w).Encode(allhosts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
	}
}

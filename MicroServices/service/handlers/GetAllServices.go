package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/service/serv"
	"github.com/matscus/Hamster/Package/Services/service"
)

//GetAllServices -  handle for response all services
func GetAllServices(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["project"]
	if ok {
		if len(serv.GetResponseAllData) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"len services slice equally 0\"}"))
		} else {
			serv.CheckService()
			l := len(serv.GetResponseAllData)
			res := make([]service.Service, 0, l)
			iter := 0
			for i := 0; i < l; i++ {
				projects := serv.GetResponseAllData[i].Projects
				for i := 0; i < len(projects); i++ {
					if projects[i] == page {
						res = append(res, serv.GetResponseAllData[iter])
						break
					}
				}
				iter++
			}
			err := json.NewEncoder(w).Encode(res)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
			}
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"Message\":\" params hot found \"}"))
	}
}

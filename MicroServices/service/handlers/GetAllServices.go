package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/Package/Services/service"
	"github.com/matscus/Hamster/Package/httperror"
)

//GetAllServices -  handle for response all services
func GetAllServices(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	project, ok := params["project"]
	if ok {
		res := make([]service.Service, 0, 20)
		CheckService()
		for _, v := range AllService {
			l := len(v.Projects)
			for i := 0; i < l; i++ {
				if v.Projects[i] == project {
					res = append(res, v)
				}
			}
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"Message\":\" params hot found \"}"))
}

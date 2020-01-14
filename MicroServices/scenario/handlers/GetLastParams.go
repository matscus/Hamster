package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
)

//GetLastParams - init slace for response last scenario params
func GetLastParams(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["name"]
	if ok {
		lastParamName, ok := params["project"]
		if ok {
			res, ок := scn.LastRunsParams.Load(page + lastParamName)
			if ок {
				params := res.(scn.StartRequest)
				err := json.NewEncoder(w).Encode(params)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					_, errWrite := w.Write([]byte("{\"Message\":\"Scenario encode json error: " + err.Error() + "\"}"))
					if errWrite != nil {
						log.Printf("[ERROR] Scenario encode json error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
					}
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			nilres := scn.StartRequest{}
			err := json.NewEncoder(w).Encode(nilres)
			if err != nil {
				_, errWrite := w.Write([]byte("{\"Message\":\"Scenario encode json error: " + err.Error() + "\"}"))
				if errWrite != nil {
					log.Printf("[ERROR] Scenario encode json error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
				}
			}
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	_, errWrite := w.Write([]byte("{\"Message\":\"Scenario params not found\"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Scenario params not found, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/httperror"
)

//GetScenarios - handle return state all scenario and generators
func GetScenarios(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["project"]
	if ok {
		if len(scn.GetResponseAllData.Scenarios) == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			_, errWrite := w.Write([]byte("{\"Message\":\"Scenario Len GetResponceAllData slice equally 0\"}"))
			if errWrite != nil {
				log.Printf("[ERROR] Scenario len GetResponceAllData slice equally 0, but Not Writing to ResponseWriter due: %s", errWrite.Error())
			}
			return
		}
		res := scn.GetResponse{}
		l := len(scn.GetResponseAllData.Scenarios)
		for i := 0; i < l; i++ {
			if scn.GetResponseAllData.Scenarios[i].Projects == page {
				res.Scenarios = append(res.Scenarios, scn.GetResponseAllData.Scenarios[i])
			}
		}
		scn.CheckRunsGen()
		res.Generators = scn.GetResponseAllData.Generators
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			httperror.WriteError(w, http.StatusInternalServerError, err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	_, errWrite := w.Write([]byte("{\"Message\":\"Scenario params hot found \"}"))
	if errWrite != nil {
		log.Printf("[ERROR] Scenario params not found, but Not Writing to ResponseWriter due: %s", errWrite.Error())
	}
}

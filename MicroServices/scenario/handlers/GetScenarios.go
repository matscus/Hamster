package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//GetScenarios - handle return state all scenario and generators
func GetScenarios(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	page, ok := params["project"]
	if ok {
		if len(scn.GetResponseAllData.Scenarios) == 0 {
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Scenario Len GetResponceAllData slice equally 0", nil))
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
			errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Encode res error", err))
			return
		}
		return
	}
	errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Params not found", nil))
}

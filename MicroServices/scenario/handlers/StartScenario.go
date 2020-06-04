package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//StartScenario - handle to start scenario
func StartScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StartRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Decode error", err))
		return
	}
	runsgen, err := scn.CheckGen(s.Generators)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(w).Encode(runsgen)
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Encode runsgen error", err))
			return
		}
		return
	}
	if len(runsgen) == 0 {
		scn.LastRunsParams.Store(s.Name+s.Projects, s)
		scn.RunsGenerators.Store(s.Name, s)
		err = s.Start()
		if err != nil {
			errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Start error", err))
			return
		}
		w.WriteHeader(http.StatusOK)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario launched\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario launched, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	err = json.NewEncoder(w).Encode(runsgen)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Encode runsgen error", err))
		return
	}
}

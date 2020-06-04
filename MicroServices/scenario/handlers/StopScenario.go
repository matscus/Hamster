package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/errorImpl"
)

//StopScenario - handle to stop scenario
func StopScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StopRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Decode error", err))
		return
	}
	err = s.Stop()
	if err != nil {
		errorImpl.WriteHTTPError(w, http.StatusOK, errorImpl.ScenarioError("Stop error", err))
		return
	}
}

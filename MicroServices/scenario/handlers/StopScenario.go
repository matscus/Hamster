package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/httperror"
)

//StopScenario - handle to stop scenario
func StopScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StopRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = s.Stop()
	if err != nil {
		httperror.WriteError(w, http.StatusInternalServerError, err)
		return
	}
}

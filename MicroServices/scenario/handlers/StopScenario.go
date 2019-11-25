package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/matscus/Hamster/MicroServices/scenario/scn"
)

//StopScenario - handle to stop scenario
func StopScenario(w http.ResponseWriter, r *http.Request) {
	var s scn.StopRequest
	var err error
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario decode json error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario decode json error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	err = s.Stop()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"Scenario stop error: " + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Scenario stop error, but Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

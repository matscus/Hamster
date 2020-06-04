package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matscus/Hamster/MicroServices/scenario/scn"
	"github.com/matscus/Hamster/Package/errorImpl"
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
					errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Encode params error", err))
				}
				return
			}
			w.WriteHeader(http.StatusOK)
			nilres := scn.StartRequest{}
			err := json.NewEncoder(w).Encode(nilres)
			if err != nil {
				errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Encode nilres error", err))
				return
			}
		}
		return
	}
	errorImpl.WriteHTTPError(w, http.StatusInternalServerError, errorImpl.ScenarioError("Params not found", nil))
}

package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	uuid "github.com/nu7hatch/gouuid"

	"github.com/matscus/Hamster/Mock/dadata/cache"
)

var (
	Mean       float64
	Deviation  float64
	Requestlog bool
)

func GetFIO(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Accel-Expires", "0")
	switch r.Method {
	case "GET":
		params := mux.Vars(r)
		query := params["query"]
		temp, ok := cache.FIO.Load(query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := fioResponce{}
			f.Suggestion = make([]suggestionFIO, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionFIO{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				sug.Data.Surname = res[i]
				f.Suggestion = append(f.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}
	case "POST":
		var fiorequest request
		err := json.NewDecoder(r.Body).Decode(&fiorequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
		temp, ok := cache.FIO.Load(fiorequest.Query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := fioResponce{}
			f.Suggestion = make([]suggestionFIO, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionFIO{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				sug.Data.Surname = res[i]
				f.Suggestion = append(f.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}
	}

}
func GetAddress(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Accel-Expires", "0")
	switch r.Method {
	case "GET":
		params := mux.Vars(r)
		query := params["query"]
		temp, ok := cache.Address.Load(query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := addressResponce{}
			f.Suggestion = make([]suggestionAddress, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionAddress{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				f.Suggestion = append(f.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}
	case "POST":
		var addressrequest request
		err := json.NewDecoder(r.Body).Decode(&addressrequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
		temp, ok := cache.Address.Load(addressrequest.Query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := addressResponce{}
			f.Suggestion = make([]suggestionAddress, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionAddress{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				f.Suggestion = append(f.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}
	}

}

func GetOrganization(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Accel-Expires", "0")
	switch r.Method {
	case "GET":
		params := mux.Vars(r)
		query := params["query"]
		temp, ok := cache.Organization.Load(query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := orgResponce{}
			f.ActualTimestamp = time.Now().Unix()
			f.Status = "success"
			f.Data.Suggestion = make([]suggestionOrg, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionOrg{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				f.Data.Suggestion = append(f.Data.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}

	case "POST":
		var orgrequest request
		err := json.NewDecoder(r.Body).Decode(&orgrequest)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		}
		temp, ok := cache.Organization.Load(orgrequest.Query)
		if ok {
			res := temp.([]string)
			l := len(res)
			f := orgResponce{}
			f.ActualTimestamp = time.Now().Unix()
			f.Status = "success"
			f.Data.Suggestion = make([]suggestionOrg, 0, l)
			for i := 0; i < l; i++ {
				var sug = suggestionOrg{}
				sug.Value = res[i]
				sug.UnrestrictedValue = res[i]
				f.Data.Suggestion = append(f.Data.Suggestion, sug)
			}
			json.NewEncoder(w).Encode(f)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			e := errResponce{}
			e.Status = "error"
			e.ActualTimestamp = time.Now().Unix()
			uuid, _ := uuid.NewV4()
			e.Error.ID = uuid.String()
			e.Error.Code = "exceptionData"
			e.Error.DisplayMessage = "params not found"
			e.Error.SystemID = "ADP-WILL"
			json.NewEncoder(w).Encode(e)
		}
	}

}

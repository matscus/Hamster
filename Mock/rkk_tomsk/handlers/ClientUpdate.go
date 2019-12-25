package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func ClientUpdateHandler(w http.ResponseWriter, r *http.Request) {
	var res structs.ClientUpdateResponse
	res.ReturnCode = 0
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
	err := xml.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

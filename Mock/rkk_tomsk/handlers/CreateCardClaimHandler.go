package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreateCardClaimRHandler(w http.ResponseWriter, r *http.Request) {
	var req structs.ClaimStatusRequest
	err := xml.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	var res structs.CreateCardClaimResponse
	res.ReturnCode = 11
	res.Message = "У клиента есть действующий договор на кредитную карту"
	res.InstitutionId = 0
	res.ContractReqId = 0
	res.CardReqId = 0
	res.DateExpirationOut = "0001-01-01T00:00:00"
	res.CrPsk = 0
	res.CrPskMoney = 0
	res.WorkDay = "1900-01-01T00:00:00"
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
	err = xml.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}

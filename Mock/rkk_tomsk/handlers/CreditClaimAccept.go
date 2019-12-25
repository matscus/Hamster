package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func CreditClaimAcceptHandler(w http.ResponseWriter, r *http.Request) {
	var res structs.CreditClaimAcceptResponse
	res.ReturnCode = 0
	res.InstitutionId = 500058
	res.ContractReqId = 441014
	res.CardReqId = 1091122
	res.CardReqStatus = "На ожидании выдачи"
	res.ContractId = 1058716
	res.ContractStatus = "Введен"
	res.ResourceId = 4086337
	res.AccNumber = "40817810810002702546"
	res.WorkDay = "1900-01-01T00:00:00"
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

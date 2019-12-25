package handlers

import (
	"encoding/xml"
	"log"
	"net/http"

	"github.com/matscus/Hamster/Mock/rkk_tomsk/structs"
)

func ClaimStatusHandler(w http.ResponseWriter, r *http.Request) {
	var statusReq structs.ClaimStatusRequest
	err := xml.NewDecoder(r.Body).Decode(&statusReq)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
		return
	}
	//необходимо обработать id инициатора когда будет пул данных
	var statusRes structs.ClaimStatusResponse
	statusRes.ReturnCode = 0
	statusRes.InstitutionId = 467465
	statusRes.ContractReqId = 441007
	statusRes.ContractStatus = "На ожидании подписи клиента"
	statusRes.ContractReqId = 441014
	statusRes.CardReqId = 1078994
	statusRes.CardReqStatus = "Действует"
	statusRes.CardBlankReqStatus = "Персонализирована принята по отчету"
	statusRes.ContractId = 1052248
	statusRes.ContractStatus = "Предоставлен"
	statusRes.CardReqLimitStatus = 0
	statusRes.CardNumberOut = 4249753212000506
	statusRes.DateExpirationOut = "2022-12-12T00:00:00+07:00"
	statusRes.CardEmbossingOut = "OLEG NIKIFOROV"
	statusRes.CrPsk = 27.6510
	statusRes.CrPskMoney = 2099.0800
	statusRes.WorkDay = "1900-01-01T00:00:00"
	w.Header().Set("content-type ", "text/xml")
	w.WriteHeader(http.StatusOK)
	err = xml.NewEncoder(w).Encode(statusRes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, errWrite := w.Write([]byte("{\"Message\":\"" + err.Error() + "\"}"))
		if errWrite != nil {
			log.Printf("[ERROR] Not Writing to ResponseWriter error %s due: %s", err.Error(), errWrite.Error())
		}
	}
}
